package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/harish876/forge/cli/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateStepCodeParser(t *testing.T) {
	currDir, _ := os.Getwd()
	parentDir := filepath.Dir(currDir)
	filePath := filepath.Join(parentDir, "code", "example_extract.py")
	result, _ := utils.GetCaseValues(filePath, "ExtractorFactory")
	expected := []string{"extract_json", "extract_harish", "extract_girish"}
	assert.ElementsMatchf(t, expected, result, "Result Does not match")
	fmt.Println("Debug", result)
}
