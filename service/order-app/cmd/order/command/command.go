package commands

import (
	"github.com/adryanchiko/x-order/service/order-app/cmd/order/command/initdb"
	"github.com/adryanchiko/x-order/service/order-app/pkg/registry"
)

func init() {
	registry.RegisterCommandFactory(initdb.InitDbCommand)
}
