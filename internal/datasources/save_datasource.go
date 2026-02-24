package datasources

import (
	"encoding/json"
	"fmt"
	"grafana-ops/internal/client"
	"grafana-ops/internal/utils"

	"github.com/sirupsen/logrus"
)

type Datasource struct {
	ID    int    `json:"id"`
	UID   string `json:"uid"`
	Name  string `json:"name"`
	OrgID int    `json:"orgid"`
}

type DatasourceService struct {
	client *client.Client
	writer *utils.FileWriter
	logger *logrus.Logger
}

func NewDatasourceService(client *client.Client, writer *utils.FileWriter, log *logrus.Logger) *DatasourceService {
	return &DatasourceService{
		client: client,
		writer: writer,
		logger: log,
	}
}

func (ds *DatasourceService) SaveDatasources() error {
	datasources, err := ds.listDatasources()
	if err != nil {
		return err
	}

	for _, datasource := range datasources {
		ds.logger.Debugf("saving datasource uid=%s", datasource.UID)
		err = ds.saveById(datasource.UID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ds *DatasourceService) listDatasources() ([]Datasource, error) {
	url := "/api/datasources"
	ds.logger.Debugf("calling Grafana API %s", url)
	dts, err := ds.client.Get(url)
	if err != nil {
		return nil, err
	}

	var datasources []Datasource
	err = json.Unmarshal(dts, &datasources)
	if err != nil {
		return nil, err
	}

	return datasources, nil
}

func (ds *DatasourceService) saveById(datasourceUid string) error {
	url := fmt.Sprintf("/api/datasources/uid/%s", datasourceUid)
	ds.logger.Debugf("calling Grafana API %s", url)
	data, err := ds.client.Get(url)
	if err != nil {
		return err
	}
	ds.logger.Debugf("writing datasource %s to disk", datasourceUid)
	return ds.writer.SavePrettyJSONToFile("datasources/"+datasourceUid, data)
}
