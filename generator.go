package main

import (
	"fmt"
	"log"
	"os"
)

// GenerateFolderStruct ...
func GenerateFolderStruct(basePath string, cleanCreation bool) error {
	if cleanCreation {
		err := os.Mkdir(basePath, 0755)
		if os.IsExist(err) {
			return err
		}
	}

	helpersFolder := Join(basePath, HelpersFolderName)
	err := os.Mkdir(helpersFolder, 0755)
	if os.IsExist(err) {
		return err
	}

	serverFolder := Join(basePath, ServerFolderName)
	err = os.Mkdir(serverFolder, 0755)
	if os.IsExist(err) {
		return err
	}

	apiFolder := Join(basePath, APIFolderName)
	err = os.Mkdir(apiFolder, 0755)
	if os.IsExist(err) {
		return err
	}

	typesFolder := Join(basePath, TypesFolderName)
	err = os.Mkdir(typesFolder, 0755)
	if os.IsExist(err) {
		return err
	}

	return nil
}

// GenerateHelpers ...
func GenerateHelpers(config *ConfigFile) error {
	typesDefinitions, err := GetTypesDefinitions(config)
	if err != nil {
		return err
	}

	template, err := GetTemplate(HelperTemplate)
	if err != nil {
		return err
	}

	for _, typeDef := range typesDefinitions {
		filePath := fmt.Sprintf(
			"./%s/helpers/%s_helper.go",
			config.Project.Name,
			typeDef.TypePackage,
		)

		file, err := os.Create(filePath)
		if err != nil {
			if err == os.ErrExist {
				log.Println("Already exist")
			} else {
				panic(err)
			}

		}
		err = template.Execute(file, typeDef)
		if err != nil {
			panic(err)
		}

	}
	return nil
}
