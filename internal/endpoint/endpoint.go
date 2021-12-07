package endpoint

import "github.com/google/wire"

type Endpoints struct {
}

func NewEndpoints() *Endpoints {
	return &Endpoints{}
}

var ProviderSet = wire.NewSet(NewEndpoints)
