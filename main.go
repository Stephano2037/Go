package main

import (
	"chanhyun/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const FILE_NAME string = "jobs.csv"

func handleHome(c echo.Context) error {
	//return c.String(http.StatusOK, "hello,world!")
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(FILE_NAME)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	//fmt.Println(c.FormValue("term"))
	scrapper.Scrape(term)

	return c.Attachment(FILE_NAME, FILE_NAME) //nil
	//return c.File("home.html")
}

func main() {

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	//scrapper.Scrape("term")
}
