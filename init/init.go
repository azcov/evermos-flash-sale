package init

import "go.uber.org/zap"

// StartAppInit init all server needs
func StartAppInit() (config *Config, logger *zap.Logger) {
	logger = setupLogger()

	config = setupMainConfig()

	return
}
