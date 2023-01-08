package registry

import (
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"

	"github.com/spf13/cobra"
)

type CommandFactory func(*settings.Settings) *cobra.Command

var _commands []CommandFactory

func RegisterCommandFactory(factory CommandFactory) {
	_commands = append(_commands, factory)
}

func CommandFactories() []CommandFactory {
	return _commands
}
