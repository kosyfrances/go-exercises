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

func cleanStrings(chars string) string {
    return strings.ToUpper(strings.TrimSpace(chars))
}

type QuizDetail struct {
    quiz       [][]string
    score      int
    quizLength int
}

func loopQuiz(quizDetail *QuizDetail, timeout chan bool) {
    quizDetail.quizLength = len(quizDetail.quiz)
    quizDetail.score = 0

    for index, value := range quizDetail.quiz {

        reader := bufio.NewReader(os.Stdin)
        fmt.Printf("Problem #%d: %s ", index, value[0])
        answer, _ := reader.ReadString('\n')

        if cleanStrings(answer) == cleanStrings(value[1]) {
            quizDetail.score++
        }
    }

    timeout <- true
}

func main() {

    csv_file := flag.String(
        "csv",
        "problems.csv",
        "a csv file in the format of 'question,answer'",
    )
    timeLimit := flag.Duration("limit", 30*time.Second, "quiz time limit in seconds")
    flag.Parse()

    file, err := os.Open(*csv_file)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    quizDetail := &QuizDetail{}

    quizDetail.quiz, err = csv.NewReader(file).ReadAll()

    if err != nil {
        log.Fatal(err)
    }

    quizEnded := make(chan bool)

    time.AfterFunc(*timeLimit, func() { quizEnded <- true })

    go loopQuiz(quizDetail, quizEnded)
    <-quizEnded

    fmt.Printf("\nYou scored %d out of %d", quizDetail.score, quizDetail.quizLength)
}
