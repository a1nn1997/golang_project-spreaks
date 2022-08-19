package config

import(
	"github.com/jinzhu/gorm" //for connecting db with go
	_ "github.com/jinzhu/gorm/dialects/mysql"  //mysql db
	_ "github.com/lib/pq"   //postgresql db
)

var (
	db * gorm.DB //to configure db with gorm
)

func Connect(){
	d, err := gorm.Open("postgres","host=localhost dbname=go1 user=postgres password=123 port=5432") //db specs
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db  //send to other fns
}