package tree_sitter_solidity

// #cgo CFLAGS: -std=c11 -fPIC -I./src -I../../src
// #include "parser.c"
import "C"

import "unsafe"

// Language returns the tree-sitter language for this grammar.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_solidity())
}
