package helper

import (
	"fmt"
	"regexp"
)

func ParseSQLQuery(query string) (tableName string, fields []string, err error) {
	// Regular expression patterns
	tablePattern := `CREATE\s+TABLE\s+IF\s+NOT\s+EXISTS\s+"?(\w+)"?\s*\(`
	fieldPattern := `"(\w+)"\s+(\w+(?:\(\d+\))?)(?:\s+(?:NOT\s+NULL|PRIMARY\s+KEY|DEFAULT\s+".*"))?`

	// Compile regular expressions
	tableRe := regexp.MustCompile(tablePattern)
	fieldRe := regexp.MustCompile(fieldPattern)

	// Find table name
	match := tableRe.FindStringSubmatch(query)
	if match == nil {
		return "", nil, fmt.Errorf("table name not found")
	}
	tableName = match[1]

	// Find fields and types
	fields = []string{}
	fieldMatches := fieldRe.FindAllStringSubmatch(query, -1)
	for _, fieldMatch := range fieldMatches {
		re := regexp.MustCompile(`\(\d+\)`)
		fieldMatch[2] = re.ReplaceAllString(fieldMatch[2], "")
		fields = append(fields, fieldMatch[1]+":"+fieldMatch[2])
	}

	return tableName, fields, nil
}
