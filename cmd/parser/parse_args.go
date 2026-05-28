package parser

const AmountOfArgs = 3

type Arg = string

type Parser struct {
	args   []Arg
	parsed map[string]Arg
}

type ParsedArgs struct {
	Count        int
	WindowWidth  int32
	WindowHeight int32
}

func NewParser(args []Arg) *Parser {
	return &Parser{args: args, parsed: make(map[string]Arg)}
}

func (p *Parser) Parse() *ParsedArgs {
	return &ParsedArgs{}
}
