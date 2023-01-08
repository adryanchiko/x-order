package migrate

import (
	"context"
	"log"
	"os"

	"github.com/adryanchiko/x-order/service/order-app/pkg/company"
	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

type (
	MigrateCompany struct {
		Data []Company
	}

	Company struct {
		CompanyID   int    `csv:"company_id"`
		CompanyName string `csv:"company_name"`
	}
)

func (m *MigrateCompany) ReadFromCSV(file string) error {
	in, err := os.Open(file)
	if err != nil {
		return err
	}
	defer in.Close()

	companies := []Company{}

	if err := gocsv.UnmarshalFile(in, &companies); err != nil {
		return err
	}

	m.Data = companies

	return nil
}

func (m *MigrateCompany) WriteToDB() error {
	records := make([]company.NewCompany, len(m.Data))
	for i, data := range m.Data {
		records[i] = company.NewCompany{
			ID:          data.CompanyID,
			CompanyName: data.CompanyName,
		}
	}

	store := company.NewCompanyEnt()
	if err := store.BulkCreate(context.Background(), records); err != nil {
		return err
	}

	return nil
}

func newMigrateCompany() Migrate {
	return &MigrateCompany{}
}

func MigrateCompanyCommand(s *settings.Settings) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "company",
		Short: "import data company from csv to db",
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

			migration := newMigrateCompany()
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

	cmd.Flags().String("file", "companies.csv", "CSV file path")

	return cmd
}
