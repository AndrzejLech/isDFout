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
	router.GET("/get", getJSON)
	router.Run()

}

func getInfo(context *gin.Context) {
	result := visitDwarfFortress()

	context.HTML(200, "view.html", gin.H{"result": result})
}

func getJSON(context *gin.Context) {

	result := visitDwarfFortress()

	context.JSON(200, result)
}

func visitDwarfFortress() string {
	result := ""
	url := "https://store.steampowered.com/app/975370/Dwarf_Fortress/"
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.date", func(div *colly.HTMLElement) {
		result = div.Text
	})

	c.Visit(url)

	return result
}
