package main

import (
	"fmt"
)

func main() {

	input := `int x = 10;
float y = 3.14;
/* comentário de múltiplas linhas */
if (x >= y) {
    print(x);
}`

	scanner := NewScanner(input)
	tokens := scanner.ScanAll()

	printTokensGrouped(tokens)

}
func printTokensGrouped(tokens []Token) {
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
	
	// Operators
	ASSIGN     // =
	PLUS       // +
	MINUS      // -
	MULTIPLY   // *
	DIVIDE     // /
	MODULO     // %
	
	// Comparison
	EQUAL      // ==
	NOT_EQUAL  // !=
	LESS       // <
	LESS_EQUAL // <=
	GREATER    // >
	GTE        // >=
	
	// Logical
	AND        // &&
	OR         // ||
	NOT        // !
	
	// Delimiters
	SEMICOLON  // ;
	COMMA      // ,
	LPAREN     // (
	RPAREN     // )
	LBRACE     // {
	RBRACE     // }
	LBRACKET   // [
	RBRACKET   // ]
	
	// Special
	EOF
	ERROR
)

func (t TokenType) String() string {
	names := []string{
		"IDENTIFIER", "NUMBER",
		"INT", "FLOAT", "IF", "ELSE", "WHILE", "FOR", "RETURN", "PRINT",
		"ASSIGN", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO",
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

// Scanner represents the lexical analyzer
type Scanner struct {
	input    string
	position int  // current position (current char)
	line     int  // current line number
	column   int  // current column number
	keywords map[string]TokenType
}

// NewScanner creates a new lexical scanner
func NewScanner(input string) *Scanner {
	scanner := &Scanner{
		input:  input,
		line:   1,
		column: 1,
		keywords: map[string]TokenType{
			"int":    INT,
			"float":  FLOAT,
			"if":     IF,
			"else":   ELSE,
			"while":  WHILE,
			"for":    FOR,
			"return": RETURN,
			"print":  PRINT,
		},
	}
	return scanner
}

// currentChar returns the current character or 0 if at end
func (s *Scanner) currentChar() byte {
	if s.position >= len(s.input) {
		return 0
	}
	return s.input[s.position]
}

// peekChar returns the next character without advancing position
func (s *Scanner) peekChar() byte {
	if s.position+1 >= len(s.input) {
		return 0
	}
	return s.input[s.position+1]
}

// advance moves to the next character
func (s *Scanner) advance() {
	if s.position < len(s.input) && s.input[s.position] == '\n' {
		s.line++
		s.column = 1
	} else {
		s.column++
	}
	s.position++
}

// skipWhitespace skips whitespace characters
func (s *Scanner) skipWhitespace() {
	for s.currentChar() != 0 && isWhitespace(s.currentChar()) {
		s.advance()
	}
}

// skipLineComment skips // style comments
func (s *Scanner) skipLineComment() {
	for s.currentChar() != 0 && s.currentChar() != '\n' {
		s.advance()
	}
}

// skipBlockComment skips /* */ style comments
func (s *Scanner) skipBlockComment() error {
	s.advance() // skip '/'
	s.advance() // skip '*'
	
	for s.currentChar() != 0 {
		if s.currentChar() == '*' && s.peekChar() == '/' {
			s.advance() // skip '*'
			s.advance() // skip '/'
			return nil
		}
		s.advance()
	}
	
	return fmt.Errorf("unterminated block comment")
}

// readIdentifier reads an identifier or keyword
func (s *Scanner) readIdentifier() string {
	start := s.position
	
	for s.currentChar() != 0 && (isAlpha(s.currentChar()) || isDigit(s.currentChar()) || s.currentChar() == '_') {
		s.advance()
	}
	
	return s.input[start:s.position]
}

// readNumber reads a numeric literal (integer or float)
func (s *Scanner) readNumber() string {
	start := s.position
	
	// Read integer part
	for s.currentChar() != 0 && isDigit(s.currentChar()) {
		s.advance()
	}
	
	// Check for decimal point
	if s.currentChar() == '.' && isDigit(s.peekChar()) {
		s.advance() // skip '.'
		
		// Read fractional part
		for s.currentChar() != 0 && isDigit(s.currentChar()) {
			s.advance()
		}
	}
	
	return s.input[start:s.position]
}

// NextToken returns the next token from the input
func (s *Scanner) NextToken() Token {
	for s.currentChar() != 0 {
		// Skip whitespace
		s.skipWhitespace()
		
		if s.currentChar() == 0 {
			break
		}
		
		// Save current position for token
		line := s.line
		column := s.column
		position := s.position
		
		char := s.currentChar()
		
		// Handle comments
		if char == '/' {
			if s.peekChar() == '/' {
				s.skipLineComment()
				continue
			} else if s.peekChar() == '*' {
				if err := s.skipBlockComment(); err != nil {
					return Token{ERROR, err.Error(), line, column, position}
				}
				continue
			}
		}
		
		// Identifiers and keywords
		if isAlpha(char) || char == '_' {
			value := s.readIdentifier()
			tokenType := IDENTIFIER
			
			// Check if it's a keyword
			if keyword, exists := s.keywords[value]; exists {
				tokenType = keyword
				value = "" // Don't store value for keywords in this format
			}
			
			return Token{tokenType, value, line, column, position}
		}
		
		// Numbers
		if isDigit(char) {
			value := s.readNumber()
			return Token{NUMBER, value, line, column, position}
		}
		
		// Two-character operators
		if char == '=' && s.peekChar() == '=' {
			s.advance()
			s.advance()
			return Token{EQUAL, "", line, column, position}
		}
		
		if char == '!' && s.peekChar() == '=' {
			s.advance()
			s.advance()
			return Token{NOT_EQUAL, "", line, column, position}
		}
		
		if char == '<' && s.peekChar() == '=' {
			s.advance()
			s.advance()
			return Token{LESS_EQUAL, "", line, column, position}
		}
		
		if char == '>' && s.peekChar() == '=' {
			s.advance()
			s.advance()
			return Token{GTE, "", line, column, position}
		}
		
		if char == '&' && s.peekChar() == '&' {
			s.advance()
			s.advance()
			return Token{AND, "", line, column, position}
		}
		
		if char == '|' && s.peekChar() == '|' {
			s.advance()
			s.advance()
			return Token{OR, "", line, column, position}
		}
		
		// Single-character tokens
		s.advance()
		
		switch char {
		case '=':
			return Token{ASSIGN, "", line, column, position}
		case '+':
			return Token{PLUS, "", line, column, position}
		case '-':
			return Token{MINUS, "", line, column, position}
		case '*':
			return Token{MULTIPLY, "", line, column, position}
		case '/':
			return Token{DIVIDE, "", line, column, position}
		case '%':
			return Token{MODULO, "", line, column, position}
		case '<':
			return Token{LESS, "", line, column, position}
		case '>':
			return Token{GREATER, "", line, column, position}
		case '!':
			return Token{NOT, "", line, column, position}
		case ';':
			return Token{SEMICOLON, "", line, column, position}
		case ',':
			return Token{COMMA, "", line, column, position}
		case '(':
			return Token{LPAREN, "", line, column, position}
		case ')':
			return Token{RPAREN, "", line, column, position}
		case '{':
			return Token{LBRACE, "", line, column, position}
		case '}':
			return Token{RBRACE, "", line, column, position}
		case '[':
			return Token{LBRACKET, "", line, column, position}
		case ']':
			return Token{RBRACKET, "", line, column, position}
		default:
			return Token{ERROR, fmt.Sprintf("unexpected character: '%c'", char), line, column, position}
		}
	}
	
	return Token{EOF, "", s.line, s.column, s.position}
}

// ScanAll scans all tokens from the input
func (s *Scanner) ScanAll() []Token {
	var tokens []Token
	
	for {
		token := s.NextToken()
		if token.Type == EOF {
			break
		}
		if token.Type == ERROR {
			fmt.Printf("Lexical error: %s\n", token.Value)
			continue
		}
		tokens = append(tokens, token)
	}
	
	return tokens
}

// Helper functions
func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}



