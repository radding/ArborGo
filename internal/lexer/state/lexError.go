package state

import (
	"github.com/arborlang/ArborGo/internal/lexer/internal"
)

func lexError(msg string) State {
	return func(lex *internal.Lexer) State {
		lex.Errorf(msg)
		return nil
	}
}
