package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	//return c.String(http.StatusOK, "hello,world!")
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	//term := strings.Tolower(scrapper.CleanString(c.FormValue("term")))
	fmt.Println(c.FormValue("term"))
	return nil
	//return c.File("home.html")
}

func main() {

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	//scrapper.Scrape("term")
}
