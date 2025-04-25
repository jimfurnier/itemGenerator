package generator

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jimfurnier/itemGenerator/config"
)

type Generator struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Generator {
	gofakeit.Seed(time.Now().UnixNano())
	return &Generator{cfg}
}

func (g *Generator) GenerateRow(id int) []string {
	row := make([]string, len(g.cfg.Columns))
	for i, col := range g.cfg.Columns {
		switch col.Type {
		case "string":
			row[i] = gofakeit.Word()
		case "integer":
			row[i] = strconv.Itoa(rand.Intn(100000))
		case "float":
			row[i] = fmt.Sprintf("%.2f", gofakeit.Price(1, 1000))
		case "name":
			row[i] = gofakeit.Name()
		case "email":
			row[i] = gofakeit.Email()
		case "datetime":
			row[i] = gofakeit.Date().Format(time.RFC3339)
		default:
			row[i] = "unknown"
		}
	}
	return row
}

