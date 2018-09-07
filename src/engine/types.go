package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url string
	ParserFunc ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items []Item
}

//Type string
type Item struct {
	Url string
	Id string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
