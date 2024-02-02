package model

import (
	"github.com/big-dust/DreamBridge/internal/pkg/common"
	"gorm.io/gorm"
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
