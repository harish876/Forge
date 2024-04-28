package commands

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/harish876/forge/utils"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DATABASE          = "./history.db"
	JOBS_BASE_PATH    = "/Users/harishgokul/forge/jobs/"
	FACTORY_BASE_PATH = "/Users/harishgokul/forge/factory"
	TEST              = true
	DB_CLIENT         *sql.DB
)

func InitDbClient() (*sql.DB, error) {
	if DB_CLIENT != nil {
		return DB_CLIENT, nil
	}

	db, err := sql.Open("sqlite3", DATABASE)
	if err != nil {
		return nil, err
	}
	DB_CLIENT = db
	return DB_CLIENT, nil
}

type Step struct {
	StepType string // Type of Step  - Extractor, Loader, Transformer
	Dir      string
	StepName string // Name of the step, like extract_json
	Prefix   string //the prefix in front of the step file in the files and in code, like extract_json_job
}

func NewStep(stepType, stepName string) *Step {
	return &Step{
		StepType: stepType,
		StepName: stepName,
	}
}

func (s *Step) InitPrefixAndDirectory() {
	switch s.StepType {
	case "extract", "extractor":
		s.Prefix = "extract"
		s.Dir = "extractors"
	case "transform", "transformer":
		s.Prefix = "transformer"
		s.Dir = "transformers"
	case "load", "loader":
		s.Prefix = "loader"
		s.Dir = "loaders"
	case "report", "reporter":
		s.Prefix = "reporter"
		s.Dir = "reporters"
	default:
		s.Prefix = ""
		s.Dir = ""
	}
}

// return the step name in the format stepType_stepName_job
func (s *Step) GetformattedStepName() string {
	return fmt.Sprintf("%s_%s_job", s.Prefix, s.StepName)
}

func (s *Step) GetStepHistory() ([]Row, error) {
	db, err := sql.Open("sqlite3", DATABASE)
	if err != nil {
		slog.Error("Error reading existing steps", err)
	}
	defer db.Close()

	step := s.GetformattedStepName()

	err = s.SetupHistoryTable(db)

	if err != nil {
		slog.Error("Error creating Step", err, step)
	}

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", s.Dir))
	if err != nil {
		slog.Error("Error Reading from History Table", s.Dir, err)
	}
	defer rows.Close()

	var result []Row
	for rows.Next() {
		var row Row
		if err := rows.Scan(&row.Id, &row.Name); err != nil {
			slog.Error("Error Reading from History Table", s.Dir, err)
		}
		result = append(result, row)
	}

	slog.Info("SQLite History database setup successfully!")
	return result, nil
}

func (s *Step) SetupHistoryTable(db *sql.DB) error {
	_, err := db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)`, s.Dir))

	if err != nil {
		slog.Error("Error creating Step", err, s.GetformattedStepName())
		return err
	}
	return nil
}

func (s *Step) GetPythonJobCode() string {
	var pythonCode string
	caser := cases.Title(language.AmericanEnglish)
	pythonCode += "from jobs.job_interface import ETLJob\n\n"
	pythonCode += "class %s%sJob(ETLJob):\n"
	pythonCode += "\tdef __init__(self, config):\n"
	pythonCode += "\t\tsuper().__init__()\n\n"
	pythonCode += "\tdef execute(self, data=None):\n"
	pythonCode += "\t\tself.next()\n"

	return fmt.Sprintf(pythonCode, caser.String(s.Prefix), caser.String(s.StepName))
}

func (s *Step) GeneratePythonJobCode() {

	directory := fmt.Sprintf("%s/%s", JOBS_BASE_PATH, s.Dir)

	if err := os.MkdirAll(directory, 0755); err != nil {
		fmt.Println("Error:", err)
		return
	}

	filename := fmt.Sprintf("%s.py", s.GetformattedStepName())
	filePath := filepath.Join(directory, filename)
	if !TEST {
		if _, err := os.Stat(filePath); err == nil {
			fmt.Printf("File for %s already exists\n", filename)
			return
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	pythonText := s.GetPythonJobCode()
	_, err = file.WriteString(pythonText)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// Get the directory name of the job in the singular form
func (s *Step) GetFactoryCodeFileName() string {
	formattedName := s.Dir
	return formattedName[:len(formattedName)-1]
}

func (s *Step) GeneratePythonFactoryCode() {

	s.SetupHistoryTable(DB_CLIENT)
	fileNeedsCodeGen := s.InsertNewStep()

	if !fileNeedsCodeGen {
		fmt.Println("File does not any updation. All steps up to date")
		return
	}

	stepHistory, _ := s.GetStepHistory()
	fileContent := s.GetPythonFactoryCode(stepHistory)

	fileName := s.GetFactoryCodeFileName()
	filePath := filepath.Join(FACTORY_BASE_PATH, fmt.Sprintf("%s_factory.py", fileName))
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fileContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("File %s_factory.py created/updated successfully!\n", fileName)
}

func (s *Step) GetPythonFactoryCode(options []Row) string {
	var pythonCode string
	pythonCode += "from factory.factory_interface import Factory\n"

	for _, option := range options {
		pythonCode += fmt.Sprintf("from jobs.%s.%s import %s\n", s.Dir, option.Name, utils.SnakeToCamel(option.Name))
	}

	pythonCode += "\n"
	pythonCode += "class Factory(Factory):\n"
	pythonCode += "\tdef __init__(self):\n"
	pythonCode += "\t\tsuper().__init__()\n\n"
	pythonCode += "\tdef create(self, mode, **kwargs):\n"
	pythonCode += "\t\tmerged_config = self.get_config(mode)\n\n"
	pythonCode += "\t\t#Autogenerated File Section. Do not Edit this file\n\n"
	pythonCode += "\t\tmatch mode:\n"

	for _, option := range options {
		pythonCode += fmt.Sprintf("\t\t\tcase \"%s\":\n", option.Name)
		pythonCode += fmt.Sprintf("\t\t\t\treturn %s(config=merged_config)\n", utils.SnakeToCamel(option.Name))
	}
	pythonCode += "\t\t\tcase _:\n"
	pythonCode += "\t\t\t\traise ValueError(\"Invalid extract type\")\n"

	return pythonCode
}

func (s *Step) InsertNewStep() bool {
	formattedStepName := s.GetformattedStepName()
	stmt, err := DB_CLIENT.Prepare(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE name = ?", s.Dir))
	if err != nil {
		slog.Error("Error preparing statement:", err)
		s.InsertIntoHistoryTable(DB_CLIENT)
		return true
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(formattedStepName).Scan(&count)
	if err != nil {
		slog.Error("Error executing query:", err)
		return false
	}

	if count == 0 {
		s.InsertIntoHistoryTable(DB_CLIENT)
		return true
	}
	return false

}

func (s *Step) InsertIntoHistoryTable(db *sql.DB) {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO %s (name) VALUES (?)", s.Dir), s.GetformattedStepName())
	if err != nil {
		slog.Error("Error Inserting into ", s.Dir, err)
	}
}

type Row struct {
	Id   int
	Name string
}

func CreateStep(cmd *cobra.Command, args []string) {
	stepType, _ := cmd.Flags().GetString("type")
	stepName, _ := cmd.Flags().GetString("name")

	s := NewStep(stepType, stepName)
	s.InitPrefixAndDirectory()
	InitDbClient()

	s.GeneratePythonJobCode()
	s.GeneratePythonFactoryCode()

	fmt.Printf("Successfully Finished Action for %s step named %s\n", stepType, stepName)
}