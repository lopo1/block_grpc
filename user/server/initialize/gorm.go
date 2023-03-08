package initialize

import (
	"block_tool/user/server/global"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	//case "oracle":
	//	return GormOracle()
	//case "mssql":
	//	return GormMssql()
	default:
		return GormMysql()
	}
}
