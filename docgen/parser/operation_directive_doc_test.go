package parser

import (
	"go/ast"
	goparser "go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

func kthFixture(t *testing.T) string {
	t.Helper()
	candidates := []string{
		filepath.Join("..", "..", "..", "dsa", "linkedlist", "kthNodeFromEnd.go"), // dsa-pavilion/docgen/parser
		filepath.Join("..", "..", "dsa", "linkedlist", "kthNodeFromEnd.go"),       // legacy sibling docgen/parser
	}
	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	t.Skip("dsa fixture not found (expected ../dsa next to dsa-pavilion)")
	return ""
}

func TestKthOperationDocAttached(t *testing.T) {
	root := kthFixture(t)
	data, err := os.ReadFile(root)
	if err != nil {
		t.Skip("kth file not at", root, err)
	}
	fset := token.NewFileSet()
	n, err := goparser.ParseFile(fset, root, data, goparser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range n.Decls {
		fn, ok := d.(*ast.FuncDecl)
		if !ok || fn.Name == nil || fn.Name.Name != "KthFromEnd_BruteForce" {
			continue
		}
		if fn.Doc == nil {
			t.Fatal("expected FuncDecl.Doc before KthFromEnd_BruteForce")
		}
		desc, tim, sp := ParseOperationDirectiveDoc(fn.Doc)
		if desc == "" || tim == "" || sp == "" {
			t.Fatalf("ParseOperationDirectiveDoc: desc=%q time=%q space=%q", desc, tim, sp)
		}
		return
	}
	t.Fatal("func not found")
}

func TestExtractOperationCodeOrderedKthMeta(t *testing.T) {
	root := kthFixture(t)
	snips, err := ExtractOperationCodeOrdered(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(snips) < 1 {
		t.Fatal("expected snippets")
	}
	if snips[0].Description == "" || snips[0].Time == "" || snips[0].Space == "" {
		t.Fatalf("snippet0: %+v", snips[0])
	}
}
