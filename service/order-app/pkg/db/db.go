package db

import (
	"log"

	"github.com/adryanchiko/x-order/service/order-app/pkg/db/entsql"
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
)

func Open(s *settings.Settings) error {
	if s.Conn.Sql.Enabled {
		log.Println("Initialize sql connection")
		if err := entsql.Open(&s.Conn.Sql); err != nil {
			return err
		}
	}

	return nil
}

func Close(s *settings.Settings) error {
	if s.Conn.Sql.Enabled {
		if err := entsql.Close(); err != nil {
			return err
		}
	}

	return nil
}
