package tree_sitter_solidity

// #cgo CFLAGS: -std=c11 -fPIC -I./src
// #include "parser.c"
// // NOTE: if your language has an external scanner, add it here.
import "C"

import "unsafe"

// Language returns the tree-sitter language for this grammar.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_solidity())
}
