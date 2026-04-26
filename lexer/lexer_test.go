package lexer

import (
	"testing"

	"github.com/MihoZaki/MonkeyGo/token"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;

		let add = fn(x,y){
			x + y ;
		};

		let result = add(five, ten);
 		!-/5*;
 		5 < 10 > 5;
 		if (5 < 10){
 			return true;
 		}else {
 			return false;
 		}

 		10 == 10;
 		10 != 9; 
 		i++;
 		i+= 1;
 		"foobar"
 		"foo bar"
 		[1, 2];
 		{"foo": "bar"}
 		3.143
 		1.
 		.5
 		`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.INT, "5"},
		{token.ASTERISK, "*"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.INCREMENT, "++"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.PLUS_ASSIGN, "+="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.FLOAT, "3.143"},
		{token.FLOAT, "1."},
		{token.FLOAT, ".5"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d]- tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d]- literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestSingleLineComment(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.TokenType
	}{
		{
			name:     "comment_before_code",
			input:    "// this is a comment\nlet x = 5;",
			expected: []token.TokenType{token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.EOF},
		},
		{
			name:     "inline_comment",
			input:    "let x = 5;// inline comment\nlet y = 6;",
			expected: []token.TokenType{token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.EOF},
		},
		{
			name:     "empty_comments",
			input:    "//\n// empty comment\nlet z = 0;",
			expected: []token.TokenType{token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.EOF},
		},
		{
			name:     "comment_at_eof_no_newline",
			input:    "let z = 0; // comment without newline at EOF",
			expected: []token.TokenType{token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.EOF},
		},
		{
			name:     "comment_with_special_chars",
			input:    "// comment with symbols: !@#$%^&*()\nlet a = 1;",
			expected: []token.TokenType{token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.EOF},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)
			for i, exp := range tt.expected {
				tok := l.NextToken()
				if tok.Type != exp {
					t.Errorf("step %d: expected token %s, got %s (literal: %q)",
						i, exp, tok.Type, tok.Literal)
				}
			}
			if final := l.NextToken(); final.Type != token.EOF {
				t.Errorf("expected EOF after test, got %s", final.Type)
			}
		})
	}
}

func TestMultiLineComments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.TokenType
	}{
		{
			"basic inline",
			"/* comment */ let x = 5;",
			[]token.TokenType{token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.EOF},
		},
		{
			"multi-line with newlines",
			"/* line1\nline2\nline3 */ let y = 10;",
			[]token.TokenType{token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.EOF},
		},
		{
			"empty comment",
			"/**/ let z = 0;",
			[]token.TokenType{token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.EOF},
		},
		{
			"comment before and after code",
			"/* start */ let a = 1; /* end */",
			[]token.TokenType{token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.EOF},
		},
		{
			"unclosed comment (should consume to EOF)",
			"let x = 5; /* never closed",
			[]token.TokenType{token.LET, token.IDENT, token.ASSIGN, token.INT, token.SEMICOLON, token.EOF},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)
			for _, exp := range tt.expected {
				tok := l.NextToken()
				if tok.Type != exp {
					t.Errorf("expected %s, got %s", exp, tok.Type)
				}
			}
		})
	}
}

func TestUnicodeIdentifiers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []struct {
			tokenType token.TokenType
			literal   string
		}
	}{
		{
			name:  "chinese_identifier",
			input: "let 你好 = 5;",
			expected: []struct {
				tokenType token.TokenType
				literal   string
			}{
				{token.LET, "let"},
				{token.IDENT, "你好"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name:  "mixed_ascii_unicode",
			input: "let user_名前 = 10; 名前;",
			expected: []struct {
				tokenType token.TokenType
				literal   string
			}{
				{token.LET, "let"},
				{token.IDENT, "user_名前"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.IDENT, "名前"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name:  "unicode_whitespace_handling",
			input: "let　x　=　5;", // full-width spaces (U+3000)
			expected: []struct {
				tokenType token.TokenType
				literal   string
			}{
				{token.LET, "let"},
				{token.IDENT, "x"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name:  "emoji_should_be_illegal_in_ident",
			input: "let 🚀 = 1;",
			expected: []struct {
				tokenType token.TokenType
				literal   string
			}{
				{token.LET, "let"},
				{token.ILLEGAL, "🚀"}, // or skip? we'll decide semantics
				{token.ASSIGN, "="},
				{token.INT, "1"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)
			for i, exp := range tt.expected {
				tok := l.NextToken()
				if tok.Type != exp.tokenType {
					t.Errorf("step %d: expected type %s, got %s", i, exp.tokenType, tok.Type)
				}
				if tok.Literal != exp.literal {
					t.Errorf("step %d: expected literal %q, got %q", i, exp.literal, tok.Literal)
				}
			}
		})
	}
}
