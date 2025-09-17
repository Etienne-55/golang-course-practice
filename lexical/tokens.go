package lexical

import(
	"fmt"
)


func PrintTokensGrouped(tokens []Token) {
	if len(tokens) == 0 {
		return
	}

	currentLine := tokens[0].Line
	lineTokens := []Token{}

	for _, token := range tokens {
		if token.Line != currentLine {
			// Print current line tokens
			if len(lineTokens) > 0 {
				printTokenLine(lineTokens)
			}
			// Start new line
			currentLine = token.Line
			lineTokens = []Token{token}
		} else {
			lineTokens = append(lineTokens, token)
		}
	}
	// Print last line
	if len(lineTokens) > 0 {
		printTokenLine(lineTokens)
	}
}

func printTokenLine(tokens []Token) {
	fmt.Print("[")
	for i, token := range tokens {
		if i > 0 {
			fmt.Print(", ")
		}
		if token.Value != "" && token.Type != IDENTIFIER {
			fmt.Printf("%s(%s)", token.Type, token.Value)
		} else if token.Type == IDENTIFIER {
			fmt.Printf("IDENT(%s)", token.Value)
		} else {
			fmt.Print(token.Type)
		}
	}
	fmt.Println("]")
}

// TokenType represents the different types of tokens
type TokenType int

const (
	// Literals
	IDENTIFIER TokenType = iota
	NUMBER
	
	// Keywords
	INT
	FLOAT
	IF
	ELSE
	WHILE
	FOR
	RETURN
	PRINT
	
	ASSIGN     // =
	PLUS       // +
	MINUS      // -
	MULTIPLY   // *
	DIVIDE     // /
	MODULO     // %
	INCREMENT	 // ++
	DECREMENT	 // --
	
	EQUAL      // ==
	NOT_EQUAL  // !=
	LESS       // <
	LESS_EQUAL // <=
	GREATER    // >
	GTE        // >=
	
	AND        // &&
	OR         // ||
	NOT        // !
	
	SEMICOLON  // ;
	COMMA      // ,
	LPAREN     // (
	RPAREN     // )
	LBRACE     // {
	RBRACE     // }
	LBRACKET   // [
	RBRACKET   // ]
	
	EOF
	ERROR
)

func (t TokenType) String() string {
	names := []string{
		"IDENTIFIER", "NUMBER",
		"INT", "FLOAT", "IF", "ELSE", "WHILE", "FOR", "RETURN", "PRINT",
		"ASSIGN", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO", "INCREMENT",
		"EQUAL", "NOT_EQUAL", "LESS", "LESS_EQUAL", "GREATER", "GTE",
		"AND", "OR", "NOT",
		"SEMICOLON", "COMMA", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET",
		"EOF", "ERROR",
	}
	if int(t) < len(names) {
		return names[t]
	}
	return "UNKNOWN"
}

// Token represents a lexical token
type Token struct {
	Type     TokenType
	Value    string
	Line     int
	Column   int
	Position int
}

func (t Token) String() string {
	if t.Value != "" {
		return fmt.Sprintf("%s(%s) at %d:%d", t.Type, t.Value, t.Line, t.Column)
	}
	return fmt.Sprintf("%s at %d:%d", t.Type, t.Line, t.Column)
}

