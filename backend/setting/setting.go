package setting

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	mysql Mysql
}

type Mysql struct {
	Username string
	Password string
	Host     string
	Dbname   string
}

var Conf Config

func init() {
	if _, err := toml.DecodeFile("./config.toml", &Conf); err != nil {
		log.Fatal(err)
	}

}
