package storage

import (
	"fmt"
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

	fmt.Println("tableName:", tableName)
	fmt.Println("fields:", fields)

	return nil
}
