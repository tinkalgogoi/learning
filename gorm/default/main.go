package main
import (
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type UserModel struct {
	Id      int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
	RollNo  uint64 `gorm:"default:30"`
}

func main() {
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=exercise password=postgres sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	db.DropTableIfExists(&UserModel{})
	// In gorm default values works only when its auto migrated
	// Else default will restrict atleast putting zero values and allow sql default value to be inserted
	db.AutoMigrate(&UserModel{})

	user := &UserModel{Name: "John", Address: "New York"}

	db.Debug().Create(user)
}
