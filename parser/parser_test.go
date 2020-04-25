package parser

import (
	"strings"
	"testing"
)

func TestExpressionStatement(t *testing.T) {
	source := strings.NewReader(`1`)
	parser := New(source)
	parser.Parse()
}
