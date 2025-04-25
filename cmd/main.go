package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jimfurnier/itemGenerator/internal/app"
	"github.com/jimfurnier/itemGenerator/internal/config"
)

func main() {
	rows := flag.Int("rows", 100, "Number of rows to generate")
	templatePath := flag.String("template", "", "Path to template JSON file")
	compression := flag.String("compression", "", "Force the compress type, regardless of the template")
	flag.Parse()

	if *templatePath == "" {
		fmt.Println("Template path is required")
		os.Exit(1)
	}

	cfg, err := config.LoadFromJsonTemplate(*templatePath, *rows)
	if err != nil {
		panic(err)
	}
	if *compression != "" {
		fmt.Printf("Forcing compression: %s\n", *compression)
		cfg.ForceCompression(*compression)
	}

	result, err := app.Execute(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %d rows to %s\n", result.Count(), result.Path())
}
