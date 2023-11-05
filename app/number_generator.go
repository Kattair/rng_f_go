package app

import (
	"fmt"
	"math/rand"
	"strconv"
)

type NumberGenerator struct {
	rangeFrom  int32
	rangeTo    int32
	totalRange int64
	lineStart  []byte
	lineEnd    []byte
	delimiter  []byte
}

func NewNumberGenerator(rangeFrom, rangeTo int32, delimiter string) (*NumberGenerator, error) {
	totalRange := int64(rangeTo) - int64(rangeFrom)
	if totalRange <= 0 {
		return nil, fmt.Errorf("failed to create number generator with range <%d, %d) - empty range", rangeFrom, rangeTo)
	}
	lineStart := []byte("")
	lineEnd := []byte("\n")
	delimiterBytes := []byte(delimiter)
	return &NumberGenerator{rangeFrom, rangeTo, totalRange, lineStart, lineEnd, delimiterBytes}, nil
}

func (g *NumberGenerator) AppendLineStart(buffer []byte) []byte {
	return append(buffer, g.lineStart...)
}

func (g *NumberGenerator) AppendLineEnd(buffer []byte) []byte {
	return append(buffer, g.lineEnd...)
}

func (g *NumberGenerator) AppendElement(buffer []byte) []byte {
	randomNumber := rand.Int63n(g.totalRange) + int64(g.rangeFrom)
	return strconv.AppendInt(buffer, randomNumber, 10)
}

func (g *NumberGenerator) AppendDelimiter(buffer []byte) []byte {
	return append(buffer, []byte(g.delimiter)...)
}
