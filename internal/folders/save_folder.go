package folders

import (
	"grafana-ops/internal/client"
	"grafana-ops/internal/utils"

	"github.com/sirupsen/logrus"
)

type FolderService struct {
	client *client.Client
	writer *utils.FileWriter
	logger *logrus.Logger
}

func NewFolderService(client *client.Client, writer *utils.FileWriter, log *logrus.Logger) *FolderService {
	return &FolderService{
		client: client,
		writer: writer,
		logger: log,
	}
}

func (fs *FolderService) SaveFolders() error {
	url := "/api/folders"
	fs.logger.Debugf("calling Grafana API %s", url)
	data, err := fs.client.Get(url)
	if err != nil {
		return err
	}
	fs.logger.Debugf("writing folders list to disk")
	err = fs.writer.SavePrettyJSONToFile("folders", data)
	if err != nil {
		return err
	}

	return nil
}
