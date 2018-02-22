package main

func main() {

	// templateData, err := ioutil.ReadFile("./template_type/helper_type.tmpl")
	// if err != nil {
	// 	panic(err)
	// }

	// t, err := template.New("simpleCRUD").Parse(string(templateData))
	// if err != nil {
	// 	panic(err)
	// }

	// err = os.Mkdir("helpers", 0777)
	// if err != nil {
	// 	if err == os.ErrExist {
	// 		log.Println("Already exist")
	// 	} else {
	// 		panic(err)
	// 	}
	// }

	// for _, typeDecl := range types {

	// 	file, err := os.Create(fmt.Sprintf("./helpers/%s_helper.go", typeDecl.TypePackage))
	// 	if err != nil {
	// 		if err == os.ErrExist {
	// 			log.Println("Already exist")
	// 		} else {
	// 			panic(err)
	// 		}

	// 	}

	// 	err = t.Execute(file, typeDecl)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// }

	err := GenerateFolderStruct("./myproject")
	if err != nil {
		panic(err)
	}
}
