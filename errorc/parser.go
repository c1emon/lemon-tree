package errorc

type Parser interface {
	Do(any) (string, ErrorType)
	Support(any) bool
}

var ParserHolder = &parserHolder{parsers: make(map[string]Parser)}

type parserHolder struct {
	parsers map[string]Parser
}

func (h *parserHolder) Add(name string, parser Parser) {
	h.parsers[name] = parser
}

func (h *parserHolder) Remove(name string) {
	delete(h.parsers, name)
}

func (h *parserHolder) Iter() map[string]Parser {
	return h.parsers
}
