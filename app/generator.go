package app

type Generator interface {
	AppendLineStart([]byte)
	AppendLineEnd([]byte)
	AppendElement([]byte)
	AppendDelimiter([]byte)
}
