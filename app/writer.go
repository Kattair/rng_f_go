package app

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteMatrixToFile(gen Generator, rowCount, colCount uint, output string) error {
	file, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("failed to open file '%s' for writing", output)
	}
	defer file.Close()

	WriteMatrix(gen, rowCount, colCount, file)

	return nil
}

func WriteMatrix(gen Generator, rowCount, colCount uint, output io.Writer) {
	bufWriter := bufio.NewWriter(output)

	for row := uint(0); row < rowCount; row++ {
		appendToWriter(gen.AppendLineStart, bufWriter)
		for col := uint(1); col < colCount; col++ {
			appendToWriter(gen.AppendElement, bufWriter)
			appendToWriter(gen.AppendDelimiter, bufWriter)
		}
		appendToWriter(gen.AppendElement, bufWriter)
		appendToWriter(gen.AppendLineEnd, bufWriter)
	}
}

func appendToWriter(appender func([]byte) []byte, writer *bufio.Writer) {
	buffer := writer.AvailableBuffer()
	appendedBuffer := appender(buffer)
	writer.Write(appendedBuffer)
}
