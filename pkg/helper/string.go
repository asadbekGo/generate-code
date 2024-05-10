package helper

import (
	"strings"
)

func RemoveEmptyRows(input string) string {
	// Split the input string into lines
	lines := strings.Split(input, "\n")

	// Filter out the empty lines
	var nonEmptyLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	// Join the non-empty lines back together
	output := strings.Join(nonEmptyLines, "\n")

	return output
}

func Pluralize(word string) string {
	// Check for common pluralization rules
	if strings.HasSuffix(word, "y") {
		return word[:len(word)-1] + "ies"
	}

	// Add more pluralization rules as needed

	// Default pluralization: add "s" suffix
	return word + "s"
}

func SQLToGoType(sqlType string) string {
	// Convert PostgreSQL types to Go types
	sqlType = strings.ToLower(sqlType)
	switch sqlType {
	case "uuid":
		return "string"
	case "varchar", "text", "char", "character varying":
		return "string"
	case "integer", "int", "smallint", "number":
		return "int32"
	case "bigint":
		return "int64"
	case "float", "double precision", "numeric", "decimal", "real":
		return "double"
	case "boolean":
		return "bool"
	case "date", "timestamp", "time":
		return "string"
	case "json", "jsonb":
		return "interface{}" // Handle JSON types
	default:
		return "interface{}" // Fallback to interface{} for unknown types
	}
}

func SnakeToCamel(s string) string {
	// Split the string by underscores
	parts := strings.Split(s, "_")

	// Iterate through the parts and capitalize the first letter of each part except the first one
	for i := 1; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}

	// Join the parts back together
	return strings.Join(parts, "")
}
