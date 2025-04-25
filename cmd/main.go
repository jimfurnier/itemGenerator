package main

import (
	"flag"
	"fmt"
	"github.com/jimfurnier/itemGenerator/internal/config"
	"github.com/jimfurnier/itemGenerator/internal/generator"
	"github.com/jimfurnier/itemGenerator/internal/writer"
	"os"
)

func main() {
	rows := flag.Int("rows", 100, "Number of rows to generate")
	templatePath := flag.String("template", "", "Path to template JSON file")
	flag.Parse()

	if *templatePath == "" {
		fmt.Println("Template path is required")
		os.Exit(1)
	}

	cfg, err := config.LoadConfig(*templatePath)
	if err != nil {
		panic(err)
	}

	// Create writer spec from config
	spec, err := writer.NewWriteSpec(cfg, rows)
	if err != nil {
		panic(err)
	}

	fileWriter := writer.NewCompressedFileWriter(
		writer.NewDefaultFileWriter(generator.New(cfg)),
	)
	result, err := fileWriter.Write(spec)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %d rows to %s\n", result.Count(), result.Path())
}
