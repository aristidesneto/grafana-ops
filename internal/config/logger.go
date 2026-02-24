package config

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

// InitLogger inicializa o logger com o nível configurado via ENV
func InitLogger(loglevel string) *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()

		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
		})

		logger.SetOutput(os.Stdout)

		level, err := logrus.ParseLevel(loglevel)
		if err != nil {
			logger.Warnf("Nível de log inválido '%s', usando 'info'", loglevel)
			level = logrus.InfoLevel
		}

		logger.SetLevel(level)
	})

	return logger
}
