//go:build go1.18
// +build go1.18

package deepcopy

import (
	"encoding/json"
	"net/http"
	"testing"
)

type ComplexStruct struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	IsEnabled   bool             `json:"is_enabled"`
	Count       int              `json:"count"`
	Price       float64          `json:"price"`
	Tags        []string         `json:"tags"`
	CreatedAt   string           `json:"created_at"`
	Options     map[string]bool  `json:"options"`
	Nested      NestedStruct     `json:"nested"`
	Items       []ItemStruct     `json:"items"`
	ExtraData   json.RawMessage  `json:"extra_data"`
	Handler     http.HandlerFunc `json:"handler"`
	Opts        map[string]any   `json:"opts"`
}

type NestedStruct struct {
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
	ItemAStruct
	*ItemBStruct
}

type ItemAStruct struct {
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
}

type ItemBStruct struct {
	ItemType string `json:"item_type"`
	ItemFrom string `json:"item_from"`
}

type ItemStruct struct {
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
}

func BenchmarkDeepCopy(b *testing.B) {
	complexData := ComplexStruct{
		ID:          1,
		Name:        "Complex Object",
		Description: "A struct with various types for JSON benchmarking",
		IsEnabled:   true,
		Count:       5,
		Price:       99.99,
		Tags:        []string{"tag1", "tag2", "tag3"},
		CreatedAt:   "2023-08-07T12:34:56Z",
		Options:     map[string]bool{"option1": true, "option2": false, "option3": true},
		Nested: NestedStruct{
			Field1: 42,
			Field2: "Nested Field",
			ItemAStruct: ItemAStruct{
				ItemID:   100,
				ItemName: "Item A",
			},
			ItemBStruct: &ItemBStruct{
				ItemType: "Item Type",
				ItemFrom: "Item From",
			},
		},
		Items: []ItemStruct{
			{ItemID: 101, ItemName: "Item 1"},
			{ItemID: 102, ItemName: "Item 2"},
		},
		ExtraData: json.RawMessage(`{"key": "value"}`),
		Handler:   func(w http.ResponseWriter, r *http.Request) {},
		Opts: map[string]any{
			"opt1": "value1",
			"opt2": 2,
			"opt3": true,
		},
	}

	for i := 0; i < b.N; i++ {
		_ = Copy[ComplexStruct](complexData)
		_ = Copy[*ComplexStruct](&complexData)
	}
}
