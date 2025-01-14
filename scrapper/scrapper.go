/*
21.05.23
Learning GO language with nomad code

Using simple print function  (Hello World!)


add something package (21.05.23)


learning const variable & nomral variable (21.05.23)

add comment for github third (21.05.23)


21.05.27 Functions part One
--Naked Function
-- Defer

21.05.27 for loop
- range


21.05.29
- if else | switch
- Pointer in GO

21.05.31
- Array and slices
- Map
- Struct


21.06.02 [Bank & Dictionary projects]
- Account + NewAccount


21.06.03 [Dictionary]


21.06.07  [URL checker project]
- url check (url hit)
- using map

21.06.09 [parallel]
- 'go' routine
- 'channel' variable
- get multi messages with channel variable


- using go routine & channel with HITURL project

*******************[Bank & Dictionary projects] ****************


21.06.13 [Scrapper projects]
- download go query from relative github
-- add getPages - 21.06.23
-- extract
-- code review 21.06.30

21.06.22 renew go path


21.07.03 [WEB SERVER WITH ECHO]

- 21.07.04 using echo .GET, POST

-21.07.08 download files






*/

//package main
package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	salary   string
	location string
	summary  string
}

//Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50" //https://kr.indeed.com/jobs?q=python&limit=50"
	var jobs []extractedJob
	c := make(chan []extractedJob)

	//fmt.Println("hi")
	totalPages := getPages(baseURL)
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		//how many go routines?
		go getPage(i, baseURL, c)

	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	//fmt.Println(jobs[0])
	writeJobs(jobs)
	fmt.Println("Done, extracted:", len(jobs))
} //end of main

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) //bytes
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	//create 50 go routine
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
		//jobs = append(jobs, job)
	})

	for i := 0; i < (searchCards.Length()); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()
	headers := []string{"LINK", "TITLE", "LOCATION", "SALARY", "SUMMARY"}
	wErr := w.Write(headers)

	checkErr(wErr)
	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}

}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk") //job card struct

	title := CleanString(card.Find(".title>a").Text())

	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())

	//receiving messages
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
	//fmt.Println(id, title, location, salary, summary)
}

//how many pages are there
func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)

	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) //bytes
	checkErr(err)
	//fmt.Println(doc)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Html())
		fmt.Println(s.Find("a").Html()) // just get first one
		pages = (s.Find("a").Length())  // get 5 link (Knowing how many pages)

	})
	return pages

}

//Checking error all the time

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status: ", res.StatusCode)
	}
}

//clean string
func CleanString(str string) string {
	// no make empty space
	//양쪽끝 스페이스 없애고,
	//추가로 스페이스 또 없앤 후, 배열안에 텍스트만 얻음
	//example hello.              f.        1-> (trimspace로 공백사라지고, fields로 배열이 나오고, "hello.","f.","1." -> join으로 hello f 1
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")

}
