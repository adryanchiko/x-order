package entsql

import (
	"fmt"
	"log"
	"sync"

	"github.com/adryanchiko/x-order/service/order-app/ent"
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
	_ "github.com/lib/pq"
)

const defaultKey = "default"

var (
	mapper sync.Map
)

func Open(opts *settings.SqlOption) error {
	return OpenWithKey(defaultKey, opts)
}

func OpenWithKey(key string, opts *settings.SqlOption) error {
	if _, ok := mapper.Load(key); ok {
		return fmt.Errorf("DB with key '%s' already initialized", key)
	}

	db, err := open(opts)
	if err != nil {
		return err
	}

	mapper.Store(key, db)

	return nil
}

func open(opts *settings.SqlOption) (*ent.Client, error) {
	c, err := ent.Open(opts.Driver, opts.URI)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func DB(keys ...string) *ent.Client {
	key := defaultKey
	if len(keys) > 0 && len(keys[0]) > 0 {
		key = keys[0]
	}

	instance, ok := mapper.Load(key)
	if !ok {
		log.Fatalf("No declared '%s' found, need to open with Open() or OpenWithKey().", key)
	}

	return instance.(*ent.Client)
}

func Close() error {
	mapper.Range(func(k, v interface{}) bool {
		log.Printf("Closing Ent DB: %s", k)
		err := (v.(*ent.Client)).Close()
		if err != nil {
			log.Println(err)
		}

		mapper.Delete(k)

		return true
	})

	return nil
}
