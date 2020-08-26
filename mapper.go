package mapper

import (
	"strings"
)

/*
mapper provides an alternative to the default sqlx field name mapper
(which is simply strings.ToLower()).  This mapper converts Pascal-cased,
camel-cased, snake-cased and kebab-cased field names into lower-cased,
snake-cased names.
*/
func Mapper(name string) string {
	const alphaOffset = 0x20

	sb := strings.Builder{}
	lastSep := true
	lastUpper := true

	for i := 0; i < len(name); i++ {
		cin := name[i]
		cout := cin | alphaOffset

		if lastSep && isSeparator(cin) {
			lastUpper = false

			continue
		}

		if !lastSep && isSeparator(cin) {
			sb.WriteByte('_')

			lastSep = true
			lastUpper = false

			continue
		}

		if !isAlpha(cout) {
			continue
		}

		if !lastUpper && !lastSep && isUpper(cin) {
			sb.WriteByte('_')
		}

		lastSep = false
		lastUpper = isUpper(cin)

		sb.WriteByte(cout)
	}

	return strings.Trim(sb.String(), "_-")
}

func isAlpha(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isSeparator(c byte) bool {
	return c == '-' || c == '_'
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}
