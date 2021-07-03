package fixtures

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"errors"

	"github.com/JoaoLeal92/goals_backend/repositories"
)

// PrepareDatabase cleans up database for testing
func PrepareDatabase(conn *repositories.Connection, sqlFilePath string) (err error) {
	requests, err := readSQLFile(sqlFilePath)
	if err != nil {
		return err
	}

	for _, request := range requests {
		result := conn.Db.Exec(request)

		if result.Error != nil {
			return errors.New(result.Error.Error())
		}
	}

	return nil
}

func readSQLFile(filePath string) (requests []string, err error) {
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return requests, err
	}

	file, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		return requests, err
	}

	requests = strings.Split(string(file), ";\n")

	return requests, nil
}
