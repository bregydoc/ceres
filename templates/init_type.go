package templatetype

import (
	"log"

	"github.com/asdine/storm"
)

// DbUser ...
var DbUser *storm.DB

func init() {
	var err error
	DbUser, err = storm.Open("")
	if err != nil {
		log.Fatalln("storm.Open()", err.Error())
	}
}
