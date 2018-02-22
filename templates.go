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

// GetTemplate ...
func GetTemplate(templateKind int, templatePath ...string) (*template.Template, error) {
	var templateData []byte
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
			break
		case APITemplate:
			templateData, err = ioutil.ReadFile("./templates/helper_type.gotemplate")
			if err != nil {
				return nil, err
			}
			break
		case TypeTemplate:
			templateData, err = ioutil.ReadFile("./templates/helper_type.gotemplate")
			if err != nil {
				return nil, err
			}
			break
		default:
			break
		}

	}
	var nameForTemplate string

	switch templateKind {
	case HelperTemplate:
		nameForTemplate = "helperTemplate"
	case APITemplate:
		nameForTemplate = "helperTemplate"
	case TypeTemplate:
		nameForTemplate = "helperTemplate"
	default:
		break
	}

	t, err := template.New(nameForTemplate).Parse(string(templateData))
	if err != nil {
		return nil, err
	}

	return t, nil
}
