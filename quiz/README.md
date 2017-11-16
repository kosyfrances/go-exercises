# Quiz
Exercise details from https://gophercises.com/exercises/quiz
### Part 1

Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to problems.csv, but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.
```
5+5,10
7+3,10
1+1,2
8+3,11
```
You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

### Part 2
... Coming soon
