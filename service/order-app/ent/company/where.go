// Code generated by ent, DO NOT EDIT.

package company

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/adryanchiko/x-order/service/order-app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Company {
	return predicate.Company(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Company {
	return predicate.Company(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Company {
	return predicate.Company(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Company {
	return predicate.Company(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Company {
	return predicate.Company(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Company {
	return predicate.Company(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Company {
	return predicate.Company(sql.FieldLTE(FieldID, id))
}

// CompanyName applies equality check predicate on the "company_name" field. It's identical to CompanyNameEQ.
func CompanyName(v string) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldCompanyName, v))
}

// CompanyNameEQ applies the EQ predicate on the "company_name" field.
func CompanyNameEQ(v string) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldCompanyName, v))
}

// CompanyNameNEQ applies the NEQ predicate on the "company_name" field.
func CompanyNameNEQ(v string) predicate.Company {
	return predicate.Company(sql.FieldNEQ(FieldCompanyName, v))
}

// CompanyNameIn applies the In predicate on the "company_name" field.
func CompanyNameIn(vs ...string) predicate.Company {
	return predicate.Company(sql.FieldIn(FieldCompanyName, vs...))
}

// CompanyNameNotIn applies the NotIn predicate on the "company_name" field.
func CompanyNameNotIn(vs ...string) predicate.Company {
	return predicate.Company(sql.FieldNotIn(FieldCompanyName, vs...))
}

// CompanyNameGT applies the GT predicate on the "company_name" field.
func CompanyNameGT(v string) predicate.Company {
	return predicate.Company(sql.FieldGT(FieldCompanyName, v))
}

// CompanyNameGTE applies the GTE predicate on the "company_name" field.
func CompanyNameGTE(v string) predicate.Company {
	return predicate.Company(sql.FieldGTE(FieldCompanyName, v))
}

// CompanyNameLT applies the LT predicate on the "company_name" field.
func CompanyNameLT(v string) predicate.Company {
	return predicate.Company(sql.FieldLT(FieldCompanyName, v))
}

// CompanyNameLTE applies the LTE predicate on the "company_name" field.
func CompanyNameLTE(v string) predicate.Company {
	return predicate.Company(sql.FieldLTE(FieldCompanyName, v))
}

// CompanyNameContains applies the Contains predicate on the "company_name" field.
func CompanyNameContains(v string) predicate.Company {
	return predicate.Company(sql.FieldContains(FieldCompanyName, v))
}

// CompanyNameHasPrefix applies the HasPrefix predicate on the "company_name" field.
func CompanyNameHasPrefix(v string) predicate.Company {
	return predicate.Company(sql.FieldHasPrefix(FieldCompanyName, v))
}

// CompanyNameHasSuffix applies the HasSuffix predicate on the "company_name" field.
func CompanyNameHasSuffix(v string) predicate.Company {
	return predicate.Company(sql.FieldHasSuffix(FieldCompanyName, v))
}

// CompanyNameEqualFold applies the EqualFold predicate on the "company_name" field.
func CompanyNameEqualFold(v string) predicate.Company {
	return predicate.Company(sql.FieldEqualFold(FieldCompanyName, v))
}

// CompanyNameContainsFold applies the ContainsFold predicate on the "company_name" field.
func CompanyNameContainsFold(v string) predicate.Company {
	return predicate.Company(sql.FieldContainsFold(FieldCompanyName, v))
}

// HasCustomers applies the HasEdge predicate on the "customers" edge.
func HasCustomers() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CustomersTable, CustomersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomersWith applies the HasEdge predicate on the "customers" edge with a given conditions (other predicates).
func HasCustomersWith(preds ...predicate.Customer) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CustomersTable, CustomersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		p(s.Not())
	})
}
