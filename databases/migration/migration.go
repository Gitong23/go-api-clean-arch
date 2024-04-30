package main

import (
	"github.com/Gitong23/go-api-clean-arch/config"
	"github.com/Gitong23/go-api-clean-arch/databases"
	"github.com/Gitong23/go-api-clean-arch/entities"
)

func main() {
	conf := config.ConfigGetting()
	// fmt.Println(conf.Database.HostUrl)
	db := databases.NewPostgresDatabase(conf.Database)

	playerMigration(db)
	adminMigration(db)
	itemMigration(db)
	playerCoinMigration(db)
	inventoryMigration(db)
	purchaseHistoryMigration(db)
}

func playerMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Player{})
}

func adminMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PurchaseHistory{})
}
