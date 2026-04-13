package sqlite

import (
	"log"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func SetSqlite(dbname string) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("创建本地sqlite数据库目录失败:%s", err.Error())
	}
	location := filepath.Join(home, dbname)
	db, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	if err != nil {
		log.Fatalf("打开本地sqlite数据库失败:%s", err.Error())
	}

	// 优化SQLite性能配置
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取底层数据库连接失败:%s", err.Error())
	}

	// 启用WAL模式（Write-Ahead Logging），提高并发写入性能
	sqlDB.Exec("PRAGMA journal_mode=WAL")
	// 设置同步模式为NORMAL，平衡性能和安全性
	sqlDB.Exec("PRAGMA synchronous=NORMAL")
	// 增加缓存大小（单位：KB）
	sqlDB.Exec("PRAGMA cache_size=-64000") // 64MB缓存
	// 设置临时存储到内存
	sqlDB.Exec("PRAGMA temp_store=MEMORY")

	gormDB = db
	log.Println("本地sqlite数据库初始化完成")
}

func GetSqlite() *gorm.DB {
	return gormDB
}
