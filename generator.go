package main

import "os"

// GenerateFolderStruct ...
func GenerateFolderStruct(basePath string) error {
	err := os.Mkdir(basePath, 0755)
	if os.IsExist(err) {
		return err
	}

	helpersFolder := Join(basePath, HelpersFolderName)
	err = os.Mkdir(helpersFolder, 0755)
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

	return nil
}
