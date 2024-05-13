package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/harish876/forge/cli/utils"
)

func TestCreateStepCodeParser(t *testing.T) {
	currDir, _ := os.Getwd()
	parentDir := filepath.Dir(currDir)
	filePath := filepath.Join(parentDir, "code", "example_extract.py")
	result, _, _ := utils.GetCaseValues(filePath, "ExtractorFactory")

	lines := []string{
		fmt.Sprintf("from jobs.extractors.extract_json_job import ExtractJsonJob"),
	}
	idx, _ := utils.GetImportStatementEndRow(filePath)
	fmt.Println("Insert Posisition: ", idx)
	utils.InsertContentAtPosition(filePath, idx, lines)
	_, nidx, _ := utils.GetCaseValues(filePath, "ExtractorFactory")

	lines = []string{
		fmt.Sprintf("\t\t\tcase \"%s\":", "extract_json"),
		fmt.Sprintf("\t\t\t\treturn %s(config=merged_config)", utils.SnakeToCamel("extract_json"+"_job")),
	}
	utils.InsertContentAtPosition(filePath, nidx, lines)
	fmt.Println("Debug", result)
}

