package mapper_test

import (
	"testing"

	"github.com/PennState/mapper"
	"github.com/stretchr/testify/assert"
)

func TestMapper(t *testing.T) {
	tests := []struct {
		name string
		inp  string
	}{
		{"Pascal", "NoSql"},
		{"Camel", "noSql"},
		{"Snake", "no_sql"},
		{"Kebab", "no-sql"},
		{"Underscores", "__No__SQL__"},
		{"Dashes", "--No--SQL--"},
		{"Dashes", "--no--sql--"},
		{"Acronym", "NoSQL"},
		{"Illegal", "No$SQL"},
	}

	exp := "no_sql"

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, exp, mapper.Mapper(test.inp))
		})
	}
}
