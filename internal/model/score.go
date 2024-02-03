package model

import (
	"github.com/big-dust/DreamBridge/internal/pkg/common"
	"gorm.io/gorm"
)

// 定义 Scores 表格的模型
type Score struct {
	SchoolID   int
	Location   int    //地区
	Year       int    //年份
	TypeId     int    //科类
	BatchName  string //批次
	Tag        string //类型
	SgName     string //专业组
	Lowest     int
	LowestRank int
}

// 创建记录
func CreateScore(score *Score) {
	if err := common.DB.Create(score).Error; err != nil {
		common.LOG.Error("create score: " + err.Error())
	}
}

// 查询记录
func GetScoresBySchoolID(db *gorm.DB, schoolID int) ([]Score, error) {
	var scores []Score
	err := db.Where("school_id = ?", schoolID).Find(&scores).Error
	return scores, err
}

// 更新记录
func UpdateScore(db *gorm.DB, score *Score) {
	db.Save(score)
}

// 删除记录
func DeleteScore(db *gorm.DB, schoolID int) {
	db.Delete(&Score{}, "school_id = ?", schoolID)
}
