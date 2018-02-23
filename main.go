package main

import (
	"github.com/k0kubun/pp"
)

func main() {

	configProject, err := GetConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}

	pp.Println(configProject)

	err = GenerateDbLinks(configProject)
	if err != nil {
		panic(err)
	}

	err = GenerateHelpers(configProject)
	if err != nil {
		panic(err)
	}

}
