package storage

import (
	"log"
	"strings"

	"githubc.com/asadbekGo/generate-code/pkg/helper"
)

func MakeStorage(sqlBody []byte) error {

	var sqlTable = helper.RemoveEmptyRows(string(sqlBody))
	if len(sqlTable) <= 0 {
		return nil
	}

	tableName, fields, err := helper.ParseSQLQuery(sqlTable)
	if err != nil {
		log.Println("Error while ParseSQLQuery:", err.Error())
		return err
	}

	for index, field := range fields {
		var fieldType = strings.Split(field, ":")
		fields[index] = fieldType[0] + ":" + helper.SQLToGoType(fieldType[1])
	}

	var templateGoFilename = "./storage/template_service.txt"
	templateGoBody, err := helper.ReadFile(templateGoFilename)
	if err != nil {
		log.Println("Error while ReadFile:", err.Error())
		return err
	}

	var (
		templateGo         = string(templateGoBody)
		camelCaseText      = helper.SnakeToCamel(tableName)
		upperHeadTableName = strings.ToUpper(string(camelCaseText[0])) + camelCaseText[1:]
	)

	templateGo = strings.ReplaceAll(templateGo, "Template", upperHeadTableName)
	templateGo = strings.ReplaceAll(templateGo, "template", tableName)

	err = helper.WriteFile("generates/service/"+tableName+".go", templateGo)
	if err != nil {
		log.Println("Error while WriteFile:", err.Error())
		return err
	}

	return nil
}
