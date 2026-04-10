package swiss_kit

import "testing"

func TestMysqlEscapeStringNoEscape(t *testing.T) {
	cases := []struct {
		input string
	}{
		{"hello"},
		{"abc123"},
		{""},
		{"select 1"},
		{"中文测试"},
	}
	for _, tc := range cases {
		got := MysqlEscapeString(tc.input)
		if got != tc.input {
			t.Errorf("MysqlEscapeString(%q) = %q, want %q (no escape needed)", tc.input, got, tc.input)
		}
	}
}

func TestMysqlEscapeStringSpecialChars(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		// null byte
		{"ab\x00cd", "'ab\\0cd'"},
		// newline
		{"line1\nline2", "'line1\\nline2'"},
		// carriage return
		{"line1\rline2", "'line1\\rline2'"},
		// backslash
		{"path\\to\\file", "'path\\\\to\\\\file'"},
		// single quote (escaped as '' per MySQL convention)
		{"it's", "'it''s'"},
		// double quote (passed through, not escaped)
		{`say "hello"`, `'say "hello"'`},
		// Ctrl-Z (0x1A)
		{"end\x1a", "'end\\Z'"},
	}
	for _, tc := range cases {
		got := MysqlEscapeString(tc.input)
		if got != tc.want {
			t.Errorf("MysqlEscapeString(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestMysqlEscapeStringSQLInjection(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		// classic SQL injection
		{"'; DROP TABLE users; --", "'''; DROP TABLE users; --'"},
		// single quote injection
		{"admin' OR '1'='1", "'admin'' OR ''1''=''1'"},
		// backslash + quote combo
		{"test\\' OR 1=1", "'test\\\\'' OR 1=1'"},
	}
	for _, tc := range cases {
		got := MysqlEscapeString(tc.input)
		if got != tc.want {
			t.Errorf("MysqlEscapeString(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestMysqlEscapeStringMultipleSpecial(t *testing.T) {
	// string with multiple special characters
	input := "hello\x00world\n\r\\'\"\x1a"
	got := MysqlEscapeString(input)
	want := "'hello\\0world\\n\\r\\\\''\"\\Z'"
	if got != want {
		t.Errorf("MysqlEscapeString(%q) = %q, want %q", input, got, want)
	}
}

func TestMysqlEscapeStringUnicode(t *testing.T) {
	// \u00a5 (Yen sign) and \u20a9 (Won sign) should pass through
	cases := []struct {
		input string
		want  string
	}{
		{"\u00a5", "\u00a5"},
		{"\u20a9", "\u20a9"},
		{"价格\u00a5100", "价格\u00a5100"},
	}
	for _, tc := range cases {
		got := MysqlEscapeString(tc.input)
		if got != tc.want {
			t.Errorf("MysqlEscapeString(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestMysqlIsEscapeNeeded(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"hello", false},
		{"abc\x00", true},
		{"abc\n", true},
		{"abc\r", true},
		{"abc\\", true},
		{"abc'", true},
		{`abc"`, true},
		{"abc\x1a", true},
		{"中文", false},
	}
	for _, tc := range cases {
		got := mysqlIsEscapeNeeded(tc.input)
		if got != tc.want {
			t.Errorf("mysqlIsEscapeNeeded(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
