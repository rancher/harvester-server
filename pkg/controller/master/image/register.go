package image

import (
	"context"

	"github.com/rancher/harvester-server/pkg/config"
)

func Register(ctx context.Context, management *config.Management) error {
	RegisterController(ctx, management)
	return nil
}
