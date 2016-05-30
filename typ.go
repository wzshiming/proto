package proto

import "go/ast"

type pair struct {
	Key string
	Typ string
}

var m = map[*ast.Object]bool{}
