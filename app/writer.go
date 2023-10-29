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

func WriteMatrix(gen Generator, rowCount, colCount uint, output io.Writer) error {
	bufWriter := bufio.NewWriter(output)

	for row := uint(0); row < rowCount; row++ {
		if err := appendToWriter(gen.AppendLineStart, bufWriter); err != nil {
			return fmt.Errorf("failed to write line start to the output")
		}
		for col := uint(1); col < colCount; col++ {
			if err := appendToWriter(gen.AppendElement, bufWriter); err != nil {
				return fmt.Errorf("failed to write element to the output")
			}
			if err := appendToWriter(gen.AppendDelimiter, bufWriter); err != nil {
				return fmt.Errorf("failed to write delimiter to the output")
			}
		}
		if err := appendToWriter(gen.AppendElement, bufWriter); err != nil {
			return fmt.Errorf("failed to write element to the output")
		}
		if err := appendToWriter(gen.AppendLineEnd, bufWriter); err != nil {
			return fmt.Errorf("failed to write line end to the output")
		}
	}

	err := bufWriter.Flush()
	return err
}

func appendToWriter(appender func([]byte) []byte, writer *bufio.Writer) error {
	buffer := writer.AvailableBuffer()
	appendedBuffer := appender(buffer)
	_, err := writer.Write(appendedBuffer)
	return err
}
