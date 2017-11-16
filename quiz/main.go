package main

import (
        "os"
        "encoding/csv"
        "fmt"
        "log"
        "flag"
)


func main() {

    var csv_file = flag.String("file", "problems.csv", "a csv file in the format of 'question,answer'")
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

    fmt.Print(quiz)
}

