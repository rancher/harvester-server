package auth

import (
	"context"

	"github.com/rancher/harvester-server/pkg/config"
	"github.com/rancher/harvester-server/pkg/settings"

	"github.com/rancher/steve/pkg/server"
)

func Register(ctx context.Context, scaled *config.Scaled, server *server.Server) error {
	go WatchSecret(ctx, scaled, config.Namespace, settings.AuthSecretName.Get())
	return nil
}
