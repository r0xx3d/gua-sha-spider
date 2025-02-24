package main

import (
	"github.com/gocolly/colly"
	"encoding/csv"
	"os"
	"log"
)

type Product struct {
	Url, Image, Name, Price string
}


func main() {

	var products []Product

	//...implementation of scraping logic to be done

	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)


	// triggered when a CSS selector matches an element
	c.OnHTML("li.product", func(e *colly.HTMLElement) {

		product := Product{}

		product.Url = e.ChildAttr("a", "href")
		product.Image = e.ChildAttr("img", "src")
		product.Name = e.ChildText(".product-name")
		product.Price = e.ChildText(".price")

		products = append(products, product)
	})


	// triggered once scraping is done
	c.OnScraped(func(r *colly.Response) {
		// save to csv

		file, err := os.Create("products.csv")
		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}
		defer file.Close()

		// initialize file writer
		writer := csv.NewWriter(file)

		// write the CSV headers
		headers := []string{
			"Url",
			"Image",
			"Name",
			"Price",
		}
		writer.Write(headers)

		for _, product := range products {
			// convert a Product to an array of strings
			record := []string {
				product.Url,
				product.Image,
				product.Name,
				product.Price,
			}
			writer.Write(record)
		}
		defer writer.Flush()
	})

	c.Visit("https://www.scrapingcourse.com/ecommerce")


}

