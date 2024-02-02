package model

import "gorm.io/gorm"

// 定义 MajorScores 表格的模型
type MajorScore struct {
	SpecialID         int
	Location          string
	Year              int
	Kelei             string
	RecruitmentNumber int
	HighestScore      int
	LowestScore       int
	HighestRank       int
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
