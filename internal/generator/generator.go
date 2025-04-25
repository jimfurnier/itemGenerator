package generator

import (
	"fmt"
	"github.com/jimfurnier/itemGenerator/internal/config"
	"math/rand"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Generator struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Generator {
	gofakeit.Seed(time.Now().UnixNano())
	return &Generator{cfg}
}

func (g *Generator) GenerateRow(id int) []string {
	row := make([]string, len(g.cfg.GetColumns()))
	for i, col := range g.cfg.GetColumns() {
		if col.Optional < config.OptionalDefault {
			if rand.Intn(config.OptionalDefault) >= col.Optional {
				row[i] = ""
				continue
			}
		}

		switch col.Type {
		case "rowNumber":
			row[i] = strconv.Itoa(id)
		case "string":
			row[i] = gofakeit.Word()
		case "sentence":
			row[i] = gofakeit.Sentence(rand.Intn(10))
		case "integer":
			row[i] = strconv.Itoa(rand.Intn(100000))
		case "float":
			row[i] = fmt.Sprintf("%.2f", gofakeit.Price(1, 1000))
		case "percent":
			row[i] = fmt.Sprintf("%.2f%%", gofakeit.Price(1, 1000)/100)
		case "name":
			row[i] = gofakeit.Name()
		case "email":
			row[i] = gofakeit.Email()
		case "datetime":
			row[i] = gofakeit.Date().Format(time.RFC3339)
		case "image":
			row[i] = gofakeit.URL()
		case "url":
			row[i] = gofakeit.URL()
		case "gender":
			row[i] = gofakeit.Gender()
		case "color":
			row[i] = gofakeit.Color()
		case "uuid":
			row[i] = gofakeit.UUID()
		case "product":
			row[i] = gofakeit.ProductName()
		case "category":
			row[i] = gofakeit.ProductCategory()
		case "description":
			row[i] = gofakeit.ProductDescription()
		case "material":
			row[i] = gofakeit.ProductMaterial()
		case "null":
			row[i] = ""
		case "random":
			if len(col.Random) > 0 {
				row[i] = pickRandom(col.Random)
			} else {
				row[i] = ""
			}
		default:
			row[i] = "unknown"
		}
	}
	return row
}

func pickRandom(options []string) string {
	if len(options) == 0 {
		return ""
	}
	return options[rand.Intn(len(options))]
}
