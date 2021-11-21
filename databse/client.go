package databse

import (
	"github.com/jinzhu/gorm"
	"log"
	"project/models"
)

var Connector *gorm.DB

func ConnectToMysql(c string) error {
	var err error
	Connector, err = gorm.Open("mysql", c)
	if err != nil {
		return err
	}

	log.Println("connect to db")
	return nil
}

func Migrate(t *models.Product) {
	Connector.AutoMigrate(&t)

	log.Println("migrated successfully")
}
