package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"

	"github.com/gocolly/colly/v2"
)

type PageInfo struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func main() {
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	//Create JSON lines file
	file, err := os.Create("items.jl")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer file.Close()

	// Create a new collector instance
	c := colly.NewCollector()

	c.OnHTML("h1#firstHeading", func(e *colly.HTMLElement) {
		e.Request.Ctx.Put("title", e.Text)
	})

	c.OnHTML("#bodyContent", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()
		title := e.Request.Ctx.Get("title")
		text := stripHTML(e.Text) // Use the stripHTML function here

		// Create a PageInfo object
		pageInfo := PageInfo{URL: url, Title: title, Text: text}

		err := writeJSON(pageInfo, file) // Use the encapsulated function
		if err != nil {
			log.Printf("Failed to write to file: %s", err)
		}
	})

	// Start scraping the URLs
	for _, url := range urls {
		c.Visit(url)
	}
}

// JSON writing function
func writeJSON(info PageInfo, file *os.File) error {
	jsonData, err := json.Marshal(info)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(jsonData) + "\n")
	return err
}

// Use RegEx to strip HTML tags
func stripHTML(content string) string {
	// Regular expression to match HTML tags
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(content, "")
}
