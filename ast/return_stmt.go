package ast

import (
	"bytes"

	"github.com/fcidade/monkey-lang/token"
)

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

var _ Statement = &ReturnStatement{}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}
