package app

import (
	"fmt"
	"math/rand"
	"strconv"
)

type NumberGenerator struct {
	RangeFrom int32
	RangeTo   int32
	Delimiter string
}

func (g *NumberGenerator) AppendLineStart(buffer []byte) []byte {
	return fmt.Append(buffer, "")
}

func (g *NumberGenerator) AppendLineEnd(buffer []byte) []byte {
	return fmt.Appendln(buffer, "")
}

func (g *NumberGenerator) AppendElement(buffer []byte) []byte {
	randomNumber := rand.Intn(int(g.RangeTo)-int(g.RangeFrom)) + int(g.RangeFrom)
	return strconv.AppendInt(buffer, int64(randomNumber), 10)
}

func (g *NumberGenerator) AppendDelimiter(buffer []byte) []byte {
	return fmt.Append(buffer, g.Delimiter)
}
