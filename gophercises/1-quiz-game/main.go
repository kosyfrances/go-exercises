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
    "math/rand"
)

func cleanStrings(chars string) string {
    return strings.ToUpper(strings.TrimSpace(chars))
}

type QuizDetail struct {
    quiz       [][]string
    score      int
    quizLength int
}

func shuffleQuiz(quiz [][]string) [][]string {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    shuffledQuizArray := make([][]string, len(quiz))
    perm := r.Perm(len(quiz))

    for index, randIndex := range perm {
        shuffledQuizArray[index] = quiz[randIndex]
    }
    return shuffledQuizArray
}

func loopQuiz(quizDetail *QuizDetail, timeout chan bool, shuffle bool) {
    quizDetail.quizLength = len(quizDetail.quiz)
    quizDetail.score = 0

    var quiz [][] string

    if shuffle {
        quiz = shuffleQuiz(quizDetail.quiz)
    } else {
        quiz = quizDetail.quiz
    }

    for index, value := range quiz {

        reader := bufio.NewReader(os.Stdin)
        fmt.Printf("Problem #%d: %s ", index+1, value[0])
        answer, _ := reader.ReadString('\n')

        if cleanStrings(answer) == cleanStrings(value[1]) {
            quizDetail.score++
        }
    }

    timeout <- true
}

func main() {

    csvFile := flag.String(
        "csv",
        "problems.csv",
        "A csv file in the format of 'question,answer'",
    )
    timeLimit := flag.Duration("limit", 30*time.Second, "Quiz time limit in seconds")
    shuffle := flag.Bool("shuffle", false, "Randomly shuffle quiz questions")
    flag.Parse()

    file, err := os.Open(*csvFile)

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

    go loopQuiz(quizDetail, quizEnded, *shuffle)
    <-quizEnded

    fmt.Printf("\nYou scored %d out of %d", quizDetail.score, quizDetail.quizLength)
}
