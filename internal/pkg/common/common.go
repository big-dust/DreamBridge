package common

import (
	"github.com/go-sql-driver/mysql"
	"github.com/gookit/config/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

// 数据库映射
var (
	DB        *gorm.DB
	CONFIG    *config.Config
	LOG       *zap.Logger
	HuBei     = 42
	Page      = 1
	Mu        = &sync.Mutex{}
	Count     = (Page - 1) * 5
	T_li      = 1
	T_wen     = 2
	T_Physics = 2073
	T_History = 2074
)

var (
	ErrMysqlDuplicate = &mysql.MySQLError{
		Number:   1062,
		SQLState: [5]byte{2, 3, 0, 0, 0},
		Message:  "",
	}
)

// typeId: 2-7文科，1-7理科，2074-14历史类，2073-14物理类
func Kelei(type_id string) string {
	if type_id[1] == '-' {
		type_id = type_id[:2]
	}
	switch type_id {
	case "2-":
		return "文科"
	case "1-":
		return "理科"
	case "2074-14":
		return "历史类"
	case "2073-14":
		return "物理类"
	}
	return ""
}
