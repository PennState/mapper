package mapper_test

import (
	"testing"

	"github.com/PennState/mapper"
	"github.com/stretchr/testify/assert"
)

type test struct {
	name string
	inp  string
	exp  string
}

func tests() []test {
	return []test{
		{"Pascal", "PascalCase", "pascal_case"},
		{"Camel", "camelCase", "camel_case"},
		{"Snake", "snake_case", "snake_case"},
		{"Kebab", "kebab_case", "kebab_case"},
		{"Multiple underscores", "__Why__This__", "why_this"},
		{"Dashes", "--Why--This--", "why_this"},
		{"Dashes", "--why--this--", "why_this"},
		// {"Prefix acronym", "IRSWantsYou", "irs_wants_you"},
		// {"Infix acronym", "ExternalURLResolver", "external_url_resolver"},
		{"Postfix acronym", "NoSQL", "no_sql"},
		{"Mixed up mess", "Not_Pascal_Case", "not_pascal_case"},
		{"Illegal", "No$SQL", "no_sql"},
	}
}

func TestMapper(t *testing.T) {
	tests := tests()
	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.exp, mapper.Mapper(test.inp))
		})
	}
}

func BenchmarkMapper(b *testing.B) {
	tests := tests()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(tests); j++ {
			mapper.Mapper(tests[j].inp)
		}
	}
}
