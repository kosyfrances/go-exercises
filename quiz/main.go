package main

import (
        "os"
        "encoding/csv"
        "fmt"
        "log"
        "flag"
        "bufio"
        "strings"
)


func main() {

    csv_file := flag.String("file", "problems.csv", "a csv file in the format of 'question,answer'")
    flag.Parse()

    file, err := os.Open(*csv_file)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    quiz, err := csv.NewReader(file).ReadAll()

    if err != nil {
        log.Fatal(err)
    }

    quiz_length := len(quiz)
    score := 0

    for index, value := range quiz {

        reader := bufio.NewReader(os.Stdin)
        fmt.Printf("Problem #%d: %s ", index, value[0])
        answer, _ := reader.ReadString('\n')

        if strings.TrimSpace(answer) == strings.TrimSpace(value[1]) {
            score ++
        }
    }

    fmt.Printf("You scored %d out of %d", score, quiz_length)
}
