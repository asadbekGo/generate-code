package protos

import (
	"fmt"
	"log"
	"strings"

	"githubc.com/asadbekGo/generate-code/pkg/helper"
)

func MakeProtos(sqlBody []byte) error {

	var sqlTable = helper.RemoveEmptyRows(string(sqlBody))
	tableName, fields, err := helper.ParseSQLQuery(sqlTable)
	if err != nil {
		log.Println("Error while ParseSQLQuery:", err.Error())
		return err
	}

	for index, field := range fields {
		var fieldType = strings.Split(field, ":")
		fields[index] = fieldType[0] + ":" + helper.SQLToGoType(fieldType[1])
	}

	var templateProtoFilename = "./protos/template.proto"
	templateProtoBody, err := helper.ReadFile(templateProtoFilename)
	if err != nil {
		log.Println("Error while ReadFile:", err.Error())
		return err
	}

	var (
		templateProto      = string(templateProtoBody)
		upperHeadTableName = strings.ToUpper(string(tableName[0])) + tableName[1:]
	)

	templateProto = strings.ReplaceAll(templateProto, "message Template {}", GenerateProtoMessage(upperHeadTableName, fields))

	var (
		createFields []string
		updateFields []string
	)
	for _, field := range fields {
		if !strings.Contains(field, "id:string") && !strings.Contains(field, "created_at:string") && !strings.Contains(field, "updated_at:string") {
			createFields = append(createFields, field)
		}

		if !strings.Contains(field, "created_at:string") && !strings.Contains(field, "updated_at:string") {
			updateFields = append(updateFields, field)
		}
	}

	templateProto = strings.ReplaceAll(templateProto, "message CreateTemplateRequest {}", GenerateProtoMessage(fmt.Sprintf("Create%sRequest", upperHeadTableName), createFields))
	templateProto = strings.ReplaceAll(templateProto, "message UpdateTemplateRequest {}", GenerateProtoMessage(fmt.Sprintf("Update%sRequest", upperHeadTableName), updateFields))
	templateProto = strings.ReplaceAll(templateProto, "Template", upperHeadTableName)
	templateProto = strings.ReplaceAll(templateProto, "templates", helper.Pluralize(tableName))

	err = helper.WriteFile("./generates/"+tableName+".proto", templateProto)
	if err != nil {
		log.Println("Error while WriteFile:", err.Error())
		return err
	}

	return nil
}

func GenerateProtoMessage(messageName string, fields []string) string {
	var text = fmt.Sprintf("message %s {\n", messageName)
	for fieldIndex, field := range fields {
		parts := strings.Split(strings.TrimSpace(field), ":")
		if len(parts) == 2 {
			fieldName := strings.TrimSpace(parts[0])
			fieldType := strings.TrimSpace(parts[1])

			if fieldType == "interface{}" {
				fieldType = "string"
			}

			text += fmt.Sprintf("    %s %s = %d;\n", fieldType, fieldName, fieldIndex+1)
		}
	}
	text += "}"

	return text
}
