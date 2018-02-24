package main

import "os"

// CheckAndFixDependencies ...
func CheckAndFixDependencies(config *ConfigFile) error {

	// Repair Db path
	err := os.Mkdir(config.Db.Path, 0755)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	return nil
}
