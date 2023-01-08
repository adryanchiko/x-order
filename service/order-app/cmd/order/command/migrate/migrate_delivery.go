package migrate

import (
	"context"
	"log"
	"os"

	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
	"github.com/adryanchiko/x-order/service/order-app/pkg/delivery"
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

type (
	MigrateDelivery struct {
		Data []Delivery
	}

	Delivery struct {
		ID                int `csv:"id"`
		OrderItemID       int `csv:"order_item_id"`
		DeliveredQuantity int `csv:"delivered_quantity"`
	}
)

func (m *MigrateDelivery) ReadFromCSV(file string) error {
	in, err := os.Open(file)
	if err != nil {
		return err
	}
	defer in.Close()

	deliveries := []Delivery{}

	if err := gocsv.UnmarshalFile(in, &deliveries); err != nil {
		return err
	}

	m.Data = deliveries

	return nil
}

func (m *MigrateDelivery) WriteToDB() error {
	records := make([]delivery.NewDelivery, len(m.Data))
	for i, data := range m.Data {
		records[i] = delivery.NewDelivery{
			ID:                data.ID,
			DeliveredQuantity: data.DeliveredQuantity,
			OrderItemID:       data.OrderItemID,
		}
	}

	store := delivery.NewDeliveryEnt()
	if err := store.BulkCreate(context.Background(), records); err != nil {
		return err
	}

	return nil
}

func newMigrateDelivery() Migrate {
	return &MigrateDelivery{}
}

func MigrateDeliveryCommand(s *settings.Settings) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delivery",
		Short: "import data delivery from csv to db",
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

			migration := newMigrateDelivery()
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

	cmd.Flags().String("file", "delivery.csv", "CSV file path")

	return cmd
}
