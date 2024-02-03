package common

import (
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
	Page      = 29
	Mu        = &sync.Mutex{}
	Count     = Page * 5
	T_li      = 1
	T_wen     = 2
	T_Physics = 2073
	T_History = 2074
)
