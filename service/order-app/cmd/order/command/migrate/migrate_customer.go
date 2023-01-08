package migrate

import (
	"context"
	"log"
	"os"

	"github.com/adryanchiko/x-order/service/order-app/pkg/customer"
	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

type (
	MigrateCustomer struct {
		Data []Customer
	}

	Customer struct {
		UserID      string   `csv:"user_id"`
		Login       string   `csv:"login"`
		Password    string   `csv:"password"`
		Name        string   `csv:"name"`
		CompanyID   int      `csv:"company_id"`
		CreditCards []string `csv:"credit_cards"`
	}
)

func (m *MigrateCustomer) ReadFromCSV(file string) error {
	in, err := os.Open(file)
	if err != nil {
		return err
	}
	defer in.Close()

	customers := []Customer{}

	if err := gocsv.UnmarshalFile(in, &customers); err != nil {
		return err
	}

	m.Data = customers

	return nil
}

func (m *MigrateCustomer) WriteToDB() error {
	records := make([]customer.NewCustomer, len(m.Data))
	for i, data := range m.Data {
		records[i] = customer.NewCustomer{
			ID:          data.UserID,
			Login:       data.Login,
			Password:    data.Password,
			Name:        data.Name,
			CreditCards: data.CreditCards,
			CompanyID:   data.CompanyID,
		}
	}

	store := customer.NewCustomerEnt()
	if err := store.BulkCreate(context.Background(), records); err != nil {
		return err
	}

	return nil
}

func newMigrateCustomer() Migrate {
	return &MigrateCustomer{}
}

func MigrateCustomerCommand(s *settings.Settings) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "customer",
		Short: "import data customer from csv to db",
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

			migration := newMigrateCustomer()
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

	cmd.Flags().String("file", "customer.csv", "CSV file path")

	return cmd
}
