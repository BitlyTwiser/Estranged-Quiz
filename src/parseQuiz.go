package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

//Set Score globally
var scoreNums score

func main() {
	fmt.Println("You will have 30 seconds to complete this quiz.")
	fmt.Println("Press enter to start quiz: ")
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	if char != '\x00' {
		fmt.Printf("You must press enter to being.. you pressed %q\r\n restarting\n", char)
		main()
	} else {
		primaryExec()
	}
	//primaryExec()
}

func primaryExec() {
	//Core for loop performing all primary actions
	//Intialize vars
	dataLoc := "/home/smallz/go/src/gophercises/quiz/problems.csv"
	dataValue := openFile(dataLoc)
	go timer()
	for i := 0; i < len(dataValue); i++ {
		var answer int64
		fmt.Println("Question: What is the answer to the addition of the two followig values: ", dataValue[i])
		_, err := fmt.Scan(&answer)
		if err != nil {
			fmt.Printf("Ahem, something happened.. Closing program.\nError: %s", err.Error())
			os.Exit(0)
		}
		if addValues(dataValue[i]) == answer {
			scoreNums.correct++
		} else {
			scoreNums.incorrect++
		}
	}
	//Just printing the results
	fmt.Println("Quiz is done")
	fmt.Println("Correct: ", scoreNums.correct)
	fmt.Println("Incorrect: ", scoreNums.incorrect)
}

type score struct {
	correct   int
	incorrect int
}

func openFile(dataLoc string) []string {
	var myReturn []string
	if _, err := os.Stat(dataLoc); os.IsNotExist(err) {
		log.Fatal("Object on given path does not exist")
	}
	file, err := os.Open(dataLoc)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Println("File found, reading file..")
	in := csv.NewReader(file)
	for {
		record, err := in.Read()
		if err == io.EOF {
			log.Println("EOF reached.. moving on")
			break
		}
		if err != nil {
			log.Fatal(err)
		} else {
			myReturn = append(myReturn, record[0])
		}
	}
	return myReturn
}

func addValues(valuesArray string) int64 {
	string1, _ := strconv.ParseInt(strings.Split(valuesArray, "+")[0], 0, 64)
	string2, _ := strconv.ParseInt(strings.Split(valuesArray, "+")[1], 0, 64)
	return string1 + string2
}

func timer() {
	time.Sleep(time.Second * 30)
	fmt.Println("Time value limit has been reached. Quiz is over!")
	fmt.Println("Correct: ", scoreNums.correct)
	fmt.Println("Incorrect: ", scoreNums.incorrect)
	os.Exit(0)
}
