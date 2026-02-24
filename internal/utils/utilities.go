package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

type FileWriter struct {
	BaseDir string
	Logger  *logrus.Logger
}

func NewFileWriter(baseDir string, log *logrus.Logger) *FileWriter {
	return &FileWriter{
		BaseDir: baseDir,
		Logger:  log,
	}
}

func (fw *FileWriter) SaveToFile(filename string, data []byte) error {
	fw.Logger.Debugf("saving file %s", filename)

	err := os.MkdirAll(fw.BaseDir, 0755)
	if err != nil {
		return fmt.Errorf("erro ao criar diretório: %w", err)
	}

	pathname := path.Join(fw.BaseDir, filename+".json")
	fw.Logger.Debugf("writing to %s (size %d)", pathname, len(data))
	err = os.WriteFile(pathname, data, 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar arquivo: %w", err)
	}

	return nil
}

func (fw *FileWriter) SavePrettyJSONToFile(filename string, data []byte) error {
	fw.Logger.Debugf("saving pretty JSON file %s", filename)

	if strings.Contains(filename, "/") {
		dir := path.Dir(filename)
		err := os.MkdirAll(path.Join(fw.BaseDir, dir), 0755)
		if err != nil {
			return fmt.Errorf("erro ao criar diretório: %w", err)
		}
	} else {
		err := os.MkdirAll(fw.BaseDir, 0755)
		if err != nil {
			return fmt.Errorf("erro ao criar diretório: %w", err)
		}
	}

	fw.Logger.Debugf("parsing JSON for %s", filename)
	var parsedJSON interface{}
	err := json.Unmarshal(data, &parsedJSON)
	if err != nil {
		return fmt.Errorf("erro ao parsear JSON: %w", err)
	}

	prettyData, err := json.MarshalIndent(parsedJSON, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao formatar JSON: %w", err)
	}

	pathname := path.Join(fw.BaseDir, filename+".json")
	fw.Logger.Debugf("writing to %s (size %d)", pathname, len(prettyData))
	err = os.WriteFile(pathname, prettyData, 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar arquivo: %w", err)
	}

	return nil
}
