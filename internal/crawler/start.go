package crawler

import (
	"github.com/big-dust/DreamBridge/internal/crawler/migration"
	"github.com/big-dust/DreamBridge/internal/pkg/common"
	"github.com/big-dust/DreamBridge/pkg/config"
	"github.com/big-dust/DreamBridge/pkg/gorm"
	"github.com/big-dust/DreamBridge/pkg/proxy"
	"github.com/big-dust/DreamBridge/pkg/zap"
	"time"
)

func Start() {
	// 初始化配置
	common.CONFIG = config.New("./config/config.toml")
	// 日志配置
	common.LOG = zap.AddZap()
	// 连接数据库
	DB, err := gorm.NewGorm()
	if err != nil {
		panic("gorm:" + err.Error())
	}
	common.DB = DB
	// 开启代理
	if common.CONFIG.Bool("proxy.switchon") {
		go proxy.ChangeHttpProxyIP()
		time.Sleep(2 * time.Second)
	}
	// 爬取数据迁移到数据库
	migration.Migrate()
}
