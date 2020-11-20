package api

import (
	"context"

	"github.com/rancher/harvester-server/pkg/api/datavolume"
	"github.com/rancher/harvester-server/pkg/api/image"
	"github.com/rancher/harvester-server/pkg/api/keypair"
	"github.com/rancher/harvester-server/pkg/api/network"
	"github.com/rancher/harvester-server/pkg/api/user"
	"github.com/rancher/harvester-server/pkg/api/vm"
	"github.com/rancher/harvester-server/pkg/api/vmtemplate"
	"github.com/rancher/harvester-server/pkg/config"

	"github.com/rancher/steve/pkg/server"
)

type registerSchema func(scaled *config.Scaled, server *server.Server) error

func registerSchemas(scaled *config.Scaled, server *server.Server, registers ...registerSchema) error {
	for _, register := range registers {
		if err := register(scaled, server); err != nil {
			return err
		}
	}
	return nil
}

func Setup(ctx context.Context, server *server.Server) error {
	scaled := config.ScaledWithContext(ctx)
	return registerSchemas(scaled, server,
		image.RegisterSchema,
		keypair.RegisterSchema,
		vmtemplate.RegisterSchema,
		vm.RegisterSchema,
		user.RegisterSchema,
		network.RegisterSchema,
		datavolume.RegisterSchema)
}
