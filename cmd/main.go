package main

import (
	"log"
	"strings"

	"githubc.com/asadbekGo/generate-code/handlers"
	"githubc.com/asadbekGo/generate-code/pkg/helper"
	"githubc.com/asadbekGo/generate-code/storage"
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
			// err = handlers.MakeHandlerss([]byte(table))
			// if err != nil {
			// 	log.Println("Error while MakeHandlerss:", err.Error())
			// 	return
			// }

			// err = protos.MakeProtos([]byte(table))
			// if err != nil {
			// 	log.Println("Error while read file:", err.Error())
			// 	return
			// }

			// err = storage.MakeService([]byte(table))
			// if err != nil {
			// 	log.Println("Error while read file:", err.Error())
			// 	return
			// }

			err = storage.MakeStorage([]byte(table))
			if err != nil {
				log.Println("Error while read file:", err.Error())
				return
			}
		}
	}

	err = handlers.MakeApi()
	if err != nil {
		log.Println("Error while MakeApi:", err.Error())
		return
	}

}
