package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"strings"
)

// GetTypesDefinitions ...
func GetTypesDefinitions(config *ConfigFile) ([]*TypeDefinition, error) {
	fset := token.NewFileSet()

	src, _ := ioutil.ReadFile(config.Types.Definitions)

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
				PluralInternalTypeName: strings.ToLower(Pluralize(ret.Name.Name)),
				TypeDbPath:             Join(config.Db.Path, strings.ToLower(ret.Name.Name)+".db"),
				InternalAPIName:        Snakeazer(ret.Name.Name),
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

				bufFieldType := bytes.NewBufferString("")
				printer.Fprint(bufFieldType, fset, f.Type)

				fields = append(fields, &TypeField{
					Name: f.Names[0].Name,
					Type: bufFieldType.String(),
					Tags: fmt.Sprintf("%s", tags),
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

// func GetImportsFromTypesFile(config *ConfigFile)
