package main

import (
	"fmt"

	"github.com/Gitong23/go-api-clean-arch/config"
	"github.com/Gitong23/go-api-clean-arch/databases"
)

func main() {
	conf := config.ConfigGetting()
	// fmt.Println(conf.Database.HostUrl)
	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.ConnectionGetting())
}
