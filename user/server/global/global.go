package global

import (
	"block_tool/user/server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	GVA_DB *gorm.DB

	GVA_VP                  *viper.Viper
	GVA_CONFIG              *config.ServerConfig
	GVA_LOG                 *zap.SugaredLogger
	GVA_Concurrency_Control = &singleflight.Group{}
)
