package model

import "gorm.io/gorm"

// 定义 Majors 表格的模型
type Major struct {
	ID                 int
	Name               string `gorm:"not null"`
	NationalFeature    bool
	Level              string
	DisciplineCategory string
	MajorCategory      string
	LimitYear          string
	SchoolID           int
	SpecialId          string
}

// 创建记录
func CreateMajor(db *gorm.DB, major *Major) {
	db.Create(major)
}

// 查询记录
func GetMajorByID(db *gorm.DB, id int) (Major, error) {
	var major Major
	err := db.Where("id = ?", id).First(&major).Error
	return major, err
}

// 更新记录
func UpdateMajor(db *gorm.DB, major *Major) {
	db.Save(major)
}

// 删除记录
func DeleteMajor(db *gorm.DB, id int) {
	db.Delete(&Major{}, id)
}
