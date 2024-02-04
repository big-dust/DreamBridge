package model

import (
	"fmt"
	"github.com/big-dust/DreamBridge/internal/pkg/common"
	"gorm.io/gorm"
	"time"
)

// 定义 Schools 表格的模型
type School struct {
	ID                          int
	Name                        string `gorm:"unique;not null"`
	BriefIntroduction           string
	SchoolCode                  string `gorm:"unique"`
	MasterPoint                 int
	PhdPoint                    int
	ResearchProject             int
	TitleDoubleFirstClass       bool
	Title_985                   bool
	Title_211                   bool
	TitleCollege                bool
	TitleUndergraduate          bool
	Region                      string
	Website                     string
	RecruitmentPhone            string
	Email                       string
	PromotionRate               string
	AbroadRate                  string
	EmploymentRate              string
	DoubleFirstClassDisciplines string
}

// 创建记录
func CreateSchool(school *School) error {
	return common.DB.Create(school).Error
}

// 查询记录
func GetSchoolByID(db *gorm.DB, id int) (School, error) {
	var school School
	err := db.Where("id = ?", id).First(&school).Error
	return school, err
}

// 更新记录
func UpdateSchool(db *gorm.DB, school *School) {
	db.Save(school)
}

// 删除记录
func DeleteSchool(db *gorm.DB, id int) {
	db.Delete(&School{}, id)
}

func CreateSchoolScore(school *School, scores map[string]*Score) error {
	tx := common.DB.Begin()

	if err := tx.Create(school).Error; err != nil {
		common.LOG.Error("CreateSchoolScore: " + err.Error())
		tx.Rollback()
		if common.ErrMysqlDuplicate.Is(err) {
			return nil
		}
		return err
	}
	for _, score := range scores {
		if err := tx.Create(score).Error; err != nil {
			tx.Rollback()
			common.LOG.Error("CreateSchoolScore: " + err.Error())
			return err
		}
	}
	tx.Commit()
	return nil
}

func MustCreateSchoolScore(school *School, scores map[string]*Score) {
	tryCount := 0
this:
	for {
		errChan := make(chan error, 1)
		nilChan := make(chan error, 1)
		go func() {
			err := CreateSchoolScore(school, scores)
			if err != nil {
				errChan <- err
				return
			}
			nilChan <- nil
		}()
		ticker := time.NewTicker(10 * time.Second)
		select {
		case <-nilChan:
			break this
		case <-ticker.C:
			select {
			case err := <-errChan:
				common.LOG.Error("MustCreateSchoolScore: err:" + err.Error())
			default:
			}
			tryCount++
			common.LOG.Error(fmt.Sprintf("MustCreateSchoolScore:	Time Out 10s TryCount:%d schoolName:", tryCount, school.Name))
		}
	}
}
