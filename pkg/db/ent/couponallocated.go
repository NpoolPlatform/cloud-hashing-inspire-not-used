// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponallocated"
)

// CouponAllocated is the model entity for the CouponAllocated schema.
type CouponAllocated struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CouponAllocated) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case couponallocated.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CouponAllocated", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CouponAllocated fields.
func (ca *CouponAllocated) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case couponallocated.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ca.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this CouponAllocated.
// Note that you need to call CouponAllocated.Unwrap() before calling this method if this CouponAllocated
// was returned from a transaction, and the transaction was committed or rolled back.
func (ca *CouponAllocated) Update() *CouponAllocatedUpdateOne {
	return (&CouponAllocatedClient{config: ca.config}).UpdateOne(ca)
}

// Unwrap unwraps the CouponAllocated entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ca *CouponAllocated) Unwrap() *CouponAllocated {
	tx, ok := ca.config.driver.(*txDriver)
	if !ok {
		panic("ent: CouponAllocated is not a transactional entity")
	}
	ca.config.driver = tx.drv
	return ca
}

// String implements the fmt.Stringer.
func (ca *CouponAllocated) String() string {
	var builder strings.Builder
	builder.WriteString("CouponAllocated(")
	builder.WriteString(fmt.Sprintf("id=%v", ca.ID))
	builder.WriteByte(')')
	return builder.String()
}

// CouponAllocateds is a parsable slice of CouponAllocated.
type CouponAllocateds []*CouponAllocated

func (ca CouponAllocateds) config(cfg config) {
	for _i := range ca {
		ca[_i].config = cfg
	}
}
