package migrate

import (
	"context"
	"log"
	"os"

	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
	"github.com/adryanchiko/x-order/service/order-app/pkg/orderitem"
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

type (
	MigrateOrderItem struct {
		Data []OrderItem
	}

	OrderItem struct {
		ID           int    `csv:"id"`
		OrderID      int    `csv:"order_id"`
		PricePerUnit int    `csv:"price_per_unit"`
		Quantity     int    `csv:"quantity"`
		Product      string `csv:"product"`
	}
)

func (m *MigrateOrderItem) ReadFromCSV(file string) error {
	in, err := os.Open(file)
	if err != nil {
		return err
	}
	defer in.Close()

	orderItems := []OrderItem{}

	if err := gocsv.UnmarshalFile(in, &orderItems); err != nil {
		return err
	}

	m.Data = orderItems

	return nil
}

func (m *MigrateOrderItem) WriteToDB() error {
	records := make([]orderitem.NewOrderItem, len(m.Data))
	for i, data := range m.Data {
		records[i] = orderitem.NewOrderItem{
			ID:           data.ID,
			PricePerUnit: data.PricePerUnit,
			Quantity:     data.Quantity,
			Product:      data.Product,
			OrderID:      data.OrderID,
		}
	}

	store := orderitem.NewOrderItemEnt()
	if err := store.BulkCreate(context.Background(), records); err != nil {
		return err
	}

	return nil
}

func newMigrateOrderItem() Migrate {
	return &MigrateOrderItem{}
}

func MigrateOrderItemCommand(s *settings.Settings) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-item",
		Short: "import data order item from csv to db",
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

			migration := newMigrateOrderItem()
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

	cmd.Flags().String("file", "orderitem.csv", "CSV file path")

	return cmd
}
