package parser

import (
	"testing"
)

func TestJsonParser(t *testing.T) {
	input := `{"d":"red","e":[1,2,3,4],"f":5,"g":[[[-3,"red",{"i":"red"}],"red"]],"h":{"i":"red"}}`
	p := New(input)

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{OpenObject, "{"},
		{PropertyName, "d"},
		{StringValue, "red"},
		{PropertyName, "e"},
		{OpenArray, "["},
		{IntValue, "1"},
		{IntValue, "2"},
		{IntValue, "3"},
		{IntValue, "4"},
		{CloseArray, "]"},
		{PropertyName, "f"},
		{IntValue, "5"},
		{PropertyName, "g"},
		{OpenArray, "["},
		{OpenArray, "["},
		{OpenArray, "["},
		{IntValue, "-3"},
		{StringValue, "red"},
		{OpenObject, "{"},
		{PropertyName, "i"},
		{StringValue, "red"},
		{CloseObject, "}"},
		{CloseArray, "]"},
		{StringValue, "red"},
		{CloseArray, "]"},
		{CloseArray, "]"},
		{PropertyName, "h"},
		{OpenObject, "{"},
		{PropertyName, "i"},
		{StringValue, "red"},
		{CloseObject, "}"},
		{CloseObject, "}"},
	}

	r := p.ReadObject()

	for i, tt := range tests {
		if r[i].Type != tt.expectedType {
			t.Fatalf("tests[%d] - type wrong. expected=%q, got %q", i, tt.expectedType, r[i].Type)
		}

		if r[i].Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, r[i].Literal)
		}
	}
}
