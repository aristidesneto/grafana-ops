package internal

import (
	"grafana-ops/internal/client"
	"grafana-ops/internal/dashboards"
	"grafana-ops/internal/datasources"
	"grafana-ops/internal/folders"
	"grafana-ops/internal/utils"

	"github.com/sirupsen/logrus"
)

type SaveService struct {
	client *client.Client
	writer *utils.FileWriter
	logger *logrus.Logger
}

func NewSaveService(client *client.Client, writer *utils.FileWriter, log *logrus.Logger) *SaveService {
	return &SaveService{
		client: client,
		writer: writer,
		logger: log,
	}
}

func (ds *SaveService) Save() error {
	ds.logger.Info("starting full backup")

	backupFunctions := []struct {
		name string
		fn   func() error
	}{
		{"dashboards", dashboards.NewDashboardService(ds.client, ds.writer, ds.logger).SaveDashboards},
		{"datasources", datasources.NewDatasourceService(ds.client, ds.writer, ds.logger).SaveDatasources},
		{"folders", folders.NewFolderService(ds.client, ds.writer, ds.logger).SaveFolders},
	}

	for _, entry := range backupFunctions {
		ds.logger.Infof("running backup step: %s", entry.name)
		if err := entry.fn(); err != nil {
			return err
		}
	}

	ds.logger.Info("backup finished successfully")
	return nil
}
