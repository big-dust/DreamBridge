package migration

import (
	"encoding/json"
	"fmt"
	"github.com/big-dust/DreamBridge/internal/crawler/scraper"
	"github.com/big-dust/DreamBridge/internal/model"
	"github.com/big-dust/DreamBridge/internal/pkg/common"
	"strconv"
)

func Migrate() {
	page := 1
	for {
		//学校基础信息
		list, err := scraper.SchoolList(page)
		if list == nil {
			break //结束
		}
		if err != nil {
			common.LOG.Panic("获取学校基础信息失败:" + err.Error())
		}
		page++ //下一页

		for _, item := range list.Data.Item {
			//go func(item response.Item) {
			// 学校具体信息
			info, err := scraper.SchoolInfo(item.SchoolID)
			if err != nil {
				common.LOG.Panic("学校具体信息：" + err.Error())
			}
			master, _ := strconv.Atoi(info.Data.NumMaster)
			doctor, _ := strconv.Atoi(info.Data.NumDoctor)
			gbh, _ := strconv.Atoi(info.Data.GbhNum)

			promotion, abroad, job := ToWhere(item.SchoolID)

			// 双一流学科text
			var text string
			for _, dualclass := range info.Data.Dualclass {
				text += dualclass.Class + " "
			}

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
			err = model.CreateSchool(school)
			if err != nil {
				common.LOG.Error("create school:" + err.Error())
			}
			// 学校各省历年分数（这里只针对湖北省）
			for year := 2021; year <= 2023; year++ {
				// 物理类
				DownloadScore(school.ID, common.HuBei, common.T_Physics, year)
				//历史类
				DownloadScore(school.ID, common.HuBei, common.T_History, year)
			}
			for year := 2018; year <= 2020; year++ {
				// 理科
				DownloadScore(school.ID, common.HuBei, common.T_li, year)
				// 文科
				DownloadScore(school.ID, common.HuBei, common.T_wen, year)
			}
			//}(item)
		}
	}
}

func ToWhere(schoolId int) (promotion string, abroad string, job string) {
	// 学生毕业去向
	detail, err := scraper.JobDetail(schoolId)
	if err != nil {
		if _, ok := err.(*json.UnmarshalTypeError); !ok {
			common.LOG.Info("jobdetail: " + err.Error())
		} else {
			common.LOG.Panic("jobdetail:" + err.Error())
		}
	} else {
		promotion = detail.Data.Jobrate.Postgraduate.One
		abroad = detail.Data.Jobrate.Abroad.One
		job = detail.Data.Jobrate.Job.One
	}
	return
}

func DownloadScore(schooldId int, provinceId int, type_id int, year int) {
	score, err := scraper.ProvinceScore(schooldId, provinceId, type_id, year)
	if err != nil {
		if _, ok := err.(*json.UnmarshalTypeError); !ok {
			common.LOG.Info(fmt.Sprintf("get score: schoolId:%d , tyid: %d, year: %d, %s", schooldId, type_id, year, err.Error()))
			return
		} else {
			common.LOG.Panic("jobdetail:" + err.Error())
		}
	}
	for _, item := range score.Data.Item {
		score := &model.Score{
			SchoolID: schooldId,
			Location: common.HuBei,
			Year:     year,
			TypeId:   type_id,
			Tag:      item.ZslxName, // 专项...
			Lowest:   item.Min,
			//LowestRank: item.MinSection,
		}
		model.CreateScore(score)
	}
}
