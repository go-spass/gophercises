package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const defaultProblemsFile = "problems.csv"

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
func runQuiz(quiz []problem) int {
	numCorrect := 0
	scanner := bufio.NewScanner(os.Stdin) // Create a scanner to read from stdin
	for _, problem := range quiz {
		fmt.Printf("Question: %s\n", problem.question)
		fmt.Print("Answer: ")
		scanner.Scan()
		answer := scanner.Text()
		if answer == problem.answer {
			numCorrect++
		}
	}
	return numCorrect
}

func main() {
	fmt.Println("Welcome to the quiz game!")

	// Process Command Line Flags
	var filename string
	var timeout int
	flag.StringVar(&filename, "csv", defaultProblemsFile, "a csv file in the format of 'question,answer'")
	flag.IntVar(&timeout, "timeout", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	fmt.Printf("Using problems file: %s\n", filename)
	fmt.Printf("Using timeout: %d\n", timeout)

	// Load the quiz
	quiz, err := loadQuiz(filename)
	if err != nil {
		log.Fatalf("Failed to load the quiz: %s\n", err)
	}

	// Start the quiz
	numCorrect := runQuiz(quiz)
	fmt.Printf("You got %d out of %d correct\n", numCorrect, len(quiz))
}
