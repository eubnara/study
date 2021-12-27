package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eubnara/study/go/gopl/problems/4.10/github"
)

func main() {
	fmt.Println("Within a month.")
	currentTime := time.Now()
	t := currentTime.AddDate(0, -1, 0)
	updated := fmt.Sprintf("updated:>=%s", t.Format("2006-01-02"))
	terms := append(os.Args[1:], updated)
	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\n\n\nWithin a year.\n")
	currentTime = time.Now()
	t = currentTime.AddDate(-1, 0, 0)
	updated = fmt.Sprintf("updated:>=%s", t.Format("2006-01-02"))
	terms = append(os.Args[1:], updated)
	result, err = github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\n\n\nMore than a year.\n")
	t = currentTime.AddDate(-1, 0, 0)
	updated = fmt.Sprintf("updated:<%s", t.Format("2006-01-02"))
	terms = append(os.Args[1:], updated)
	result, err = github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
