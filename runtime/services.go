package runtime

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"

	"github.com/T-ragon/cosmos-sdk/v3/runtime/services"
	"github.com/T-ragon/cosmos-sdk/v3/types/module"
)

func (a *App) registerRuntimeServices(cfg module.Configurator) error { // nolint:staticcheck // SA1019: Configurator is deprecated but still used in runtime v1.
	autocliv1.RegisterQueryServer(cfg.QueryServer(), services.NewAutoCLIQueryService(a.ModuleManager.Modules))

	reflectionSvc, err := services.NewReflectionService()
	if err != nil {
		return err
	}
	reflectionv1.RegisterReflectionServiceServer(cfg.QueryServer(), reflectionSvc)

	return nil
}
