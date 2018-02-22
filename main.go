package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	src, _ := ioutil.ReadFile("./types/types.go")

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "", src, parser.Trace)
	if err != nil {
		fmt.Println(err)
		return
	}

	var types []TypeDefinition

	ast.Inspect(f, func(n ast.Node) bool {

		ret, ok := n.(*ast.TypeSpec)
		if ok {

			types = append(types, TypeDefinition{
				TypeName:               ret.Name.Name,
				InternalTypeName:       strings.ToLower(ret.Name.Name),
				PluralTypeName:         ret.Name.Name + "s", // Pluralize(ret.Name.Name),
				TypePackage:            strings.ToLower(ret.Name.Name),
				PluralInternalTypeName: strings.ToLower(ret.Name.Name + "s"),
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
			fields := make([]TypeField, 0)

			for _, f := range ret.Fields.List {
				tags := ""
				if f.Tag != nil {
					tags = strings.Replace(f.Tag.Value, "`", "", -1)
				}

				fields = append(fields, TypeField{
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

	pp.Println(types)

	templateData, err := ioutil.ReadFile("./template_type/helper_type.tmpl")
	if err != nil {
		panic(err)
	}

	t, err := template.New("simpleCRUD").Parse(string(templateData))
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("helpers", 0777)
	if err != nil {
		if err == os.ErrExist {
			log.Println("Already exist")
		} else {
			panic(err)
		}
	}

	for _, typeDecl := range types {

		file, err := os.Create(fmt.Sprintf("./helpers/%s_helper.go", typeDecl.TypePackage))
		if err != nil {
			if err == os.ErrExist {
				log.Println("Already exist")
			} else {
				panic(err)
			}

		}

		err = t.Execute(file, typeDecl)
		if err != nil {
			panic(err)
		}

	}

}
