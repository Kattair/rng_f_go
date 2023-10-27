package app

type Generator interface {
	AppendLineStart(buffer []byte) []byte
	AppendLineEnd(buffer []byte) []byte
	AppendElement(buffer []byte) []byte
	AppendDelimiter(buffer []byte) []byte
}
