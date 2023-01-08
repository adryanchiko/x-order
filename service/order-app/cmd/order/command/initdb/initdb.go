package initdb

import (
	"context"
	"log"

	"github.com/adryanchiko/x-order/service/order-app/ent/migrate"
	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
	"github.com/spf13/cobra"
)

func InitDbCommand(s *settings.Settings) *cobra.Command {
	return &cobra.Command{
		Use:   "init-db",
		Short: "Init or migrate database schema",
		Run: func(cmd *cobra.Command, args []string) {
			// init ent
			err := entsql.Open(&s.Conn.Sql)
			if err != nil {
				log.Fatal(err)
			}
			defer entsql.Close()

			err = entsql.DB().Debug().Schema.Create(
				context.Background(),
				migrate.WithDropIndex(true),
				migrate.WithDropColumn(true),
			)
			if err != nil {
				log.Fatalf("failed creating schema resources: %v", err)
			}
		},
	}
}
