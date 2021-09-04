package main

import (
	"chanhyun/scrapper"
	"fmt"
	"os"
	"strconv"
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

// func main() {

// 	e := echo.New()
// 	e.GET("/", handleHome)
// 	e.POST("/scrape", handleScrape)
// 	e.Logger.Fatal(e.Start(":1323"))
// 	//scrapper.Scrape("term")
// }

var i int = 42

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter int = 0
)

func main() {

	//should be readable
	//var theHTTP string = "https://google.com"

	//shadowing
	var i float32 = 42.5
	//	j := 13

	var j int
	j = int(i)
	var c string = strconv.Itoa(j)

	fmt.Printf("%v,%T\n", c, c)
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T", j, j)
	//fmt.Println(i)
	//fmt.Println(j)
}
