package setting

import (
	"io/ioutil"
	"log"
	"os"

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
	// if _,err := toml.DecodeFile("./config.toml",&conf); err != nil {
	// 	log.Fatal(err)
	// }
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = toml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal(err)
	}

}
