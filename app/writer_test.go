package app_test

import (
	"log"
	"math"
	"os"
	"testing"

	"github.com/Kattair/rng_f_go/app"
)

func BenchmarkWriteMatrixToFile(b *testing.B) {
	outputFilename := "output.txt"
	gen, _ := app.NewNumberGenerator(math.MinInt32, math.MaxInt32, " ")
	b.Cleanup(func() {
		if err := os.Remove(outputFilename); err != nil {
			log.Fatalf("failed to remove output file '%s'", outputFilename)
		}
	})

	app.WriteMatrixToFile(gen, 10000, 10000, outputFilename)
}
