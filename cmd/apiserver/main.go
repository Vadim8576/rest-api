package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/Vadim8576/REST-API-GO/cmd/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "congig-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	err1 := http.ListenAndServe(":8080", nil)

	fmt.Println("Error creating server", err1)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
