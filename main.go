package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jimfurnier/itemGenerator/config"
	"github.com/jimfurnier/itemGenerator/compressor"
	"github.com/jimfurnier/itemGenerator/generator"
	"github.com/jimfurnier/itemGenerator/writer"
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

	// Create writer based on delimiter
	w, err := writer.NewWriter(cfg)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	gen := generator.New(cfg)
	headers := cfg.GetHeaders()
	w.WriteHeader(headers)

	for i := 0; i < *rows; i++ {
		row := gen.GenerateRow(i)
		w.WriteRow(row)
	}

	if cfg.Compression != "none" {
		comp, err := compressor.GetCompressor(cfg.Compression)
		if err != nil {
			panic(err)
		}
		err = comp.Compress(cfg.OutputName+".csv", cfg.OutputName+".csv."+cfg.Compression)
		if err != nil {
			panic(err)
		}
		fmt.Println("File compressed to:", cfg.OutputName+".csv."+cfg.Compression)
	}
}

