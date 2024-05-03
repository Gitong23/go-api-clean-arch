package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/Gitong23/go-api-clean-arch/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgresDatabaseInstance *postgresDatabase
	once                     sync.Once
)

func NewPostgresDatabase(conf *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf("%s", conf.HostUrl)
		log.Printf("Connecting to Postgres: %s", dsn)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Println("Connected to Postgres")

		postgresDatabaseInstance = &postgresDatabase{conn}
	})

	return postgresDatabaseInstance
}

func (db *postgresDatabase) Connect() *gorm.DB {
	return db.DB
}

// func NewPostgresDatabase(conf *config.Database) Database{
// 	once.Do(func (){
// 		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=",
// 		conf.Host,
// 		conf.User,
// 		conf.Password,
// 		conf.DBName,
// 		conf.Port)
// 	})
// }
