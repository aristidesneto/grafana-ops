package types

type Config struct {
	GrafanaURL   string `mapstructure:"grafana-url"`
	GrafanaToken string `mapstructure:"grafana-token"`
	Output       string `mapstructure:"output"`
	LogLevel     string `mapstructure:"loglevel"`
}
