package rulesv2

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

var and = `&& 2`
var or = `|| 3`
var varnameAnd = `xyz && abc`
var varnameOr = `xyz || abc`
var numberOr = `1 || 2`
var numberAnd = `1 && 2`

func TestCanParseABooleanExpression(t *testing.T) {
	varName := &ast.VarName{}
	assert := assert.New(t)
	p := parseTest(and)
	boolean, err := boolOperation(varName, p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	boolNode := boolean.(*ast.BoolOp)
	if !assert.NotNil(boolNode, "failed to convert to BoolOp") {
		t.Fatal()
	}
	rs := boolNode.RightSide.(*ast.Constant)
	if !assert.NotNil(rs, "failed to convert to Constant") {
		t.Fatal()
	}

	p = parseTest(or)
	boolean, err = boolOperation(varName, p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	boolNode = boolean.(*ast.BoolOp)
	if !assert.NotNil(boolNode, "failed to convert to BoolOp") {
		t.Fatal()
	}
	rs = boolNode.RightSide.(*ast.Constant)
	if !assert.NotNil(rs, "failed to convert to Constant") {
		t.Fatal()
	}
}

func TestCanParseFromExpression(t *testing.T) {
	assert := assert.New(t)

	p := parseTest(varnameAnd)
	expr, err := ExpressionRule(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	assert.NotNil(expr)
	exprNode, _ := expr.(*ast.BoolOp)
	if !assert.NotNil(exprNode, "couldn't convert to BoolOP") {
		t.Fatal()
	}
	ls := exprNode.LeftSide.(*ast.VarName)
	if !assert.NotNil(ls) {
		t.Fatal()
	}
	rs := exprNode.RightSide.(*ast.VarName)
	if !assert.NotNil(rs) {
		t.Fatal()
	}
	assert.Equal("and", exprNode.Condition)

	p = parseTest(varnameOr)
	expr, err = ExpressionRule(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	assert.NotNil(expr)
	exprNode, _ = expr.(*ast.BoolOp)
	if !assert.NotNil(exprNode, "couldn't convert to BoolOP") {
		t.Fatal()
	}
	ls = exprNode.LeftSide.(*ast.VarName)
	if !assert.NotNil(ls) {
		t.Fatal()
	}
	rs = exprNode.RightSide.(*ast.VarName)
	if !assert.NotNil(rs) {
		t.Fatal()
	}
	assert.Equal("or", exprNode.Condition)

	p = parseTest(numberOr)
	expr, err = ExpressionRule(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	assert.NotNil(expr)
	exprNode, _ = expr.(*ast.BoolOp)
	if !assert.NotNil(exprNode, "couldn't convert to BoolOP") {
		t.Fatal()
	}
	lsC := exprNode.LeftSide.(*ast.Constant)
	if !assert.NotNil(lsC) {
		t.Fatal()
	}
	rsC := exprNode.RightSide.(*ast.Constant)
	if !assert.NotNil(rsC) {
		t.Fatal()
	}
	assert.Equal("or", exprNode.Condition)
}
