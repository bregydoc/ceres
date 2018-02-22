package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

// GetTypesDefinitions ...
func GetTypesDefinitions(configPath string) ([]*TypeDefinition, error) {
	config, err := GetConfigFromFile(configPath)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()

	src, _ := ioutil.ReadFile(config.Type.Definitions)

	f, err := parser.ParseFile(fset, "", src, parser.Trace)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var types []*TypeDefinition

	ast.Inspect(f, func(n ast.Node) bool {

		ret, ok := n.(*ast.TypeSpec)
		if ok {

			types = append(types, &TypeDefinition{
				TypeName:               ret.Name.Name,
				InternalTypeName:       strings.ToLower(ret.Name.Name),
				PluralTypeName:         Pluralize(ret.Name.Name),
				TypePackage:            strings.ToLower(ret.Name.Name),
				PluralInternalTypeName: Pluralize(ret.Name.Name),
			})

			return true
		}
		return true

	})

	count := 0

	ast.Inspect(f, func(n ast.Node) bool {
		// Find Return Statements
		ret, ok := n.(*ast.StructType)
		if ok {
			fields := make([]*TypeField, 0)

			for _, f := range ret.Fields.List {
				tags := ""
				if f.Tag != nil {
					tags = strings.Replace(f.Tag.Value, "`", "", -1)
				}

				fields = append(fields, &TypeField{
					Name: f.Names[0].Name,
					Type: fmt.Sprint(f.Type),
					Tags: tags,
				})
				//fmt.Println(f.Names[0], f.Type, strings.Replace(f.Tag.Value, "`", "", -1))
			}

			types[count].TypeFields = fields
			count++
			return true
		}
		return true
	})

	return types, nil
}
