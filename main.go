package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

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
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	file, err := os.Create("items.jl")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer file.Close()

	c := colly.NewCollector()

	c.OnHTML("#bodyContent", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()
		title := e.Request.Ctx.Get("title")
		text := stripHTML(e.Text) // Use the stripHTML function here

		pageInfo := PageInfo{URL: url, Title: title, Text: text}

		err := writeJSON(pageInfo, file) // Use the encapsulated function
		if err != nil {
			log.Printf("Failed to write to file: %s", err)
		}
	})

	for _, url := range urls {
		c.Visit(url)
	}
}

func writeJSON(info PageInfo, file *os.File) error {
	jsonData, err := json.Marshal(info)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(jsonData) + "\n")
	return err
}

func stripHTML(content string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(content)), " ")
}
