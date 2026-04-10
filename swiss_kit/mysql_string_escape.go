package swiss_kit

import "strings"

// MysqlEscapeString escapes a string for safe use in MySQL SQL statements.
// It wraps the result in single quotes and escapes special characters
// to prevent SQL injection. Ported from MySQL JDBC driver.
func MysqlEscapeString(x string) string {
	if !mysqlIsEscapeNeeded(x) {
		return x
	}

	var buf strings.Builder
	buf.Grow(len(x) + len(x)/10 + 2)
	buf.WriteByte('\'')

	for _, c := range x {
		switch c {
		case 0:
			buf.WriteString("\\0")
		case '\n':
			buf.WriteString("\\n")
		case '\r':
			buf.WriteString("\\r")
		case '\\':
			buf.WriteString("\\\\")
		case '\'':
			buf.WriteString("''")
		case '"':
			buf.WriteByte('"')
		case '\032':
			buf.WriteString("\\Z")
		default:
			buf.WriteRune(c)
		}
	}

	buf.WriteByte('\'')
	return buf.String()
}

func mysqlIsEscapeNeeded(x string) bool {
	for _, c := range x {
		switch c {
		case 0, '\n', '\r', '\\', '\'', '"', '\032':
			return true
		}
	}
	return false
}
