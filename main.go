package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func main() {
	router := gin.New()

	router.LoadHTMLGlob("view.html")

	router.GET("/", getInfo)
	router.GET("/get")
	router.Run()

}

func getInfo(context *gin.Context) {
	result := ""
	url := "https://store.steampowered.com/app/975370/Dwarf_Fortress/"
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.content", func(div *colly.HTMLElement) {
		div.ForEach("span", func(index int, element *colly.HTMLElement) {
			if index == 1 {
				result = element.Text
			}
		})
	})

	c.Visit(url)

	if result != "time is subjective" {
		result = "yes"
	}

	context.HTML(200, "view.html", gin.H{"result": result})
}

func getJSON(context *gin.Context) {
	result := ""
	url := "https://store.steampowered.com/app/975370/Dwarf_Fortress/"
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.content", func(div *colly.HTMLElement) {
		div.ForEach("span", func(index int, element *colly.HTMLElement) {
			if index == 1 {
				result = element.Text
			}
		})
	})

	c.Visit(url)

	if result != "time is subjective" {
		result = "yes"
	}

	context.JSON(200, result)
}
