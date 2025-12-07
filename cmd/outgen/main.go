/*
  Filename: main.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Main entry point for the outgen utility.
*/

package main

import (
	"os"

	"github.com/infologicmgmt/outgen/internal/config"
	"github.com/infologicmgmt/outgen/internal/log"
	"github.com/infologicmgmt/outgen/internal/processor"
	"github.com/spf13/cobra"
)

var (
	debug      bool
	verbose    bool
	quiet      bool
	configFile string
	workers    int
	stdin      bool
	format     string
	overwrite  bool
)

var rootCmd = &cobra.Command{
	Use:   "outgen",
	Short: "outgen is a command-line utility for generating files from templates.",
	Long: `outgen is a flexible command-line utility that reads template files 
and generates output files. It supports various template engines and allows for 
value substitution from configuration files.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logLevel := "error"
		if debug {
			logLevel = "debug"
		} else if verbose {
			logLevel = "verbose"
		} else if quiet {
			logLevel = "quiet"
		}
		log.InitLogger(logLevel)
	},
	Run: func(cmd *cobra.Command, args []string) {
		var cfg map[string]interface{}
		var err error

		if configFile != "" {
			cfg, err = config.LoadConfig(configFile)
			if err != nil {
				log.Logger.Fatal().Err(err).Msg("Failed to load config file")
			}
		}

		if stdin {
			job := processor.Job{
				InputFile:  "stdin",
				OutputFile: "stdout",
				Format:     format,
				Overwrite:  overwrite,
				Config:     cfg,
			}
			if err := job.Process(); err != nil {
				log.Logger.Fatal().Err(err).Msg("Failed to process stdin")
			}
			return
		}

		if len(args) == 0 {
			log.Logger.Fatal().Msg("No input files specified.")
		}

		jobs := make([]processor.Job, 0, len(args))
		for _, inputFile := range args {
			outputFile := processor.GetOutputFile(inputFile)
			jobs = append(jobs, processor.Job{
				InputFile:  inputFile,
				OutputFile: outputFile,
				Format:     format,
				Overwrite:  overwrite,
				Config:     cfg,
			})
		}
		processor.Run(jobs, workers)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging.")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose logging.")
	rootCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "Suppress all output except for fatal errors.")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Configuration file (YAML or JSON).")
	rootCmd.PersistentFlags().IntVarP(&workers, "workers", "w", 1, "Number of concurrent workers.")
	rootCmd.PersistentFlags().BoolVar(&stdin, "stdin", false, "Read from stdin and write to stdout.")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "", "Template format (e.g., jinja, mustache).")
	rootCmd.PersistentFlags().BoolVar(&overwrite, "overwrite", false, "Overwrite existing output files.")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Fatal().Err(err).Msg("Failed to execute command.")
		os.Exit(1)
	}
}
