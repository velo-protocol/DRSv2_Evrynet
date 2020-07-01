package vfactory

import "github.com/velo-protocol/DRSv2_Evrynet/go/cmd/gvel/utils/config"

type veloFactory struct {
	AppConfig config.Configuration
}

func NewVeloFactory(appConfig config.Configuration) Repository {
	return &veloFactory{
		AppConfig: appConfig,
	}
}
