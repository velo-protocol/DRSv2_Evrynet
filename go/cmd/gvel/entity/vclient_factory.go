package entity

import "github.com/velo-protocol/DRSv2_Evrynet/go/libs/vclient"

type NewClientInput struct {
	RpcUrl          string
	PrivateKey      string
	ContractAddress *vclient.ContractAddress
}
