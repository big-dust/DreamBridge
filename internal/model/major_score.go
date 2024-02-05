package model

import (
	"fmt"
	"github.com/big-dust/DreamBridge/internal/pkg/common"
	"gorm.io/gorm"
	"time"
)

// 定义 MajorScores 表格的模型
type MajorScore struct {
	ID                int
	SpecialID         int
	Location          string
	Year              int
	Kelei             string
	Batch             string
	RecruitmentNumber int
	LowestScore       int
	LowestRank        int
}

// 创建记录
func CreateMajorScore(db *gorm.DB, majorScore *MajorScore) {
	db.Create(majorScore)
}

// 查询记录
func GetMajorScoresBySpecialID(db *gorm.DB, specialID int) ([]MajorScore, error) {
	var majorScores []MajorScore
	err := db.Where("special_id = ?", specialID).Find(&majorScores).Error
	return majorScores, err
}

// 更新记录
func UpdateMajorScore(db *gorm.DB, majorScore *MajorScore) {
	db.Save(majorScore)
}

// 删除记录
func DeleteMajorScore(db *gorm.DB, specialID int) {
	db.Delete(&MajorScore{}, "special_id = ?", specialID)
}

func CreateMajorScores(major []*Major, majorScores map[string]*MajorScore) error {
	tx := common.DB.Begin()
	if err := tx.Create(major).Error; err != nil {
		tx.Rollback()
		if common.ErrMysqlDuplicate.Is(err) {
			return nil
		}
		common.LOG.Error("MustCreateMajorScores: " + err.Error())
		return err
	}
	for _, score := range majorScores {
		if err := tx.Create(score).Error; err != nil {
			tx.Rollback()
			if common.ErrMysqlDuplicate.Is(err) {
				return nil
			}
			common.LOG.Error("MustCreateMajorScores: " + err.Error())
			return err
		}
	}
	tx.Commit()
	return nil
}

func MustCreateMajorScores(major []*Major, majorScores map[string]*MajorScore) {
	tryCount := 0
	for {
		tryCount++
		errChan := make(chan error, 1)
		nilChan := make(chan error, 1)
		go func() {
			err := CreateMajorScores(major, majorScores)
			if err != nil {
				errChan <- err
				return
			}
			nilChan <- nil
		}()
		ticker := time.NewTicker(15 * time.Second)
		select {
		case <-nilChan:
			return
		case <-ticker.C:
			select {
			case err := <-errChan:
				common.LOG.Error(fmt.Sprintf("MustCreateSchoolScore:	Time Out 15s TryCount:%d Major.SchoolId: %v Error: %s", tryCount, major[0].SchoolID, err.Error()))
			default:
				common.LOG.Error(fmt.Sprintf("MustCreateSchoolScore:	Time Out 15s TryCount:%d Major.SchoolId: %v", tryCount, major[0].SchoolID))
			}
		}
	}
}
