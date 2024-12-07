package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestValueMarshal(t *testing.T) {
	testCases := []struct {
		name     string
		value    Value
		expected []byte
	}{
		{
			name: "Marshal String",
			value: Value{
				typ: "string",
				str: "Hello",
			},
			expected: []byte("+Hello\r\n"),
		},
		{
			name: "Marshal Array with Bulk values",
			value: Value{
				typ: "array",
				array: []Value{
					{typ: "bulk", bulk: "Hello"},
					{typ: "bulk", bulk: "World"},
				},
			},
			expected: []byte("*2\r\n$5\r\nHello\r\n$5\r\nWorld\r\n"),
		},
		{
			name: "Marshal Array with String values",
			value: Value{
				typ: "array",
				array: []Value{
					{typ: "string", str: "Hello"},
					{typ: "string", str: "World"},
				},
			},
			expected: []byte("*2\r\n+Hello\r\n+World\r\n"),
		},
		{
			name: "Marshal Bulk",
			value: Value{
				typ:  "bulk",
				bulk: "Hello",
			},
			expected: []byte("$5\r\nHello\r\n"),
		},
		{
			name: "Marshal Null",
			value: Value{
				typ: "null",
			},
			expected: []byte("$-1\r\n"),
		},
		{
			name: "Marshal Error",
			value: Value{
				typ: "error",
				str: "Error",
			},
			expected: []byte("-Error\r\n"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.value.Marshal()
			if string(result) != string(tc.expected) {
				t.Errorf("Expected %s, got %s", string(tc.expected), string(result))
			}
		})
	}
}

func TestWrite(t *testing.T) {
	testCases := []struct {
		name     string
		value    Value
		expected []byte
	}{
		{
			name: "Write String",
			value: Value{
				typ: "string",
				str: "Hello",
			},
			expected: []byte("+Hello\r\n"),
		},
		{
			name: "Write Array with Bulk values",
			value: Value{
				typ: "array",
				array: []Value{
					{typ: "bulk", bulk: "Hello"},
					{typ: "bulk", bulk: "World"},
				},
			},
			expected: []byte("*2\r\n$5\r\nHello\r\n$5\r\nWorld\r\n"),
		},
		{
			name: "Write Array with String values",
			value: Value{
				typ: "array",
				array: []Value{
					{typ: "string", str: "Hello"},
					{typ: "string", str: "World"},
				},
			},
			expected: []byte("*2\r\n+Hello\r\n+World\r\n"),
		},
		{
			name: "Write Bulk string",
			value: Value{
				typ:  "bulk",
				bulk: "Hello",
			},
			expected: []byte("$5\r\nHello\r\n"),
		},
		{
			name: "Write Null",
			value: Value{
				typ: "null",
			},
			expected: []byte("$-1\r\n"),
		},
		{
			name: "Write Error",
			value: Value{
				typ: "error",
				str: "Error",
			},
			expected: []byte("-Error\r\n"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//create buffer to capture written bytes
			buf := &bytes.Buffer{}
			writer := NewWriter(buf)
			err := writer.Write(tc.value)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			//check for the written bytes
			// fmt.Println(string(buf.Bytes()))
			if string(buf.Bytes()) != string(tc.expected) {
				t.Errorf("Expected %s, got %s", string(tc.expected), string(buf.Bytes()))
			}
		})
	}
}

func TestWriterError(t *testing.T) {
	// create mock writer that always returns an error
	mockErrorWriter := &errorWriter{}
	writer := NewWriter(mockErrorWriter)

	// Try to write a value

	err := writer.Write(Value{typ: "string", str: "test"})
	if err == nil {
		t.Error("Expected an error got a nil")
	}
}

type errorWriter struct{}

func (*errorWriter) Write([]byte) (int, error) {
	return 0, errors.New("writer error")
}
