package base

import (
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImplementsVisitor(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*ast.Visitor)(nil), new(Visitor))
}