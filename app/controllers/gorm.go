package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"keepassweb/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    // mysql package
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres package
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // mysql package
	r "github.com/revel/revel"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// DB gorm db
var DB *gorm.DB

// InitDB database
func InitDB() {
	var DBDRIVER = r.Config.StringDefault("app.dbdriver", "postgres")
	var DBHOST = r.Config.StringDefault("app.dbhost", "localhost")
	if DBDRIVER == "sqlite" && DBHOST == "localhost" {
		DBHOST = "/tmp/gorm.db"
	}
	var DBUSER = r.Config.StringDefault("app.dbuser", "postgres")
	var DBPASSWORD = r.Config.StringDefault("app.dbpassword", "")
	var DBNAME = r.Config.StringDefault("app.dbname", "shopping")
	dbinfo := fmt.Sprintf("")

	switch DBDRIVER {
	default:
		dbinfo = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", DBHOST, DBUSER, DBNAME, DBPASSWORD)
	case "postgres":
		dbinfo = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", DBHOST, DBUSER, DBNAME, DBPASSWORD)
	case "mysql":
		dbinfo = fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", DBUSER, DBPASSWORD, DBHOST, DBNAME)
	case "sqlite":
		dbinfo = fmt.Sprintf(DBHOST)
	}

	db, err := gorm.Open(DBDRIVER, dbinfo)
	if err != nil {
		checkErr(err, "sql.Open failed")
	}
	DB = db
	DB.AutoMigrate(&models.User{})

}

// GormController controllers begin, commit and rollback transactions
type GormController struct {
	*r.Controller
	Txn *gorm.DB
}

// Begin GormController to connect db
func (c *GormController) Begin() r.Result {
	txn := DB.Begin()
	if txn.Error != nil {
		fmt.Println(c.Txn.Error)
		panic(txn.Error)
	}

	c.Txn = txn
	return nil
}

// Commit database transaction
func (c *GormController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}

	c.Txn.Commit()
	if c.Txn.Error != nil && c.Txn.Error != sql.ErrTxDone {
		fmt.Println(c.Txn.Error)
		panic(c.Txn.Error)
	}

	c.Txn = nil
	return nil
}

// Rollback transaction
func (c *GormController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}

	c.Txn.Rollback()
	if c.Txn.Error != nil && c.Txn.Error != sql.ErrTxDone {
		fmt.Println(c.Txn.Error)
		panic(c.Txn.Error)
	}

	c.Txn = nil
	return nil
}
