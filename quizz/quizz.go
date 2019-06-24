package main

import (
    "fmt"
    "os"
    "encoding/csv"
    "log"
    "bufio"
    "strings"
    "flag"
    "time"
)

var right_answer, questions_count, seconds int = 0, 0, 0

func countdown() {
    for i := 0; i < seconds; i++ {
        time.Sleep(1 * time.Second)
    }
    fmt.Println("\nScored:", right_answer, "out of ", questions_count)
    os.Exit(0)
}

func main() {
    fmt.Println("here starts the quizz")    
    flag.IntVar(&seconds, "t", 30, "Quiz duration in seconds") 
	flag.Parse()

    fmt.Println("Quizz Duration:", seconds, "\nPress Enter to start")
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
    go countdown()

    csv_file, _ := os.Open("problems.csv")
    r := csv.NewReader(csv_file)

    records, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    questions_count = len(records)

    for i:=0;i<questions_count;i++ {
        question := records[i][0]
        answer := records[i][1]
        fmt.Println(question)
        user_answer, _ := reader.ReadString('\n')
        user_answer = strings.Trim(user_answer, "\n")
        if user_answer == answer {
            right_answer = right_answer + 1
        }
    }
    fmt.Println("\nScored:", right_answer, "out of ", questions_count)
    return
}
