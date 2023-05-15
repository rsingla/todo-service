package mydb

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rsingla/todo-service/entity"
)

var db *gorm.DB

func Connect() *gorm.DB {

	var datetimePrecision = 2
	var err error

	mysqlConfig := mysql.Config{
		DSN:                       "root:my-secret-pw@tcp(localhost:3306)/todo_service?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                                                                                        // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,                                                                                       // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision,                                                                         // default datetime precision
		DontSupportRenameIndex:    true,                                                                                       // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                       // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                      // smart configure based on used version
	}

	db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}

	err = db.AutoMigrate(&entity.TodoEntity{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
