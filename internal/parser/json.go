package parser

const (
	OpenObject   = "OpenObject"
	CloseObject  = "CloseObject"
	OpenArray    = "OpenArray"
	CloseArray   = "CloseArray"
	Comma        = "Comma"
	PropertyName = "PropertyName"
	StringValue  = "String"
	IntValue     = "Int"
)

type TokenType string

type JsonToken struct {
	Type    TokenType
	Literal string
}

type JsonParser struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *JsonParser {
	p := &JsonParser{input: input}
	p.readChar()
	return p
}

func (p *JsonParser) readChar() {
	if p.readPosition >= len(p.input) {
		p.ch = 0
	} else {
		p.ch = p.input[p.readPosition]
	}
	p.position = p.readPosition
	p.readPosition += 1
}

func (p *JsonParser) ReadObject() []JsonToken {
	tokens := []JsonToken{{OpenObject, "{"}}
	p.readChar()
	for p.ch != '}' {
		p.readChar()
		name := p.readText()
		p.readChar()
		tokens = append(tokens, JsonToken{PropertyName, name})
		tokens = append(tokens, p.readValue()...)
	}
	tokens = append(tokens, JsonToken{CloseObject, "}"})
	p.readChar()
	return tokens
}

func (p *JsonParser) readValue() []JsonToken {
	tokens := []JsonToken{}

loop:
	for {
		switch p.ch {
		case '{':
			tokens = append(tokens, p.ReadObject()...)
		case '[':
			p.readChar()
			tokens = append(tokens, p.readArray()...)
		case '"':
			p.readChar()
			tokens = append(tokens, p.readString())
			p.readChar()
		case ',':
			p.readChar()
			break loop
		case '}', ']':
			break loop
		default:
			if p.isDigit() {
				tokens = append(tokens, p.readInt())
			} else {
				p.readChar()
			}
		}
	}
	return tokens
}

func (p *JsonParser) readArray() []JsonToken {
	tokens := []JsonToken{{OpenArray, "["}}
	for p.ch != ']' {
		tokens = append(tokens, p.readValue()...)
		if p.ch == ',' {
			p.readChar()
		}
	}
	p.readChar()
	tokens = append(tokens, JsonToken{CloseArray, "]"})
	return tokens
}

func (p *JsonParser) readString() JsonToken {
	return JsonToken{StringValue, p.readText()}
}

func (p *JsonParser) readInt() JsonToken {
	return JsonToken{IntValue, p.readNumber()}
}

func (p *JsonParser) readText() string {
	position := p.position
	for p.ch != '"' && p.ch != 0 {
		p.readChar()
	}
	return p.input[position:p.position]
}

func (p *JsonParser) isDigit() bool {
	return p.ch == '-' || '0' <= p.ch && p.ch <= '9'
}

func (p *JsonParser) readNumber() string {
	position := p.position
	for p.isDigit() {
		p.readChar()
	}
	return p.input[position:p.position]
}
