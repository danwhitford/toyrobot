package toyrobot

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTokenise(t *testing.T) {
	table := []struct {
		input    string
		expected []Token
	}{
		{
			"3 2 NORTH PLACE",
			[]Token{
				{Type: TOKEN_NUMBER, Value: 3, Lexeme: "3"},
				{Type: TOKEN_NUMBER, Value: 2, Lexeme: "2"},
				{Type: TOKEN_DIRECTION, Value: NORTH, Lexeme: "NORTH"},
				{Type: TOKEN_WORD, Value: "PLACE", Lexeme: "PLACE"},
			},
		},
		{
			"RIGHT",
			[]Token{
				{Type: TOKEN_WORD, Value: "RIGHT", Lexeme: "RIGHT"},
			},
		},
		{
			"REPORT",
			[]Token{
				{Type: TOKEN_WORD, Value: "REPORT", Lexeme: "REPORT"},
			},
		},
		{
			"MOVE LEFT RIGHT REPORT",
			[]Token{
				{Type: TOKEN_WORD, Value: "MOVE", Lexeme: "MOVE"},
				{Type: TOKEN_WORD, Value: "LEFT", Lexeme: "LEFT"},
				{Type: TOKEN_WORD, Value: "RIGHT", Lexeme: "RIGHT"},
				{Type: TOKEN_WORD, Value: "REPORT", Lexeme: "REPORT"},
			},
		},
		{
			"NORTH SOUTH EAST WEST",
			[]Token{
				{Type: TOKEN_DIRECTION, Value: NORTH, Lexeme: "NORTH"},
				{Type: TOKEN_DIRECTION, Value: SOUTH, Lexeme: "SOUTH"},
				{Type: TOKEN_DIRECTION, Value: EAST, Lexeme: "EAST"},
				{Type: TOKEN_DIRECTION, Value: WEST, Lexeme: "WEST"},
			},
		},
		{
			"10 20 30 40",
			[]Token{
				{Type: TOKEN_NUMBER, Value: 10, Lexeme: "10"},
				{Type: TOKEN_NUMBER, Value: 20, Lexeme: "20"},
				{Type: TOKEN_NUMBER, Value: 30, Lexeme: "30"},
				{Type: TOKEN_NUMBER, Value: 40, Lexeme: "40"},
			},
		},
		{
			"+ - * /",
			[]Token{
				{Type: TOKEN_WORD, Value: "+", Lexeme: "+"},
				{Type: TOKEN_WORD, Value: "-", Lexeme: "-"},
				{Type: TOKEN_WORD, Value: "*", Lexeme: "*"},
				{Type: TOKEN_WORD, Value: "/", Lexeme: "/"},
			},
		},
		{
			"\"hello world\"",
			[]Token{
				{Type: TOKEN_STRING, Value: "hello world", Lexeme: "\"hello world\""},
			},
		},
		{
			"15 5 GT IF \"15 is bigger than 5\" . FI",
			[]Token{
				{Type: TOKEN_NUMBER, Value: 15, Lexeme: "15"},
				{Type: TOKEN_NUMBER, Value: 5, Lexeme: "5"},
				{Type: TOKEN_WORD, Value: "GT", Lexeme: "GT"},
				{Type: TOKEN_WORD, Value: "IF", Lexeme: "IF"},
				{Type: TOKEN_STRING, Value: "15 is bigger than 5", Lexeme: "\"15 is bigger than 5\""},
				{Type: TOKEN_WORD, Value: ".", Lexeme: "."},
				{Type: TOKEN_WORD, Value: "FI", Lexeme: "FI"},
			},
		},
	}

	for _, tst := range table {
		tokeniser := RobotTokeniser{}
		got, err := tokeniser.Tokenise(tst.input)
		if err != nil {
			t.Errorf("Error tokenising '%s': '%s'", tst.input, err)
		}

		if diff := cmp.Diff(tst.expected, got); diff != "" {
			t.Errorf("Tokenise(%s) mismatch (-want +got):\n%s", tst.input, diff)
		}
	}
}

func TestTokenise_Errors(t *testing.T) {
	table := []struct {
		input         string
		expectedError string
	}{
		{"10 10 EQ IF \"equal\" . ELSE \"not equal\" . \" FI", "unterminated string"},
	}

	for _, tst := range table {
		tokeniser := RobotTokeniser{}
		_, err := tokeniser.Tokenise(tst.input)
		if err == nil {
			t.Fatalf("Expected error tokenising '%s'", tst.input)
		}

		if err.Error() != tst.expectedError {
			t.Errorf("Expected error '%s' tokenising '%s'", tst.expectedError, tst.input)
		}
	}
}
