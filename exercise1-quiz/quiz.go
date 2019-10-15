package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"bufio"
	"strings"
	"strconv"
	"time"
)

func main(){
	data, totalQuestion := parseQuestionCSV()
	var correctAnswer uint8
	doneChan := make(chan bool)
	timeoutChan := make(chan bool)
	go func(){
		for _, questionSets := range data{
			answerInt := askQuestion(questionSets)
			answerFromUserInt := answerQuestion()
			isTimeout := checkTimeout(timeoutChan)
			if isTimeout{
				return
			}
			validateAnswer(answerInt, answerFromUserInt, &correctAnswer)
		}
		doneChan<-true
	}()
	blockingForTimeout(timeoutChan, doneChan)
	fmt.Println("Total Questions: ", totalQuestion, ", Correct Answer: ", correctAnswer)
}

func checkTimeout(timeoutChan <-chan bool)(timeout bool){
	select {
		case <- timeoutChan:
			timeout = true
			fmt.Println("timeout")
			return
		default:
			timeout = false
	}
	return
}

func parseQuestionCSV()(data [][]string, totalQuestion uint8){
	dir, err := os.Getwd()
	if err != nil{
		fmt.Println(err)
		return
	}
	csvPath := fmt.Sprintf("%s\\%s", dir, "problems.csv")
	f, err := os.Open(csvPath)
	if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
	reader := csv.NewReader(f)
	data, err = reader.ReadAll()
    if err != nil {
        fmt.Println(err)
        os.Exit(-2)
    }
	totalQuestion = uint8(len(data))
	return
}

func askQuestion(questionSets []string)(answerInt uint64){
	question, answer := questionSets[0], questionSets[1]
	fmt.Println("question: ", question)
	answerInt, err := strconv.ParseUint(answer, 10, 8)
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	fmt.Print("Enter answer: ")
	return
}

func answerQuestion()(answerFromUserInt uint64){
	stdinReader := bufio.NewReader(os.Stdin)
	answerFromUser, _ := stdinReader.ReadString('\n')
	answerFromUser = strings.Trim(answerFromUser, "\r\n")
	answerFromUserInt, err := strconv.ParseUint(answerFromUser, 10, 8)
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	return
}

func validateAnswer(answer uint64, answerFromUser uint64, correctAnswer *uint8){
	if answerFromUser == answer{
		fmt.Println("Correct answer!")
		*correctAnswer++
	}else{
		fmt.Println(fmt.Sprintf("Wrong answer, correct answer is %d", answer))
	}
}

func blockingForTimeout(timeoutChan chan<- bool, doneChan <-chan bool){
	select {
    case <-doneChan:
		fmt.Println("All Questions Done")
    case <-time.After(time.Second * 3):
		timeoutChan<-true
	}
	return
}