package client

import (
	"encoding/json"
	"testing"
)

func TestFilterOperatorConstants(t *testing.T) {
	tests := []struct {
		name     string
		operator FilterOperator
		expected string
	}{
		{"Equality operator", OpEq, "$eq"},
		{"Not equal operator", OpNe, "$ne"},
		{"Less than operator", OpLt, "$lt"},
		{"Less than or equal operator", OpLte, "$lte"},
		{"Greater than operator", OpGt, "$gt"},
		{"Greater than or equal operator", OpGte, "$gte"},
		{"In operator", OpIn, "$in"},
		{"Not in operator", OpNin, "$nin"},
		{"And operator", OpAnd, "$and"},
		{"Or operator", OpOr, "$or"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.operator) != tt.expected {
				t.Errorf("Expected operator %s to have value %s, got %s", tt.name, tt.expected, string(tt.operator))
			}
		})
	}
}

func TestNewFilter(t *testing.T) {
	filter := NewFilter()

	if filter == nil {
		t.Fatal("NewFilter() returned nil")
	}

	if filter.filter == nil {
		t.Fatal("NewFilter() did not initialize the filter map")
	}

	if len(filter.filter) != 0 {
		t.Errorf("Expected empty filter map, got %d items", len(filter.filter))
	}
}

func TestComparisonOperators(t *testing.T) {
	tests := []struct {
		name           string
		setupFilter    func(*FilterBuilder) *FilterBuilder
		expectedOutput string
	}{
		{
			name: "Eq operator",
			setupFilter: func(f *FilterBuilder) *FilterBuilder {
				return f.Eq("status", "open")
			},
			expectedOutput: `{"status":{"$eq":"open"}}`,
		},
		{
			name: "Ne operator",
			setupFilter: func(f *FilterBuilder) *FilterBuilder {
				return f.Ne("status", "closed")
			},
			expectedOutput: `{"status":{"$ne":"closed"}}`,
		},
		{
			name: "Lt operator",
			setupFilter: func(f *FilterBuilder) *FilterBuilder {
				return f.Lt("priority", 5)
			},
			expectedOutput: `{"priority":{"$lt":5}}`,
		},
		{
			name: "Lte operator",
			setupFilter: func(f *FilterBuilder) *FilterBuilder {
				return f.Lte("priority", 3)
			},
			expectedOutput: `{"priority":{"$lte":3}}`,
		},
		{
			name: "Gt operator",
			setupFilter: func(f *FilterBuilder) *FilterBuilder {
				return f.Gt("created_at", "2024-01-01")
			},
			expectedOutput: `{"created_at":{"$gt":"2024-01-01"}}`,
		},
		{
			name: "Gte operator",
			setupFilter: func(f *FilterBuilder) *FilterBuilder {
				return f.Gte("updated_at", "2024-01-01T00:00:00Z")
			},
			expectedOutput: `{"updated_at":{"$gte":"2024-01-01T00:00:00Z"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := NewFilter()
			result := tt.setupFilter(filter)

			// Test method chaining
			if result != filter {
				t.Error("Method should return the same FilterBuilder instance for chaining")
			}

			output := filter.Build()
			if output != tt.expectedOutput {
				t.Errorf("Expected %s, got %s", tt.expectedOutput, output)
			}
		})
	}
}

func TestInOperator(t *testing.T) {
	tests := []struct {
		name           string
		field          string
		values         []any
		expectedOutput string
	}{
		{
			name:           "In with string values",
			field:          "status",
			values:         []any{"open", "pending", "in-progress"},
			expectedOutput: `{"status":{"$in":["open","pending","in-progress"]}}`,
		},
		{
			name:           "In with integer values",
			field:          "priority",
			values:         []any{1, 2, 3},
			expectedOutput: `{"priority":{"$in":[1,2,3]}}`,
		},
		{
			name:           "In with mixed types",
			field:          "tags",
			values:         []any{"urgent", 42, true},
			expectedOutput: `{"tags":{"$in":["urgent",42,true]}}`,
		},
		{
			name:           "In with single value",
			field:          "category",
			values:         []any{"bug"},
			expectedOutput: `{"category":{"$in":["bug"]}}`,
		},
		{
			name:           "In with empty array",
			field:          "labels",
			values:         []any{},
			expectedOutput: `{"labels":{"$in":[]}}`,
		},
		{
			name:           "In with null values",
			field:          "assignee",
			values:         []any{nil, "user1", "user2"},
			expectedOutput: `{"assignee":{"$in":[null,"user1","user2"]}}`,
		},
		{
			name:           "In with boolean values",
			field:          "is_resolved",
			values:         []any{true, false},
			expectedOutput: `{"is_resolved":{"$in":[true,false]}}`,
		},
		{
			name:           "In with float values",
			field:          "score",
			values:         []any{1.5, 2.75, 3.0},
			expectedOutput: `{"score":{"$in":[1.5,2.75,3]}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := NewFilter()
			result := filter.In(tt.field, tt.values)

			// Test method chaining
			if result != filter {
				t.Error("In method should return the same FilterBuilder instance for chaining")
			}

			output := filter.Build()
			if output != tt.expectedOutput {
				t.Errorf("Expected %s, got %s", tt.expectedOutput, output)
			}

			// Verify the output is valid JSON
			var parsed map[string]any
			if err := json.Unmarshal([]byte(output), &parsed); err != nil {
				t.Errorf("Build output is not valid JSON: %v", err)
			}
		})
	}
}

func TestNinOperator(t *testing.T) {
	tests := []struct {
		name           string
		field          string
		values         []any
		expectedOutput string
	}{
		{
			name:           "Nin with string values",
			field:          "status",
			values:         []any{"closed", "cancelled"},
			expectedOutput: `{"status":{"$nin":["closed","cancelled"]}}`,
		},
		{
			name:           "Nin with integer values",
			field:          "priority",
			values:         []any{4, 5},
			expectedOutput: `{"priority":{"$nin":[4,5]}}`,
		},
		{
			name:           "Nin with empty array",
			field:          "tags",
			values:         []any{},
			expectedOutput: `{"tags":{"$nin":[]}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := NewFilter()
			result := filter.Nin(tt.field, tt.values)

			// Test method chaining
			if result != filter {
				t.Error("Nin method should return the same FilterBuilder instance for chaining")
			}

			output := filter.Build()
			if output != tt.expectedOutput {
				t.Errorf("Expected %s, got %s", tt.expectedOutput, output)
			}
		})
	}
}

func TestLogicalOperators(t *testing.T) {
	t.Run("And operator", func(t *testing.T) {
		filter1 := NewFilter().Eq("status", "open")
		filter2 := NewFilter().Gt("priority", 2)

		andFilter := NewFilter().And(filter1, filter2)
		output := andFilter.Build()

		expected := `{"$and":[{"status":{"$eq":"open"}},{"priority":{"$gt":2}}]}`
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
	})

	t.Run("Or operator", func(t *testing.T) {
		filter1 := NewFilter().Eq("status", "urgent")
		filter2 := NewFilter().Eq("priority", 5)

		orFilter := NewFilter().Or(filter1, filter2)
		output := orFilter.Build()

		expected := `{"$or":[{"status":{"$eq":"urgent"}},{"priority":{"$eq":5}}]}`
		if output != expected {
			t.Errorf("Expected %s, got %s", expected, output)
		}
	})

	t.Run("Complex logical combinations", func(t *testing.T) {
		// Create filters for: (status = "open" OR status = "pending") AND priority > 2
		filter1 := NewFilter().Eq("status", "open")
		filter2 := NewFilter().Eq("status", "pending")
		filter3 := NewFilter().Gt("priority", 2)

		orFilter := NewFilter().Or(filter1, filter2)
		finalFilter := NewFilter().And(orFilter, filter3)

		output := finalFilter.Build()

		// The output should contain both $and and $or operators
		var parsed map[string]any
		if err := json.Unmarshal([]byte(output), &parsed); err != nil {
			t.Errorf("Build output is not valid JSON: %v", err)
		}

		if _, exists := parsed["$and"]; !exists {
			t.Error("Expected $and operator in output")
		}
	})
}

func TestMethodChaining(t *testing.T) {
	filter := NewFilter()

	// Test that all methods return the same instance for chaining
	result := filter.Eq("status", "open").Ne("priority", 0).Lt("created_at", "2024-12-31")

	if result != filter {
		t.Error("Method chaining should return the same FilterBuilder instance")
	}

	output := filter.Build()

	// Verify that all conditions are present
	var parsed map[string]any
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Errorf("Build output is not valid JSON: %v", err)
	}

	expectedFields := []string{"status", "priority", "created_at"}
	for _, field := range expectedFields {
		if _, exists := parsed[field]; !exists {
			t.Errorf("Expected field %s in chained filter output", field)
		}
	}
}

func TestMultipleConditionsSameField(t *testing.T) {
	// Test adding multiple conditions to the same field
	filter := NewFilter()
	filter.Gte("priority", 1).Lte("priority", 5)

	output := filter.Build()

	var parsed map[string]any
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Errorf("Build output is not valid JSON: %v", err)
	}

	priority, exists := parsed["priority"]
	if !exists {
		t.Fatal("Expected priority field in output")
	}

	priorityMap, ok := priority.(map[string]any)
	if !ok {
		t.Fatal("Expected priority to be a map")
	}

	if _, exists := priorityMap["$gte"]; !exists {
		t.Error("Expected $gte operator for priority field")
	}

	if _, exists := priorityMap["$lte"]; !exists {
		t.Error("Expected $lte operator for priority field")
	}
}

func TestBuildWithEmptyFilter(t *testing.T) {
	filter := NewFilter()
	output := filter.Build()

	expected := "{}"
	if output != expected {
		t.Errorf("Expected %s for empty filter, got %s", expected, output)
	}
}

func TestBuildWithComplexNesting(t *testing.T) {
	// Create a complex nested filter structure
	filter1 := NewFilter().In("status", []any{"open", "pending"})
	filter2 := NewFilter().Gte("priority", 3)
	filter3 := NewFilter().Eq("assignee", "john.doe")
	filter4 := NewFilter().Lt("created_at", "2024-01-01")

	// ((status IN ["open", "pending"] AND priority >= 3) OR (assignee = "john.doe" AND created_at < "2024-01-01"))
	andFilter1 := NewFilter().And(filter1, filter2)
	andFilter2 := NewFilter().And(filter3, filter4)
	finalFilter := NewFilter().Or(andFilter1, andFilter2)

	output := finalFilter.Build()

	// Verify it's valid JSON and has the expected structure
	var parsed map[string]any
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Errorf("Build output is not valid JSON: %v", err)
	}

	if _, exists := parsed["$or"]; !exists {
		t.Error("Expected $or operator at root level")
	}

	t.Logf("Complex filter output: %s", output)
}

func TestInOperatorWithLargeArray(t *testing.T) {
	// Test In operator with a large array of values
	values := make([]any, 100)
	for i := 0; i < 100; i++ {
		values[i] = i + 1
	}

	filter := NewFilter().In("id", values)
	output := filter.Build()

	// Verify it's valid JSON
	var parsed map[string]any
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Errorf("Build output is not valid JSON: %v", err)
	}

	// Verify the array length is preserved
	id, exists := parsed["id"]
	if !exists {
		t.Fatal("Expected id field in output")
	}

	idMap, ok := id.(map[string]any)
	if !ok {
		t.Fatal("Expected id to be a map")
	}

	inArray, exists := idMap["$in"]
	if !exists {
		t.Fatal("Expected $in operator")
	}

	arrayValues, ok := inArray.([]any)
	if !ok {
		t.Fatal("Expected $in value to be an array")
	}

	if len(arrayValues) != 100 {
		t.Errorf("Expected array length 100, got %d", len(arrayValues))
	}
}

func TestBuildOutputConsistency(t *testing.T) {
	// Test that Build() produces consistent output for the same filter
	filter := NewFilter().Eq("status", "open").In("priority", []any{1, 2, 3})

	output1 := filter.Build()
	output2 := filter.Build()

	if output1 != output2 {
		t.Error("Build() should produce consistent output for the same filter")
	}
}

func TestFilterBuilderImmutability(t *testing.T) {
	// Test that creating a new filter doesn't affect existing ones
	filter1 := NewFilter().Eq("status", "open")
	output1 := filter1.Build()

	filter2 := NewFilter().Eq("status", "closed")
	output2 := filter2.Build()

	// Check that filter1 wasn't affected by creating filter2
	output1Again := filter1.Build()
	if output1 != output1Again {
		t.Error("Creating a new filter should not affect existing filters")
	}

	if output1 == output2 {
		t.Error("Different filters should produce different outputs")
	}
}
