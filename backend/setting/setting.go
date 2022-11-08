package setting

import (
	"io/ioutil"
	"log"
	"os"

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
	// if _, err := toml.DecodeFile("./config.toml", &Conf); err != nil {
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

	err = toml.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatal(err)
	}

}
