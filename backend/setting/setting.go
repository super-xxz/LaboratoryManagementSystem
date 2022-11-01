package setting

import (
	"log"

	"github.com/BurntSushi/toml"
)

type config struct {
	mysql Mysql
}

type Mysql struct {
	username string
	password string
	host     string
	dbname   string
}

var conf config

func init() {
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		log.Fatal(err)
	}

}
