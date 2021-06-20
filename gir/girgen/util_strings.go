package girgen

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	_ "embed"
)

//go:embed util_strings_specials.txt
var specialsTXT string

var (
	snakeRegex     = regexp.MustCompile(`[_0-9]+\w`)
	pascalSpecials = strings.Split(specialsTXT, "\n")
	pascalWords    = strings.NewReplacer(
		"Tolower", "ToLower",
		"Toupper", "ToUpper",
		"Totitle", "ToTitle",
		"Xdigit", "XDigit",
		"Dbus", "DBus",
		"Gicon", "GIcon",
		"Gtype", "GType",
		"Gvalue", "GValue",
		"Gvariant", "GVariant",
		"Hadj", "HAdj",
		"Vadj", "VAdj",
		"Hscroll", "HScroll",
		"Vscroll", "VScroll",
		"Nomem", "NOMEM",
		"Beos", "BeOS",
		"Directfb", "DirectFB",
		"Ipv4", "IPv4",
		"Ipv6", "IPv6",
		"ProXY", "Proxy",
	)
	pascalRegex *regexp.Regexp
)

func init() {
	fullRegex := strings.Builder{}
	fullRegex.Grow(256)
	fullRegex.WriteByte('(')

	for i, special := range pascalSpecials {
		if i > 0 {
			fullRegex.WriteByte('|')
		}
		fullRegex.WriteString(special)
	}

	fullRegex.WriteByte(')')

	// Must account for the next character being either EOF or a capitalized
	// letter to avoid cases like "IDentifier".
	fullRegex.WriteString("([A-Z0-9]|$)")

	pascalRegex = regexp.MustCompile(fullRegex.String())
}

// dots is a helper function to join strings in dots for debugging.
func dots(parts ...string) string {
	nonEmptyParts := parts[:0]

	for _, part := range parts {
		if part == "" {
			continue
		}

		if strings.Contains(part, "*") {
			part = "(" + part + ")"
		}

		nonEmptyParts = append(nonEmptyParts, part)
	}

	return strings.Join(nonEmptyParts, ".")
}

// PascalToGo converts regular Pascal case to Go.
func PascalToGo(pascal string) string {
	// Force constructors to have a New prefix instead of suffix.
	if strings.HasSuffix(pascal, "New") {
		pascal = "New" + strings.TrimSuffix(pascal, "New")
	}

	pascal = pascalRegex.ReplaceAllStringFunc(pascal, strings.ToUpper)
	pascal = pascalWords.Replace(pascal)

	return pascal
}

// FirstLetter returns the first letter in lower-case.
func FirstLetter(p string) string {
	r, sz := utf8.DecodeRuneInString(p)
	if sz > 0 && r != utf8.RuneError {
		return string(unicode.ToLower(r))
	}

	return string(p[0]) // fallback
}

// UnexportPascal converts the PascalToGo string to be unexported.
func UnexportPascal(pascal string) string {
	runes := []rune(pascal)
	if len(runes) < 1 {
		return snakeNoGo(strings.ToLower(pascal))
	}

	var i int
	for i < len(runes) && unicode.IsUpper(runes[i]) {
		i++
	}

	if i > 1 {
		i--
	}

	pascal = strings.ToLower(string(runes[:i])) + string(runes[i:])
	pascal = snakeNoGo(pascal)

	return pascal
}

// SnakeToGo converts snake case to Go's special case. If Pascal is true, then
// the first letter is capitalized.
func SnakeToGo(pascal bool, snakeString string) string {
	if pascal {
		snakeString = "_" + snakeString
	}

	snakeString = snakeRegex.ReplaceAllStringFunc(snakeString,
		func(orig string) string {
			orig = strings.ToUpper(orig)
			orig = strings.ReplaceAll(orig, "_", "")
			return orig
		},
	)

	if !pascal {
		return snakeNoGo(snakeString)
	}

	return PascalToGo(snakeString)
}

// goKeywords includes both Go types and keywords. This is primarily to prevent
// collisions with meaningful Go words.
var goKeywords = map[string]struct{}{
	// Keywords.
	"break":       {},
	"default":     {},
	"func":        {},
	"interface":   {},
	"select":      {},
	"case":        {},
	"defer":       {},
	"go":          {},
	"map":         {},
	"struct":      {},
	"chan":        {},
	"else":        {},
	"goto":        {},
	"package":     {},
	"switch":      {},
	"const":       {},
	"fallthrough": {},
	"if":          {},
	"range":       {},
	"type":        {},
	"continue":    {},
	"for":         {},
	"import":      {},
	"return":      {},
	"var":         {},

	// Types.
	"bool":       {},
	"byte":       {},
	"complex128": {},
	"complex64":  {},
	"error":      {},
	"float32":    {},
	"float64":    {},
	"int":        {},
	"int16":      {},
	"int32":      {},
	"int64":      {},
	"int8":       {},
	"rune":       {},
	"string":     {},
	"uint":       {},
	"uint16":     {},
	"uint32":     {},
	"uint64":     {},
	"uint8":      {},
	"uintptr":    {},
}

// cgoField formats the C field name to not be confused with a Go keyword.
// See https://golang.org/cmd/cgo/#hdr-Go_references_to_C.
func cgoField(field string) string {
	_, keyword := goKeywords[field]
	if keyword {
		return "_" + field
	}
	return field
}

// snakeNoGo ensures the snake-case string is never a Go keyword.
func snakeNoGo(snake string) string {
	switch snake {
	case "func":
		snake = "fn"
	case "type":
		snake = "typ"
	case "error":
		snake = "err"
	case "return":
		snake = "ret"
	}

	_, isKeyword := goKeywords[snake]
	if isKeyword {
		snake = "_" + snake
	}

	return snake
}
