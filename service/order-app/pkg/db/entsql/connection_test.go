package entsql

import (
	"testing"

	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
)

func TestDefaultPostgresDB(t *testing.T) {
	option := settings.SqlOption{
		Driver: "postgres",
		URI:    "postgresql://postgres:@localhost?sslmode=disable",
	}

	defer Close()

	if err := Open(&option); err != nil {
		t.Error(err)
	}
}

func TestMultiplePostgresDBs(t *testing.T) {
	option := settings.SqlOption{
		Driver: "postgres",
		URI:    "postgresql://postgres:@localhost?sslmode=disable",
	}

	defer Close()

	keys := []string{defaultKey, "key1", "key2"}

	for _, key := range keys {
		if err := OpenWithKey(key, &option); err != nil {
			t.Error(err)
		}
	}
}
