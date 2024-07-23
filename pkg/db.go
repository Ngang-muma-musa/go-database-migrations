package pkg

import (
	"database/sql"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/config"
)

func NewDBClient(config *config.Config) (*gorm.DB, error) {
	sqlDB, err := sql.Open("mysql", config.MysqlConnection)
	if err != nil {
		return nil, err
	}
	var db *gorm.DB
	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	return db, err
}
