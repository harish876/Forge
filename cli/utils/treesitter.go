package utils

import (
	"bufio"
	"context"
	"fmt"
	"math"
	"os"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/python"
)

var (
	IGNORE_CASE = "_"
	NEW_LINE    = "\n"
)

type QueryExecutionParams struct {
	Cursor     *sitter.QueryCursor
	Query      *sitter.Query
	Node       *sitter.Node
	SourceCode []byte
}

func NewQueryExecutionParams(cursor *sitter.QueryCursor, query *sitter.Query, node *sitter.Node, sourceCode []byte) *QueryExecutionParams {
	return &QueryExecutionParams{
		Cursor:     cursor,
		Query:      query,
		Node:       node,
		SourceCode: sourceCode,
	}
}

func GetQueryCursor(filePath string, query []byte) (*QueryExecutionParams, error) {
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

	sitterQuery, _ := sitter.NewQuery(query, lang)
	queryCursor := sitter.NewQueryCursor()

	return NewQueryExecutionParams(queryCursor, sitterQuery, node, sourceCode), nil
}

func GetCaseValues(filePath string, className string) ([]string, int, error) {
	rawQuery := fmt.Sprintf(`
	(
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
				   			) @matchStatement
						)
					)
				)
			)	 
		)
	)	
	`, className)

	q, err := GetQueryCursor(filePath, []byte(rawQuery))
	if err != nil {
		fmt.Println(err)
	}
	q.Cursor.Exec(q.Query, q.Node)
	var result []string
	rowCount := math.MaxInt

	for {
		m, ok := q.Cursor.NextMatch()
		if !ok {
			break
		}
		m = q.Cursor.FilterPredicates(m, q.SourceCode)
		for _, c := range m.Captures {
			if c.Node.Type() == "block" {
				rowCount = min(rowCount, int(c.Node.EndPoint().Row)) // get the first occurence
			}
			if c.Node.Type() == "case_pattern" {
				caseValue := c.Node.Content(q.SourceCode)
				result = append(result, strings.Trim(caseValue, "\""))
			}
		}
	}
	return result, max(0, rowCount-1), nil
}

func GetImportStatementEndRow(filePath string) (int, error) {
	rawQuery := `
	(
		module (
        	(class_definition) @classDef
        )
	)		
	`
	q, err := GetQueryCursor(filePath, []byte(rawQuery))
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	q.Cursor.Exec(q.Query, q.Node)
	result := math.MaxInt

	for {
		m, ok := q.Cursor.NextMatch()
		if !ok {
			break
		}
		m = q.Cursor.FilterPredicates(m, q.SourceCode)
		for _, c := range m.Captures {
			if c.Node.Type() == "class_definition" {
				result = min(result, int(c.Node.StartPoint().Row))
			}
		}
	}
	return max(0, result-1), nil
}

func InsertContentAtPosition(filePath string, row int, content []string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("File Open Error", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	idx := 0

	for scanner.Scan() {
		if idx == row {
			lines = append(lines, content...)
			idx += len(content)
		}
		lines = append(lines, scanner.Text())
		idx++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("File Scanning Error", err)
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println(err)
		return err
	}
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return nil
}
