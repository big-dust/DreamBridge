package migration

import (
	"fmt"
	"github.com/big-dust/DreamBridge/internal/crawler/response"
	"github.com/big-dust/DreamBridge/internal/crawler/safe"
	"github.com/big-dust/DreamBridge/internal/crawler/scraper"
	"github.com/big-dust/DreamBridge/internal/model"
	"github.com/big-dust/DreamBridge/internal/pkg/common"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func Migrate() {
	defer LOGPageCount()
	for common.Page <= 145 {
		//学校基础信息
		list := safe.GetSchoolListSafe(common.Page)
		common.Page++ //下一页
		//proxy.ChangeHttpProxyIP()
		time.Sleep(2 * time.Second)
		for i, item := range list.Data.Item {
			wg.Add(1)
			go MigrateSchoolScores(i, item)
		}
	}
	wg.Wait()
}

func MigrateSchoolScores(i int, item response.Item) {
	defer func() {
		if r := recover(); r != nil {
			common.LOG.Error(fmt.Sprintf("MigrateSchoolScores: Panic: %v", r))
		}
		wg.Done()
		common.Mu.Lock()
		common.Count++
		common.Mu.Unlock()
		fmt.Printf("\rNumber: %d", common.Count)
	}()
	common.LOG.Info(fmt.Sprintf("当前的爬取第[ %d ]所大学：%s", (common.Page-2)*5+i+1, item.Name))
	// 学校具体信息
	info := safe.GetSchoolInfoSafe(item.SchoolID)
	// 三个指标
	master, _ := strconv.Atoi(info.Data.NumMaster)
	doctor, _ := strconv.Atoi(info.Data.NumDoctor)
	gbh, _ := strconv.Atoi(info.Data.GbhNum)
	// 毕业去向
	promotion, abroad, job := safe.ToWhereSafe(item.SchoolID)
	// 双一流学科text
	text := TextDualClass(info)

	school := &model.School{
		ID:                          item.SchoolID,
		Name:                        item.Name,
		BriefIntroduction:           info.Data.Content,
		SchoolCode:                  item.CodeEnroll,
		MasterPoint:                 master,
		PhdPoint:                    doctor,
		ResearchProject:             gbh,
		TitleDoubleFirstClass:       item.DualClassName == "双一流",
		Title_985:                   item.F985 == 1,
		Title_211:                   item.F211 == 1,
		TitleCollege:                info.Data.LevelName == "普通本科",
		TitleUndergraduate:          info.Data.LevelName == "专科（高职）",
		Region:                      info.Data.CityName,
		Website:                     info.Data.Site,
		RecruitmentPhone:            info.Data.Phone,
		Email:                       info.Data.Email,
		PromotionRate:               promotion,
		AbroadRate:                  abroad,
		EmploymentRate:              job,
		DoubleFirstClassDisciplines: text,
	}
	var scores []*model.Score
	// 学校各省历年分数（这里只针对湖北省）
	for year := 2021; year <= 2023; year++ {
		// 物理类
		scores = append(scores, safe.GetScoresSafe(school.ID, common.HuBei, common.T_Physics, year)...)
		// 历史类
		scores = append(scores, safe.GetScoresSafe(school.ID, common.HuBei, common.T_History, year)...)
	}
	for year := 2018; year <= 2020; year++ {
		// 理科
		scores = append(scores, safe.GetScoresSafe(school.ID, common.HuBei, common.T_li, year)...)
		// 文科
		scores = append(scores, safe.GetScoresSafe(school.ID, common.HuBei, common.T_wen, year)...)
	}
	// 去重
	scoresDistinct := make(map[string]*model.Score, 30)
	for _, score := range scores {
		key := fmt.Sprintf("%d-%d-%d-%d-%s-%s-%s", score.SchoolID, score.Location, score.TypeId, score.Year, score.Tag, score.SgName, score.BatchName)
		scoresDistinct[key] = score
	}
	// Must！
	model.MustCreateSchoolScore(school, scoresDistinct)
}

func MigrateSpecialScores() {
	for schoolId := 0; schoolId < 4000; schoolId++ {
		specialInfos, err := scraper.SpecialInfo(schoolId)
		if err != nil {
			common.LOG.Error("SpecialInfo: " + err.Error())
		}
		// 拿到school的所有专业
		var specials []*model.Major
		for _, info := range specialInfos.Data.SpecialDetail["1"] {
			id, _ := strconv.Atoi(info.ID)
			special := &model.Major{
				ID:                 id,
				Name:               info.SpecialName,
				NationalFeature:    info.NationFeature == "1", //guo
				Level:              info.TypeName,
				DisciplineCategory: info.Level2Name,
				MajorCategory:      info.Level3Name,
				LimitYear:          info.LimitYear,
				SchoolID:           schoolId,
				SpecialId:          info.SpecialID,
			}
			specials = append(specials, special)
		}
		recruit, err := scraper.HistoryRecruit(schoolId, common.HuBei)
		if err != nil {
			common.LOG.Error("HistoryRecruit:" + err.Error())
		}

		scores := make(map[string]*model.MajorScore, 60)
		for major_id, major := range recruit.Data {
			for kelei_id, kelei := range major {
				for _, yearInfo := range kelei {
					major_id, _ := strconv.Atoi(major_id)
					year, _ := strconv.Atoi(yearInfo.Year)
					num, _ := strconv.Atoi(yearInfo.Num)
					keleiStr := common.Kelei(kelei_id)
					score := &model.MajorScore{
						SpecialID:         major_id,
						Location:          "湖北",
						Year:              year,
						Kelei:             keleiStr,
						Batch:             yearInfo.Batch,
						RecruitmentNumber: num,
						LowestScore:       0,
						LowestRank:        0,
					}
					key := fmt.Sprintf("%d-%s-%d-%s-%s", major_id, "湖北", year, keleiStr, yearInfo.Batch)
					if _, exist := scores[key]; !exist {
						scores[key] = score
						continue
					}
					common.LOG.Info("key exist: " + key)
				}
			}
		}
		admission, _ := scraper.HistoryAdmission(schoolId, common.HuBei)
		for major_id, major := range admission.Data {
			for kelei_id, kelei := range major {
				for _, yearInfo := range kelei {
					major_id, _ := strconv.Atoi(major_id)
					year, _ := strconv.Atoi(yearInfo.Year)
					keleiStr := common.Kelei(kelei_id)
					key := fmt.Sprintf("%d-%s-%d-%s-%s", major_id, "湖北", year, keleiStr, yearInfo.Batch)
					lowestScore, _ := strconv.Atoi(yearInfo.Min)
					rank, _ := strconv.Atoi(yearInfo.MinSection)
					score, exist := scores[key]
					if !exist {
						score := &model.MajorScore{
							SpecialID:         major_id,
							Location:          "湖北",
							Year:              year,
							Kelei:             keleiStr,
							Batch:             yearInfo.Batch,
							RecruitmentNumber: 0,
							LowestScore:       lowestScore,
							LowestRank:        rank,
						}
						scores[key] = score
						continue
					}
					score.LowestScore = lowestScore
					score.LowestRank = rank
				}
			}
		}
	}
}

func TextDualClass(info *response.SchoolInfoResponse) string {
	var text string
	for _, dualclass := range info.Data.Dualclass {
		text += dualclass.Class + " "
	}
	return text
}

func LOGPageCount() {
	if r := recover(); r != nil {
		common.LOG.Error(fmt.Sprintf("%v", r))
		common.LOG.Info(string(debug.Stack()))
	}
	common.LOG.Info(fmt.Sprintf("page: %d ,count: %d", common.Page, common.Count))
}
