package storage

import (
	"fmt"
	"log"
	"strings"

	"githubc.com/asadbekGo/generate-code/pkg/helper"
)

func MakeService(sqlBody []byte) error {

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
		fields[index] = fieldType[0]
	}

	var templateGoFilename = "./storage/template_storage.txt"
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

	var query = generateQuery(tableName, fields)
	templateGo = strings.ReplaceAll(templateGo, "insertQuery", query.InsertQuery)
	templateGo = strings.ReplaceAll(templateGo, "insertExecField", query.InsertExecField)
	templateGo = strings.ReplaceAll(templateGo, "getQuery", query.GetQuery)
	templateGo = strings.ReplaceAll(templateGo, "updateQuery", query.UpdateQuery)
	templateGo = strings.ReplaceAll(templateGo, "updateExecQuery", query.UpdateExecQuery)
	templateGo = strings.ReplaceAll(templateGo, "varNullString", query.VarNullString)
	templateGo = strings.ReplaceAll(templateGo, "varScan", query.VarScan)
	templateGo = strings.ReplaceAll(templateGo, "responseStruct", query.ResponseStruct)

	err = helper.WriteFile("generates/storage/"+tableName+".go", templateGo)
	if err != nil {
		log.Println("Error while WriteFile:", err.Error())
		return err
	}

	return nil
}

func generateQuery(tableName string, fields []string) Query {
	insertQuery := fmt.Sprintf(`INSERT INTO "%s" (`, tableName) + "\n"
	insertValueQuery := "\tVALUES ("
	var insertExecField, getQuery, updateQuery, updateExecQuery, varNullString, varScan, responseStruct string
	for ind, field := range fields {

		if field == "created_at" {
			continue
		}

		var text = fmt.Sprintf("\t\t\t%s,\n", field)
		if field == "updated_at" {
			text = "\t\t\tupdated_at\n"
		}
		insertQuery += text

		if field != "updated_at" {
			getQuery += "\t\t\t" + field + ",\n"
			insertValueQuery += fmt.Sprintf("$%d, ", ind+1)
		}

		var unsupportedField = []string{"id", "updated_at"}
		if !helper.Contains(unsupportedField, field) {
			var (
				fieldCamelCase = helper.SnakeToCamel(field)
				fieldUpperHead = strings.ToUpper(string(fieldCamelCase[0])) + fieldCamelCase[1:]
			)

			insertExecField += fmt.Sprintf("\t\treq.Get%s(),\n", fieldUpperHead)
			updateQuery += fmt.Sprintf("\t\t\t%s = :%s,\n", field, field)
			updateExecQuery += "\t\t" + fmt.Sprintf(`"%s": req.Get%s(),`, field, fieldUpperHead) + "\n"
			varNullString += fmt.Sprintf("\t%s sql.NullString\n", fieldCamelCase)
			varScan += fmt.Sprintf("\t\t&%s,\n", fieldCamelCase)
			responseStruct += fmt.Sprintf("\t\t%s: %s.String,\n", fieldUpperHead, fieldCamelCase)
		}
	}
	insertQuery += "\t\t)\n" + "\t" + insertValueQuery + "now())"
	insertExecField = insertExecField[:len(insertExecField)-1]
	getQuery = getQuery[:len(getQuery)-1]
	updateQuery = updateQuery[:len(updateQuery)-1]
	updateExecQuery = updateExecQuery[:len(updateExecQuery)-1]
	varNullString = varNullString[:len(varNullString)-1]
	varScan = varScan[:len(varScan)-1]
	responseStruct = responseStruct[:len(responseStruct)-1]

	return Query{
		InsertQuery:     insertQuery,
		InsertExecField: insertExecField,
		GetQuery:        getQuery,
		UpdateQuery:     updateQuery,
		UpdateExecQuery: updateExecQuery,
		VarNullString:   varNullString,
		VarScan:         varScan,
		ResponseStruct:  responseStruct,
	}
}

type Query struct {
	InsertQuery     string
	InsertExecField string
	GetQuery        string
	UpdateQuery     string
	UpdateExecQuery string
	VarNullString   string
	VarScan         string
	ResponseStruct  string
}
