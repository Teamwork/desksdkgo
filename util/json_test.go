package util

import (
	"testing"
)

type TestStruct struct {
	Name    string
	Age     int
	Address string
	Tags    []string
}

func TestMergeJSONData(t *testing.T) {
	tests := []struct {
		name     string
		target   TestStruct
		data     map[string]interface{}
		expected TestStruct
	}{
		{
			name:   "Merge with empty target",
			target: TestStruct{},
			data: map[string]interface{}{
				"Name":    "John Doe",
				"Age":     30,
				"Address": "123 Main St",
				"Tags":    []string{"developer", "golang"},
			},
			expected: TestStruct{
				Name:    "John Doe",
				Age:     30,
				Address: "123 Main St",
				Tags:    []string{"developer", "golang"},
			},
		},
		{
			name: "Merge with existing data",
			target: TestStruct{
				Name:    "Jane Doe",
				Age:     25,
				Address: "456 Oak St",
				Tags:    []string{"designer"},
			},
			data: map[string]interface{}{
				"Name": "Jane Smith",
				"Age":  26,
			},
			expected: TestStruct{
				Name:    "Jane Smith",
				Age:     26,
				Address: "456 Oak St",
				Tags:    []string{"designer"},
			},
		},
		{
			name: "Merge with partial data",
			target: TestStruct{
				Name:    "Original Name",
				Age:     40,
				Address: "Original Address",
				Tags:    []string{"original"},
			},
			data: map[string]interface{}{
				"Name": "New Name",
			},
			expected: TestStruct{
				Name:    "New Name",
				Age:     40,
				Address: "Original Address",
				Tags:    []string{"original"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeJSONData(&tt.target, tt.data)

			if tt.target.Name != tt.expected.Name {
				t.Errorf("Name: got %v, want %v", tt.target.Name, tt.expected.Name)
			}
			if tt.target.Age != tt.expected.Age {
				t.Errorf("Age: got %v, want %v", tt.target.Age, tt.expected.Age)
			}
			if tt.target.Address != tt.expected.Address {
				t.Errorf("Address: got %v, want %v", tt.target.Address, tt.expected.Address)
			}
			if len(tt.target.Tags) != len(tt.expected.Tags) {
				t.Errorf("Tags length: got %v, want %v", len(tt.target.Tags), len(tt.expected.Tags))
			} else {
				for i, tag := range tt.target.Tags {
					if tag != tt.expected.Tags[i] {
						t.Errorf("Tags[%d]: got %v, want %v", i, tag, tt.expected.Tags[i])
					}
				}
			}
		})
	}
}
