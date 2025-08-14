package client

import (
	"encoding/json"
	"fmt"
)

// FilterOperator represents the available filter operators
type FilterOperator string

const (
	// Comparison operators
	OpEq  FilterOperator = "$eq"
	OpNe  FilterOperator = "$ne"
	OpLt  FilterOperator = "$lt"
	OpLte FilterOperator = "$lte"
	OpGt  FilterOperator = "$gt"
	OpGte FilterOperator = "$gte"
	OpIn  FilterOperator = "$in"
	OpNin FilterOperator = "$nin"

	// Logical operators
	OpAnd FilterOperator = "$and"
	OpOr  FilterOperator = "$or"
)

// FilterBuilder helps build MongoDB-style filters
type FilterBuilder struct {
	filter map[string]interface{}
}

// NewFilter creates a new FilterBuilder
func NewFilter() *FilterBuilder {
	return &FilterBuilder{
		filter: make(map[string]interface{}),
	}
}

// Eq adds an equality condition
func (f *FilterBuilder) Eq(field string, value interface{}) *FilterBuilder {
	f.addCondition(field, OpEq, value)
	return f
}

// Ne adds a not-equal condition
func (f *FilterBuilder) Ne(field string, value interface{}) *FilterBuilder {
	f.addCondition(field, OpNe, value)
	return f
}

// Lt adds a less-than condition
func (f *FilterBuilder) Lt(field string, value interface{}) *FilterBuilder {
	f.addCondition(field, OpLt, value)
	return f
}

// Lte adds a less-than-or-equal condition
func (f *FilterBuilder) Lte(field string, value interface{}) *FilterBuilder {
	f.addCondition(field, OpLte, value)
	return f
}

// Gt adds a greater-than condition
func (f *FilterBuilder) Gt(field string, value interface{}) *FilterBuilder {
	f.addCondition(field, OpGt, value)
	return f
}

// Gte adds a greater-than-or-equal condition
func (f *FilterBuilder) Gte(field string, value interface{}) *FilterBuilder {
	f.addCondition(field, OpGte, value)
	return f
}

// In adds an in-list condition
func (f *FilterBuilder) In(field string, values ...interface{}) *FilterBuilder {
	f.addCondition(field, OpIn, values)
	return f
}

// Nin adds a not-in-list condition
func (f *FilterBuilder) Nin(field string, values ...interface{}) *FilterBuilder {
	f.addCondition(field, OpNin, values)
	return f
}

// And adds an AND condition
func (f *FilterBuilder) And(filters ...*FilterBuilder) *FilterBuilder {
	conditions := make([]map[string]interface{}, len(filters))
	for i, filter := range filters {
		conditions[i] = filter.filter
	}
	f.filter[string(OpAnd)] = conditions
	return f
}

// Or adds an OR condition
func (f *FilterBuilder) Or(filters ...*FilterBuilder) *FilterBuilder {
	conditions := make([]map[string]interface{}, len(filters))
	for i, filter := range filters {
		conditions[i] = filter.filter
	}
	f.filter[string(OpOr)] = conditions
	return f
}

// Build returns the filter as a JSON string
func (f *FilterBuilder) Build() string {
	data, err := json.Marshal(f.filter)
	if err != nil {
		// This should never happen as we control the structure
		panic(fmt.Sprintf("failed to marshal filter: %v", err))
	}
	return string(data)
}

// addCondition adds a condition to the filter
func (f *FilterBuilder) addCondition(field string, op FilterOperator, value interface{}) {
	if _, exists := f.filter[field]; !exists {
		f.filter[field] = make(map[string]interface{})
	}
	f.filter[field].(map[string]interface{})[string(op)] = value
}
