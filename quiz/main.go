package main

import (
        "os"
        "encoding/csv"
        "fmt"
        "log"
)


func main() {

    file, err := os.Open("problems.csv")

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    quiz, err := csv.NewReader(file).ReadAll()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(quiz)
}

