package utils

import (
	"context"
	"fmt"
	"os"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/python"
)

var (
	IGNORE_CASE = "_"
)

func GetCaseValues(filePath string) ([]string, error) {
	codeBuf := make([]byte, 10*1024*1024)
	file, err := os.Open(filePath)
	n, err := file.Read(codeBuf)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	parser := sitter.NewParser()
	defer parser.Close()

	language := python.GetLanguage()
	parser.SetLanguage(language)

	tree, _ := parser.ParseCtx(context.Background(), nil, codeBuf[:n])
	rootNode := tree.RootNode()

	result := []string{}
	extractValues(rootNode, "case_clause", "case_pattern", codeBuf, &result)
	return result, nil
}

func extractValues(node *sitter.Node, rootType string, childType string, codeBuf []byte, collector *[]string) {
	if node.Type() == rootType {
		var casePatternValue string
		for i := 0; i < int(node.ChildCount()); i++ {
			child := node.Child(i)
			if child.Type() == childType {
				casePatternValue = child.Content(codeBuf)
			}
		}
		*collector = append(*collector, strings.Trim(casePatternValue, "\""))
	}

	for i := 0; i < int(node.ChildCount()); i++ {
		extractValues(node.Child(i), rootType, childType, codeBuf, collector)
	}
}
