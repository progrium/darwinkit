package declparse

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestParser(t *testing.T) {
	deep.NilMapsAreEmpty = true

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			p := NewStringParser(normalizeStmntString(tt.s))
			p.Hint = tt.Hint
			got, err := p.Parse()
			if err != nil {
				t.Fatal("parse:", err)
			}
			if diff := deep.Equal(got, normalizeStmntNode(tt.n)); diff != nil {
				t.Error("diff:", diff)
			}
		})
	}
}

// easier to make everything a statement

func normalizeStmntString(s string) string {
	return strings.TrimRight(s, ";") + ";"
}

func normalizeStmntNode(n Node) *Statement {
	stmt := &Statement{}
	switch v := n.(type) {
	case *Statement:
		return v
	case *InterfaceDecl:
		stmt.Interface = v
	case *PropertyDecl:
		stmt.Property = v
	case *MethodDecl:
		stmt.Method = v
	case *VariableDecl:
		stmt.Variable = v
	case *ProtocolDecl:
		stmt.Protocol = v
	case *FunctionDecl:
		stmt.Function = v
	default:
		log.Fatal("bad node type:", reflect.TypeOf(n))
	}
	return stmt
}
