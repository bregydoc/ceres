package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
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
				return err
			}

		}

		tempData := bytes.NewBufferString("")

		err = template.Execute(tempData, typeDef)
		if err != nil {
			return err
		}

		rawData := tempData.String()

		finalStringData := strings.Replace(string(rawData), "&#34;", "\"", -1)

		_, err = file.Write([]byte(finalStringData))
		if err != nil {
			return err
		}

	}
	return nil
}

// GenerateDbLinks ...
func GenerateDbLinks(config *ConfigFile) error {
	types, err := GetTypesDefinitions(config)
	forTemplate := map[string][]*TypeDefinition{
		"Types": types,
	}

	template, err := GetTemplate(InitDbTemplate)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("./%s/helpers/init_db.go", config.Project.Name)

	file, err := os.Create(filePath)
	if err != nil {
		if err == os.ErrExist {
			log.Println("Already exist")
		} else {
			return err
		}

	}
	err = template.Execute(file, forTemplate)
	if err != nil {
		return err
	}
	return nil
}
