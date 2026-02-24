package cmd

import (
	"grafana-ops/internal"
	"grafana-ops/internal/client"
	"grafana-ops/internal/config"
	"grafana-ops/internal/utils"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Salva todos os componentes do Grafana",
	Long: `Comando para salvar dashboards do Grafana.

Utilize os subcomandos para listar, exportar ou gerenciar seus dashboards.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		log := config.InitLogger(Config.LogLevel)
		log.Infof("running save command; grafana=%s output=%s", Config.GrafanaURL, Config.Output)
		client := client.NewClient(Config.GrafanaURL, Config.GrafanaToken, log)
		writer := utils.NewFileWriter(Config.Output, log)

		save := internal.NewSaveService(client, writer, log)
		if err := save.Save(); err != nil {
			log.Errorf("backup failed: %v", err)
			return err
		}
		log.Info("backup completed")

		return nil
	},
}

// listCmd representa o subcomando para listar dashboards
// var listCmd = &cobra.Command{
// 	Use:   "list",
// 	Short: "Lista todos os dashboards do Grafana",
// 	Long: `Lista todos os dashboards disponíveis no Grafana.

// Exemplo:
//   grafana-backup dashboards list`,
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		return dashboards.ListDashboards()
// 	},
// }

// saveCmd representa o subcomando para salvar dashboards
// var saveCmd = &cobra.Command{
// 	Use:   "save",
// 	Short: "Salva todos os dashboards do Grafana",
// 	Long: `Salva todos os dashboards disponíveis no Grafana.

// Exemplo:
//   grafana-backup dashboards save`,
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		return internal.Save()
// 	},
// }

func init() {
	rootCmd.AddCommand(saveCmd)
}
