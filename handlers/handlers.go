package handlers

import (
	"log"
	"strings"

	"githubc.com/asadbekGo/generate-code/pkg/helper"
)

var apiTexts string

func MakeHandlerss(sqlBody []byte) error {

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

	var templateProtoFilename = "./handlers/template.txt"
	templateProtoBody, err := helper.ReadFile(templateProtoFilename)
	if err != nil {
		log.Println("Error while ReadFile:", err.Error())
		return err
	}

	var (
		templateProto      = string(templateProtoBody)
		camelCaseText      = helper.SnakeToCamel(tableName)
		upperHeadTableName = strings.ToUpper(string(camelCaseText[0])) + camelCaseText[1:]
		tableNameTire      = strings.ReplaceAll(tableName, "_", "-")
	)

	templateProto = strings.ReplaceAll(templateProto, "Template", upperHeadTableName)
	templateProto = strings.ReplaceAll(templateProto, "/template", "/"+tableNameTire)
	templateProto = strings.ReplaceAll(templateProto, "template_id", tableName+"_id")
	templateProto = strings.ReplaceAll(templateProto, "template", camelCaseText)

	err = helper.WriteFile("./generates/handlers/"+tableName+".go", templateProto)
	if err != nil {
		log.Println("Error while WriteFile:", err.Error())
		return err
	}

	var apiFilename = "./handlers/api.txt"
	apiBody, err := helper.ReadFile(apiFilename)
	if err != nil {
		log.Println("Error while ReadFile:", err.Error())
		return err
	}
	var api = string(apiBody)

	api = strings.ReplaceAll(api, "Template", upperHeadTableName)
	api = strings.ReplaceAll(api, "/template", "/"+tableNameTire)
	api = strings.ReplaceAll(api, "template_id", tableName+"_id")
	apiTexts += api + "\n"

	return nil
}

func MakeApi() error {

	err := helper.WriteFile("./generates/handlers/api.go", apiTexts)
	if err != nil {
		log.Println("Error while WriteFile:", err.Error())
		return err
	}

	return nil
}
