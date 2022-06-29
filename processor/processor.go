package processor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tonyxu1/kohomework/model"
)

var (
	conf             model.Config
	outputFile       *os.File
	err              error
	transactionList  model.TransactionList = make(map[string]model.WeeklyTransactionEntry)
	transactionEntry model.AccountEntry
)

func Process() {
	if err := conf.GetValue(); err != nil {
		log.Fatalln(err)
	}

	var startDate time.Time
	outputToConsole := true
	if conf.OutputFile != "" {
		outputToConsole = false
		outputFile, err = os.Create(conf.OutputFile)
		if err != nil {
			log.Fatalln(err)
		}
		defer outputFile.Close()
	}

	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer inputFile.Close()

	// transactionList = make(map[string]model.WeeklyTransactionEntry)

	//Read the first line from input file to get the start date
	// Assumption : input.txt contains more than 1 line of data.
	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	entryData := strings.ReplaceAll(scanner.Text(), "$", "")

	err = json.Unmarshal([]byte(entryData), &transactionEntry)
	if err != nil {
		log.Fatalln("Cannot unmarshal input data to object", err)
	}

	startDate = transactionEntry.LoadTime

	processEntry(transactionEntry, !outputToConsole, outputFile)

	for scanner.Scan() {
		entryData = strings.ReplaceAll(scanner.Text(), "$", "")
		err = json.Unmarshal([]byte(entryData), &transactionEntry)
		if err != nil {
			log.Fatalln("Cannot unmarshal input data to object", err)
		}

		if transactionList.IsDupTransaction(transactionEntry.ID, transactionEntry.CustomerID) {
			continue
		}
		loadDate := transactionEntry.LoadTime
		if foundMonday(startDate, loadDate) {
			startDate = loadDate
			transactionList.Reset()
		}
		processEntry(transactionEntry, !outputToConsole, outputFile)
	}

	log.Println("done.")

}

func writeOutput(writeToFile bool, f *os.File, text string) {
	if writeToFile {
		fmt.Fprintln(f, text)
		return
	}
	fmt.Println(text)

}

func processEntry(entry model.AccountEntry, writeToFile bool, OutputFile *os.File) {
	transactionList.Update(entry)
	t := transactionList.CreateOutput(entry, conf.DailyMaxAmout, conf.WeeklyMaxAmount, conf.DailyMaxCount)
	writeOutput(writeToFile, outputFile, t)

}

//
func foundMonday(startTime, endTime time.Time) bool {
	if !startTime.Before(endTime) {
		return false
	}
	t := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.UTC)
	for {

		t = t.AddDate(0, 0, 1)
		if t.After(endTime) {
			return false
		}
		if t.Weekday() == time.Monday {
			return true
		}
	}
}
