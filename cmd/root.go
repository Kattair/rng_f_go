/*
Copyright © 2023 Kattair <martin.kustra@tutanota.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/Kattair/rng_f_go/app"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rng_f_go ROW_COUNT COL_COUNT",
	Short: "Simple program to generate a matrix and write it into a file written in Go.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var delimiter string
var output string
var rangeFrom int32
var rangeTo int32

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rng_f_go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&delimiter, "delimiter", "d", " ", "Define a string to be used as the delimiter between columns")
	rootCmd.Flags().StringVarP(&output, "output-filename", "o", "output.txt", "Specify output filename")
	rootCmd.Flags().Int32VarP(&rangeFrom, "range-from", "f", math.MinInt32, "Specify lower limit for number range")
	rootCmd.Flags().Int32VarP(&rangeTo, "range-to", "t", math.MaxInt32, "Specify upper limit for number range")
}

func run(args []string) error {
	rowCount, rowErr := strconv.ParseUint(args[0], 0, 64)
	colCount, colErr := strconv.ParseUint(args[1], 0, 64)
	if rowErr != nil || colErr != nil {
		return fmt.Errorf("failed to parse row count (was '%s') or col count (was '%s') - both values must be parsable unsigned 64 bit integers", args[0], args[1])
	}

	generator, err := app.NewNumberGenerator(rangeFrom, rangeTo, delimiter)
	if err != nil {
		return err
	}

	fmt.Println("Starting generation")
	startTime := time.Now()

	app.WriteMatrixToFile(generator, uint(rowCount), uint(colCount), output)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("Generation took %d ms to complete", elapsedTime.Milliseconds())

	return nil
}
