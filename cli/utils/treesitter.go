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

func GetCaseValues(filePath string, className string) ([]string, error) {
	codeBuf := make([]byte, 10*1024*1024)
	file, _ := os.Open(filePath)
	n, err := file.Read(codeBuf)
	sourceCode := codeBuf[:n]

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	lang := python.GetLanguage()
	node, _ := sitter.ParseCtx(context.Background(), sourceCode, lang)

	rawQuery := fmt.Sprintf(`(
		module(
			class_definition
			name: ((identifier) @className (#match? @className %s))
		   		body: (block
					(function_definition 
						name: ((identifier) @functionName (#match? @functionName "create"))
						body: (block
							(match_statement 
								body: (block
								alternative: (
										case_clause(
											case_pattern (string)
					   				) @caseClause
					 			)
				   			)
						)
					)
				)
			)	 
		)
	)	
	`, className)

	query := []byte(rawQuery)

	q, _ := sitter.NewQuery(query, lang)
	qc := sitter.NewQueryCursor()
	qc.Exec(q, node)

	var result []string

	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}
		// Apply predicates filtering
		m = qc.FilterPredicates(m, sourceCode)
		for _, c := range m.Captures {
			if c.Node.Type() == "case_pattern" {
				caseValue := c.Node.Content(sourceCode)
				result = append(result, strings.Trim(caseValue, "\""))
			}
		}
	}

	return result, nil

}
