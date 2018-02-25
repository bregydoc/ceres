package main

import "github.com/k0kubun/pp"

func main() {

	configProject, err := GetConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}

	types, err := GetTypesDefinitions(configProject)
	if err != nil {
		panic(err)
	}

	pp.Println(types)
	// pp.Println(configProject)

	// err = GenerateFolderStruct(configProject.Project.Name, true)
	// if err != nil {
	// 	panic(err)
	// }

	// err = GenerateDbLinks(configProject)
	// if err != nil {
	// 	panic(err)
	// }

	// err = GenerateHelpers(configProject)
	// if err != nil {
	// 	panic(err)
	// }

	// err = GenerateAPILinkers(configProject)
	// if err != nil {
	// 	panic(err)
	// }

	// err = GenerateServerMain(configProject)
	// if err != nil {
	// 	panic(err)
	// }
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"

// 	"github.com/urfave/cli"
// )

// var initPath string

// func main() {
// 	app := cli.NewApp()
// 	app.Name = CeresLogo
// 	app.Usage = "A new tool for generate API REST using Boltdb and Iris"
// 	app.Description = "Ceres is a simple and hackable API REST generator."

// 	app.Commands = []cli.Command{
// 		{
// 			Name:    "init",
// 			Aliases: []string{"i"},
// 			Usage:   "create a new ceres config file",
// 			Flags: []cli.Flag{
// 				cli.StringFlag{
// 					Name:        "path, p",
// 					Value:       ".",
// 					Destination: &initPath,
// 				},
// 			},
// 			Action: func(c *cli.Context) error {
// 				reader := bufio.NewReader(os.Stdin)
// 				fmt.Print("Enter text: ")
// 				text, _ := reader.ReadString('\n')
// 				fmt.Println(text)

// 				return nil
// 			},
// 		},
// 		{
// 			Name:    "add",
// 			Aliases: []string{"a"},
// 			Usage:   "add a task to the list",
// 			Action: func(c *cli.Context) error {
// 				return nil
// 			},
// 		},
// 	}

// 	sort.Sort(cli.FlagsByName(app.Flags))
// 	sort.Sort(cli.CommandsByName(app.Commands))

// 	app.Run(os.Args)
// }
