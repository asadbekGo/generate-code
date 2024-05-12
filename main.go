package main

import (
	"log"
	"strings"

	"githubc.com/asadbekGo/generate-code/pkg/helper"
	"githubc.com/asadbekGo/generate-code/protos"
)

func main() {

	body, err := helper.ReadFile("./sql/template.sql")
	if err != nil {
		log.Println("Error while read file:", err.Error())
		return
	}

	tables := strings.Split(string(body), ";")
	for _, table := range tables {
		if len(table) > 1 {
			err = protos.MakeProtos([]byte(table))
			if err != nil {
				log.Println("Error while read file:", err.Error())
				return
			}

			// err = storage.MakeStorage([]byte(table))
			// if err != nil {
			// 	log.Println("Error while read file:", err.Error())
			// 	return
			// }
		}
	}
}
