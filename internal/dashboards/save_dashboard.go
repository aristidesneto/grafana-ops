package dashboards

import (
	"encoding/json"
	"fmt"
	"grafana-ops/internal/client"
	"grafana-ops/internal/utils"

	"github.com/sirupsen/logrus"
)

type Dashboard struct {
	ID    int      `json:"id"`
	UID   string   `json:"uid"`
	Title string   `json:"title"`
	URI   string   `json:"uri"`
	Type  string   `json:"type"`
	Tags  []string `json:"tags"`
}

type DashboardService struct {
	client *client.Client
	writer *utils.FileWriter
	logger *logrus.Logger
}

func NewDashboardService(client *client.Client, writer *utils.FileWriter, log *logrus.Logger) *DashboardService {
	return &DashboardService{
		client: client,
		writer: writer,
		logger: log,
	}
}

func (ds *DashboardService) SaveDashboards() error {
	dashs, err := ds.listDashboards()
	if err != nil {
		return err
	}

	for _, dashboard := range dashs {
		ds.logger.Debugf("saving dashboard uid=%s title=%s", dashboard.UID, dashboard.Title)
		err = ds.saveById(dashboard.UID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ds *DashboardService) saveById(dashboardUID string) error {
	url := fmt.Sprintf("/api/dashboards/uid/%s", dashboardUID)
	ds.logger.Debugf("calling Grafana API %s", url)
	data, err := ds.client.Get(url)
	if err != nil {
		return err
	}
	ds.logger.Debugf("writing dashboard %s to disk", dashboardUID)
	return ds.writer.SavePrettyJSONToFile("dashboards/"+dashboardUID, data)
}

func (ds *DashboardService) listDashboards() ([]Dashboard, error) {
	url := "/api/search?type=dash-db"
	ds.logger.Debugf("calling Grafana API %s", url)
	data, err := ds.client.Get(url)
	if err != nil {
		return nil, err
	}

	var dashboards []Dashboard
	err = json.Unmarshal(data, &dashboards)
	if err != nil {
		return nil, err
	}

	if len(dashboards) == 0 {
		return nil, nil
	}

	return dashboards, nil
}
