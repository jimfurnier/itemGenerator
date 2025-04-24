package main

import (
	"compress/gzip"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	sizeTypes   = []string{"regular", "petite", "plus"}
	sizeSystems = []string{"US", "UK", "EU"}
	ageGroups   = []string{"newborn", "infant", "toddler", "kids", "adult"}
	availabilities = []string{"in stock", "out of stock", "preorder"}
	conditions  = []string{"new", "used", "refurbished"}
	genders     = []string{"male", "female", "unisex"}
	sizes       = []string{"S", "M", "L", "XL", "XXL"}
)

func randomFloat(min, max float64) string {
	return fmt.Sprintf("%.2f", min+rand.Float64()*(max-min))
}

func randomChoice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}

func randomUUID() string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", rand.Uint32(), rand.Uint32()>>16, rand.Uint32()>>16, rand.Uint32()>>16, rand.Uint64())
}

func main() {
	var rowCount int
	var outputPath string
	var compress bool

	flag.IntVar(&rowCount, "rows", 1000, "Number of rows to generate")
	flag.StringVar(&outputPath, "out", "products.csv", "Output file path")
	flag.BoolVar(&compress, "gzip", false, "Compress output using gzip")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var writer *csv.Writer
	if compress {
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		writer = csv.NewWriter(gzWriter)
	} else {
		writer = csv.NewWriter(file)
	}
	defer writer.Flush()

	columns := []string{
		"id", "additional_image_link", "adwords_grouping", "adwords_labels", "adwords_redirect",
		"age_group", "availability", "brand", "color", "condition", "custom_label_0",
		"custom_label_1", "custom_label_2", "custom_label_3", "custom_label_4",
		"description", "display_ads_title", "gender", "google_product_category", "gtin",
		"image_link", "item_group_id", "link", "material", "mobile_link", "mpn",
		"pattern", "price", "product_type", "promotion_id", "sale_price",
		"sale_price_effective_date", "shipping", "shipping_weight", "size", "size_system",
		"size_type", "tax", "title",
	}
	writer.Write(columns)

	for i := 1; i <= rowCount; i++ {
		row := []string{
			strconv.Itoa(i),
			"https://example.com/image.jpg",
			"group" + strconv.Itoa(rand.Intn(10)),
			"label" + strconv.Itoa(rand.Intn(10)),
			"https://redirect.com",
			randomChoice(ageGroups),
			randomChoice(availabilities),
			"Brand" + strconv.Itoa(rand.Intn(100)),
			"Color" + strconv.Itoa(rand.Intn(10)),
			randomChoice(conditions),
			"label0", "label1", "label2", "label3", "label4",
			"Simple product description",
			"Ad Title " + strconv.Itoa(i),
			randomChoice(genders),
			"Category" + strconv.Itoa(rand.Intn(100)),
			strconv.Itoa(rand.Intn(999999999999)),
			"https://example.com/img.jpg",
			randomUUID(),
			"https://product.com/item",
			"cotton",
			"https://m.example.com/item",
			fmt.Sprintf("MPN-%05d", rand.Intn(99999)),
			"pattern" + strconv.Itoa(rand.Intn(10)),
			randomFloat(10, 1000),
			"type" + strconv.Itoa(rand.Intn(10)),
			randomUUID(),
			randomFloat(5, 900),
			time.Now().Format("2006-01-02T15:04:05") + "/" + time.Now().Add(time.Hour*24).Format("2006-01-02T15:04:05"),
			randomFloat(0, 20),
			randomFloat(0.1, 5.0),
			randomChoice(sizes),
			randomChoice(sizeSystems),
			randomChoice(sizeTypes),
			randomFloat(0, 10),
			"Product Title " + strconv.Itoa(i),
		}
		writer.Write(row)
	}
}
