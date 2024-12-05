package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestReadLine(t *testing.T) {
	// test := []struct{}
	resp := NewResp(strings.NewReader("*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"))
	value, err := resp.Read()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(value)
}

// func TestReadArray(t *testing.T) {
// 	resp := NewResp(strings.NewReader("*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"))
// 	value, err := resp.Read()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Expected array
// 	expected := Value{
// 		typ: "array",
// 		array: []Value{
// 			{typ: "bulk", bulk: "hello"},
// 			{typ: "bulk", bulk: "world"},
// 		},
// 	}

// 	// Validate the result
// 	if len(value.array) != len(expected.array) {
// 		t.Fatalf("Expected array length %d, got %d", len(expected.array), len(value.array))
// 	}
// 	for i, v := range value.array {
// 		if v.typ != expected.array[i].typ || v.bulk != expected.array[i].bulk {
// 			t.Fatalf("Expected element %d: %+v, got %+v", i, expected.array[i], v)
// 		}
// 	}

// 	fmt.Printf("Parsed Value: %+v\n", value)
// }
