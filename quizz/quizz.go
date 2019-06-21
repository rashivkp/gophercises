package main

import ("fmt"
		"os"
		"encoding/csv"
		"io"
		"log"
		"bufio"
		"strings"
		"flag"
	)

func main() {
	fmt.Println("here starts the quizz")	
	max_duration := flag.String("t", "", "Quiz max duration") 
	flag.Parse()
	if *max_duration == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}



	fmt.Println("Quizz Duration:", *max_duration, "\nPress Enter to start")
	reader := bufio.NewReader(os.Stdin)
	right_answer := 0
	questions_count := 0
	reader.ReadString('\n')

	// qstn := "question,answer"
	csv_file, _ := os.Open("problems.csv")
	r := csv.NewReader(csv_file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		question := record[0]
		answer := record[1]
		fmt.Println(question)
		user_answer, _ := reader.ReadString('\n')
		user_answer = strings.Trim(user_answer, "\n")
		fmt.Println(user_answer,)
		if user_answer == answer {
			right_answer = right_answer + 1
		}
		questions_count = questions_count + 1
	}
	score := right_answer*100/questions_count
	fmt.Println("Total Questions", questions_count)
	fmt.Println("Total Right Answers ", right_answer)
	fmt.Println("Your score is", score)
}
