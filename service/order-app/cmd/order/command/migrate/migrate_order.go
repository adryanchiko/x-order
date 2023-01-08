package migrate

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
	"github.com/adryanchiko/x-order/service/order-app/pkg/order"
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

type (
	MigrateOrder struct {
		Data []Order
	}

	Order struct {
		ID         int       `csv:"id"`
		CreatedAt  time.Time `csv:"created_at"`
		OrderName  string    `csv:"order_name"`
		CustomerID string    `csv:"customer_id"`
	}
)

func (m *MigrateOrder) ReadFromCSV(file string) error {
	in, err := os.Open(file)
	if err != nil {
		return err
	}
	defer in.Close()

	orders := []Order{}

	if err := gocsv.UnmarshalFile(in, &orders); err != nil {
		return err
	}

	m.Data = orders

	return nil
}

func (m *MigrateOrder) WriteToDB() error {
	records := make([]order.NewOrder, len(m.Data))
	for i, data := range m.Data {
		records[i] = order.NewOrder{
			ID:         data.ID,
			CreatedAt:  data.CreatedAt,
			OrderName:  data.OrderName,
			CustomerID: data.CustomerID,
		}
	}

	store := order.NewOrderEnt()
	if err := store.BulkCreate(context.Background(), records); err != nil {
		return err
	}

	return nil
}

func newMigrateOrder() Migrate {
	return &MigrateOrder{}
}

func MigrateOrderCommand(s *settings.Settings) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order",
		Short: "import data order from csv to db",
		RunE: func(cmd *cobra.Command, args []string) error {
			// init ent
			err := entsql.Open(&s.Conn.Sql)
			if err != nil {
				log.Fatal(err)
			}
			defer entsql.Close()

			filePath, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}

			migration := newMigrateOrder()
			if err := migration.ReadFromCSV(filePath); err != nil {
				log.Printf("failed read csv file: %v", err)
				return err
			}

			if err := migration.WriteToDB(); err != nil {
				log.Printf("failed write to db: %v", err)
				return err
			}

			return nil
		},
	}

	cmd.Flags().String("file", "order.csv", "CSV file path")

	return cmd
}
