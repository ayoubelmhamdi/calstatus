package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"strings"
	"strconv"
	"sort"
	"encoding/json"
)

var (
	addFlag  = flag.String("add", "", "Add a new text")
	timeFlag = flag.String("time", "", "Time delay for the text to appear")
	atFlag   = flag.String("at", "", "Specific time for the text to appear")
	fileName = "calendar.json"
)

type Text struct {
	Time int64
	Text string
}

func main() {
	flag.Parse()

	if *addFlag != "" {
		if *atFlag != "" {
			addTextAt(*addFlag, *atFlag)
		} else if *timeFlag != "" {
			addText(*addFlag, *timeFlag)
		} else {
			fmt.Println("Error: Please provide either --time or --at flag")
			os.Exit(1)
		}
	} else {
		printText()
	}
}

func addText(text string, delay string) {
	date := time.Now()

	if delay != "" {
		duration, err := parseDuration(delay)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		date = date.Add(duration)
	}

	textWithDate := Text{Time: date.Unix(), Text: text}

	texts := readTexts()
	texts = append(texts, textWithDate)
	sortTexts(texts)
	writeTexts(texts)

	fmt.Println("Done")
}

func addTextAt(text string, at string) {
	location, _ := time.LoadLocation("Africa/Casablanca")
	date, err := time.ParseInLocation("2006-01-02 15:04", at, location)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	textWithDate := Text{Time: date.Unix(), Text: text}

	texts := readTexts()
	texts = append(texts, textWithDate)
	sortTexts(texts)
	writeTexts(texts)

	fmt.Println("Done")
}

func printText() {
	texts := readTexts()

	var nextText *Text
	for _, text := range texts {
		if text.Time > time.Now().Unix() {
			nextText = &text
			break
		}
	}

	if nextText != nil {
		fmt.Println(nextText.Text)
	}
}

func readTexts() []Text {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Text{}
		}
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	var texts []Text
	err = json.NewDecoder(file).Decode(&texts)
	if err != nil {
		if err.Error() == "EOF" {
			return []Text{}
		}
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return texts
}

func writeTexts(texts []Text) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(texts)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func sortTexts(texts []Text) {
	sort.Slice(texts, func(i, j int) bool {
		return texts[i].Time < texts[j].Time
	})
}

func parseDuration(s string) (time.Duration, error) {
	if len(s) < 2 {
		return 0, fmt.Errorf("invalid duration")
	}

	unit := s[len(s)-1:]
	value, err := strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return 0, fmt.Errorf("invalid duration")
	}

	switch strings.ToUpper(unit) {
	case "M":
		return time.Duration(value) * time.Minute, nil
	case "H":
		return time.Duration(value) * time.Hour, nil
	default:
		return 0, fmt.Errorf("invalid duration unit")
	}
}
