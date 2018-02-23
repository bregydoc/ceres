package main

import (
	"html/template"
	"io/ioutil"
)

// HelperTemplate ...
const HelperTemplate = 0

// APITemplate ...
const APITemplate = 1

// TypeTemplate ...
const TypeTemplate = 2

// InitDbTemplate ...
const InitDbTemplate = 3

// GetTemplate ...
func GetTemplate(templateKind int, templatePath ...string) (*template.Template, error) {
	var templateData []byte
	var nameForTemplate string

	if len(templatePath) > 0 {
		templatePathPath := templatePath[0]

		var err error
		templateData, err = ioutil.ReadFile(templatePathPath)
		if err != nil {
			return nil, err
		}

	} else {
		var err error
		switch templateKind {
		case HelperTemplate:
			templateData, err = ioutil.ReadFile("./templates/helper_type.gotemplate")
			if err != nil {
				return nil, err
			}
			nameForTemplate = "helperTemplate"
			break
		case APITemplate:
			templateData, err = ioutil.ReadFile("./templates/helper_type.gotemplate")
			if err != nil {
				return nil, err
			}
			nameForTemplate = "helperTemplate"
			break
		case TypeTemplate:
			templateData, err = ioutil.ReadFile("./templates/helper_type.gotemplate")
			if err != nil {
				return nil, err
			}
			nameForTemplate = "helperTemplate"
			break
		case InitDbTemplate:
			templateData, err = ioutil.ReadFile("./templates/init_type.gotemplate")
			if err != nil {
				return nil, err
			}
			nameForTemplate = "dbLinkTemplate"
			break
		default:
			break
		}

	}

	t, err := template.New(nameForTemplate).Parse(string(templateData))
	if err != nil {
		return nil, err
	}

	return t, nil
}
