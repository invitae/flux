package remote

import (
	"context"

	"go.uber.org/zap"

	"github.com/fluxcd/flux/pkg/api"
	v10 "github.com/fluxcd/flux/pkg/api/v10"
	v11 "github.com/fluxcd/flux/pkg/api/v11"
	v6 "github.com/fluxcd/flux/pkg/api/v6"
	v9 "github.com/fluxcd/flux/pkg/api/v9"
	"github.com/fluxcd/flux/pkg/job"
	"github.com/fluxcd/flux/pkg/update"
)

var _ api.Server = &ErrorLoggingServer{}

type ErrorLoggingServer struct {
	server api.Server
	logger *zap.SugaredLogger
}

func NewErrorLoggingServer(s api.Server, l *zap.SugaredLogger) *ErrorLoggingServer {
	return &ErrorLoggingServer{s, l}
}

func (p *ErrorLoggingServer) Export(ctx context.Context) (config []byte, err error) {
	defer func() {
		if err != nil {
			// Omit config as it could be large
			p.logger.Error(zap.string("method", "Export"), zap.Error(err))
		}
	}()
	return p.server.Export(ctx)
}

func (p *ErrorLoggingServer) ListServices(ctx context.Context, maybeNamespace string) (_ []v6.ControllerStatus, err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "ListServices"), zap.Error(err))
		}
	}()
	return p.server.ListServices(ctx, maybeNamespace)
}

func (p *ErrorLoggingServer) ListServicesWithOptions(ctx context.Context, opts v11.ListServicesOptions) (_ []v6.ControllerStatus, err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "ListServicesWithOptions"), zap.Error(err))
		}
	}()
	return p.server.ListServicesWithOptions(ctx, opts)
}

func (p *ErrorLoggingServer) ListImages(ctx context.Context, spec update.ResourceSpec) (_ []v6.ImageStatus, err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "ListImages"), zap.Error(err))
		}
	}()
	return p.server.ListImages(ctx, spec)
}

func (p *ErrorLoggingServer) ListImagesWithOptions(ctx context.Context, opts v10.ListImagesOptions) (_ []v6.ImageStatus, err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "ListImagesWithOptions"), zap.Error(err))
		}
	}()
	return p.server.ListImagesWithOptions(ctx, opts)
}

func (p *ErrorLoggingServer) JobStatus(ctx context.Context, jobID job.ID) (_ job.Status, err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "JobStatus"), zap.Error(err))
		}
	}()
	return p.server.JobStatus(ctx, jobID)
}

func (p *ErrorLoggingServer) SyncStatus(ctx context.Context, ref string) (_ []string, err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "SyncStatus"), zap.Error(err))
		}
	}()
	return p.server.SyncStatus(ctx, ref)
}

func (p *ErrorLoggingServer) UpdateManifests(ctx context.Context, u update.Spec) (_ job.ID, err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "UpdateManifests"), zap.Error(err))
		}
	}()
	return p.server.UpdateManifests(ctx, u)
}

func (p *ErrorLoggingServer) GitRepoConfig(ctx context.Context, regenerate bool) (_ v6.GitConfig, err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "GitRepoConfig"), zap.Error(err))
		}
	}()
	return p.server.GitRepoConfig(ctx, regenerate)
}

func (p *ErrorLoggingServer) Ping(ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "Ping"), zap.Error(err))
		}
	}()
	return p.server.Ping(ctx)
}

func (p *ErrorLoggingServer) Version(ctx context.Context) (v string, err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "Version"), zap.Error(err), zap.String("version", v))
		}
	}()
	return p.server.Version(ctx)
}

func (p *ErrorLoggingServer) NotifyChange(ctx context.Context, change v9.Change) (err error) {
	defer func() {
		if err != nil {
			p.logger.Error(zap.string("method", "NotifyChange"), zap.Error(err))
		}
	}()
	return p.server.NotifyChange(ctx, change)
}
