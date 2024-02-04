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

func Kelei(type_id string) string {
	return ""
}
