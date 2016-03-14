package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/olebedev/config"
)

type Config struct {
	*config.Config
}

func (cfg Config) String(key string) (string, error) {
	env := os.Getenv("GIN_ENV")
	return cfg.Config.String(fmt.Sprintf("%s.%s", env, key))
}

func ConnectDB() *gorm.DB {
	path := "config/database.yml"
	file, err := os.Open(path) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.ParseYaml(string(data))
	if err != nil {
		log.Fatal(err)
	}
	x := &Config{Config: conf}
	user, _ := x.String("user")
	pass, _ := x.String("password")
	dbname, _ := x.String("database")
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, pass, dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("cannot connect to %s", dsn)
	}
	return db
}
