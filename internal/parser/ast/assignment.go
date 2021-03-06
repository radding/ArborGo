package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// AssignmentNode is a node that represents an asignment operator
type AssignmentNode struct {
	AssignTo Node
	Value    Node
	Lexeme   lexer.Lexeme
}

// Accept visits the node
func (a *AssignmentNode) Accept(v Visitor) (Node, error) {
	return v.VisitAssignmentNode(a)
}

func (a *AssignmentNode) GetType() types.TypeNode {
	return a.AssignTo.GetType()
}
