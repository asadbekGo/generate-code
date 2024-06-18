package helper

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
		keys []string
	)

	for k := range params {
		keys = append(keys, k)
	}

	for _, k := range keys {
		if k != "" {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, params[k])
			i++
		}
	}

	return namedQuery, args
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

const otpChars = "1234567890"

func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func Difference(a, b []int32) []int32 {
	mb := make(map[int32]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int32
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func ValMultipleQuery(query string, vals []int32) (string, []interface{}) {
	params := []interface{}{}

	for i, id := range vals {
		query += fmt.Sprintf("$%d,", i+1)
		params = append(params, id)
	}

	query = query[:len(query)-1] // remove trailing ","
	query += ")"

	return query, params
}

func InsertMultiple(queryInsert string, id int32, vals []int32) (string, []interface{}) {
	insertparams := []interface{}{}

	for i, d := range vals {
		p1 := i * 2 // starting position for insert params
		queryInsert += fmt.Sprintf("($%d, $%d),", p1+1, p1+2)

		insertparams = append(insertparams, d, id)
	}

	queryInsert = queryInsert[:len(queryInsert)-1] // remove trailing ","

	return queryInsert, insertparams
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewNullBool(s bool) sql.NullBool {
	if !s {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  s,
		Valid: true,
	}
}

// preprocessMap converts unsupported types in the map to supported types
func preprocessMap(data map[string]interface{}) map[string]interface{} {
	processed := make(map[string]interface{})
	for k, v := range data {
		switch val := v.(type) {
		case time.Time:
			processed[k] = val.Format(time.RFC3339) // Convert time.Time to RFC3339 string
		case []byte:
			processed[k] = hex.EncodeToString(val) // Convert byte array to hex string
		default:
			processed[k] = v
		}
	}
	return processed
}
func MarshalToStruct(data interface{}, resp interface{}) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, resp)
	if err != nil {
		return err
	}

	return nil
}

func CheckTypeAndEmpty(value interface{}) bool {
	switch v := value.(type) {
	case string:
		return v == ""
	case int:
		return v == 0
	case bool:
		return !v
	case nil:
		return true
	default:
		return reflect.ValueOf(v).IsZero()
	}
}

// Function to create a map from a slice of structs based on a specified field
func CreateMapByField(slice interface{}, fieldName string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	v := reflect.ValueOf(slice)

	// Ensure the value is a slice
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected a slice")
	}

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		// Dereference if the item is a pointer
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		// Ensure the item is a struct
		if item.Kind() != reflect.Struct {
			return nil, fmt.Errorf("expected a slice of structs or pointers to structs")
		}

		fieldValue := item.FieldByName(fieldName)
		if !fieldValue.IsValid() {
			return nil, fmt.Errorf("field %s not found", fieldName)
		}

		key := fmt.Sprintf("%v", fieldValue.Interface())
		result[key] = item.Interface()
	}

	return result, nil
}

// placeholders generates a string of placeholders for a given number of fields
func Placeholders(numFields, offset int) string {
	holders := make([]string, numFields)
	for i := 0; i < numFields; i++ {
		holders[i] = fmt.Sprintf("$%d", offset+i+1)
	}
	return strings.Join(holders, ", ")
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
