package tree_sitter_solidity_test

import (
	"testing"

	solidity "github.com/Artifex1/tree-sitter-solidity/bindings/go"
	tree_sitter "github.com/smacker/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(solidity.Language())
	if language == nil {
		t.Errorf("Error loading Solidity grammar")
	}
}
