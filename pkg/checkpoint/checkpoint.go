package checkpoint

import (
	"time"

	"github.com/weaveworks/go-checkpoint"
	"go.uber.org/zap"
)

const (
	versionCheckPeriod = 6 * time.Hour
)

func CheckForUpdates(product, version string, extra map[string]string, logger *zap.SugaredLogger) *checkpoint.Checker {
	handleResponse := func(r *checkpoint.CheckResponse, err error) {
		if err != nil {
			logger.Error(zap.Error(err))
			return
		}
		if r.Outdated {
			logger.Warn("update available", zap.String("latest", r.CurrentVersion), zap.String("URL", r.CurrentDownloadURL))
			return
		}
		logger.Info("up to date", zap.String("latest", r.CurrentVersion))
	}

	flags := map[string]string{
		"kernel-version": getKernelVersion(),
	}
	for k, v := range extra {
		flags[k] = v
	}

	params := checkpoint.CheckParams{
		Product:       product,
		Version:       version,
		SignatureFile: "",
		Flags:         flags,
	}

	return checkpoint.CheckInterval(&params, versionCheckPeriod, handleResponse)
}
