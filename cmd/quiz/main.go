package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const defaultProblemsFile = "problems.csv"

type App struct {
	filename   string
	duration   int
	doneChan   chan bool
	answerChan chan bool
}

func newApp() App {
	app := App{
		filename: defaultProblemsFile,
		duration: 30,
	}

	// Process Command Line Flags
	var filename string
	var timeout int
	flag.StringVar(&filename, "csv", defaultProblemsFile, "a csv file in the format of 'question,answer'")
	flag.IntVar(&timeout, "timeout", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	app.filename = filename
	app.duration = timeout

	// Create channels
	app.doneChan = make(chan bool)
	app.answerChan = make(chan bool)

	return app
}

type problem struct {
	question string
	answer   string
}

func loadQuiz(filename string) ([]problem, error) {
	// Open the problems file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", err)
		return nil, err
	}
	defer file.Close()

	// Read the problems file
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read the CSV file: %s\n", err)
		return nil, err
	}

	// Create the quiz
	quiz := []problem{}
	for _, record := range records {
		quiz = append(quiz, problem{question: record[0], answer: strings.TrimSpace(record[1])})
	}
	return quiz, nil
}

// runQuiz runs the quiz and returns the number of correct answers
func runQuiz(app App, quiz []problem) {
	scanner := bufio.NewScanner(os.Stdin) // Create a scanner to read from stdin
	for _, problem := range quiz {
		fmt.Printf("Question: %s\n", problem.question)
		fmt.Print("Answer: ")
		scanner.Scan()
		answer := scanner.Text()
		if answer == problem.answer {
			app.answerChan <- true
		}
	}
	app.doneChan <- true
}

func main() {
	app := newApp() // Create a new app and handle command line flags
	fmt.Println("Welcome to the quiz game!")

	// Load the quiz
	quiz, err := loadQuiz(app.filename)
	if err != nil {
		log.Fatalf("Failed to load the quiz: %s\n", err)
	}

	numCorrect := 0

	// Start the quiz
	go runQuiz(app, quiz)

	// Start the timer
	go func() {
		time.Sleep(30 * time.Second)
		fmt.Println("\nTimes Up!")
		app.doneChan <- true
	}()

	// Wait for the quiz to finish or the timer to expire
	for {
		select {
		case <-app.doneChan:
			fmt.Printf("You got %d out of %d correct\n", numCorrect, len(quiz))
			os.Exit(0)
		case <-app.answerChan:
			numCorrect++
		}
	}
}
