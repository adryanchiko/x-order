package migrate

import (
	"fmt"

	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
	"github.com/spf13/cobra"
)

type (
	Migrate interface {
		ReadFromCSV(file string) error
		WriteToDB() error
	}
)

func InitMigrateCommand(s *settings.Settings) *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate data from csv to database",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("please specify the migration data type on sub-command")
		},
	}

	initCmd.AddCommand(MigrateCompanyCommand(s))
	initCmd.AddCommand(MigrateCustomerCommand(s))
	initCmd.AddCommand(MigrateOrderCommand(s))
	initCmd.AddCommand(MigrateOrderItemCommand(s))
	initCmd.AddCommand(MigrateDeliveryCommand(s))

	return initCmd
}
