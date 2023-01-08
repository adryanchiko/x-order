// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "company_name", Type: field.TypeString},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:       "companies",
		Columns:    CompaniesColumns,
		PrimaryKey: []*schema.Column{CompaniesColumns[0]},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "login", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "credit_cards", Type: field.TypeJSON},
		{Name: "company_customers", Type: field.TypeInt, Nullable: true},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "customers_companies_customers",
				Columns:    []*schema.Column{CustomersColumns[5]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DeliveriesColumns holds the columns for the "deliveries" table.
	DeliveriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "delivered_quantity", Type: field.TypeInt},
		{Name: "order_item_deliveries", Type: field.TypeInt, Nullable: true},
	}
	// DeliveriesTable holds the schema information for the "deliveries" table.
	DeliveriesTable = &schema.Table{
		Name:       "deliveries",
		Columns:    DeliveriesColumns,
		PrimaryKey: []*schema.Column{DeliveriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deliveries_order_items_deliveries",
				Columns:    []*schema.Column{DeliveriesColumns[2]},
				RefColumns: []*schema.Column{OrderItemsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "order_name", Type: field.TypeString},
		{Name: "customer_orders", Type: field.TypeString, Nullable: true},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_customers_orders",
				Columns:    []*schema.Column{OrdersColumns[3]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderItemsColumns holds the columns for the "order_items" table.
	OrderItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price_per_unit", Type: field.TypeInt},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "product", Type: field.TypeString},
		{Name: "order_order_items", Type: field.TypeInt, Nullable: true},
	}
	// OrderItemsTable holds the schema information for the "order_items" table.
	OrderItemsTable = &schema.Table{
		Name:       "order_items",
		Columns:    OrderItemsColumns,
		PrimaryKey: []*schema.Column{OrderItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_items_orders_order_items",
				Columns:    []*schema.Column{OrderItemsColumns[4]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompaniesTable,
		CustomersTable,
		DeliveriesTable,
		OrdersTable,
		OrderItemsTable,
	}
)

func init() {
	CustomersTable.ForeignKeys[0].RefTable = CompaniesTable
	DeliveriesTable.ForeignKeys[0].RefTable = OrderItemsTable
	OrdersTable.ForeignKeys[0].RefTable = CustomersTable
	OrderItemsTable.ForeignKeys[0].RefTable = OrdersTable
}