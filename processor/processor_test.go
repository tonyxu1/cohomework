package processor

import (
	"os"
	"testing"
	"time"

	"github.com/tonyxu1/kohomework/model"
)

// Variable for testing
var (
	accountEntry = model.AccountEntry{
		ID:         "10003",
		CustomerID: "1234",
		LoadAmount: 2000.00,
		LoadTime:   time.Now().AddDate(0, 0, 1),
	}
)

func Test_writeOutput(t *testing.T) {
	type args struct {
		writeToFile bool
		f           *os.File
		text        string
	}
	testOutputFile, _ := os.Create("./temp.txt")
	defer testOutputFile.Close()

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Write text to console",
			args: args{
				writeToFile: false,
				f:           testOutputFile,
				text:        "this string should be showed up in console",
			},
		},
		{
			name: "Write text to temp.txt in current folder",
			args: args{
				writeToFile: false,
				f:           testOutputFile,
				text:        "this string should be written to  temp.txt in current folder",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeOutput(tt.args.writeToFile, tt.args.f, tt.args.text)
		})
	}
}

func Test_processEntry(t *testing.T) {
	type args struct {
		entry       model.AccountEntry
		writeToFile bool
		OutputFile  *os.File
	}
	testOutputFile, _ := os.Create("./temp.txt")
	defer testOutputFile.Close()

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Process load entry unit test",
			args: args{
				entry:       accountEntry,
				writeToFile: true,
				OutputFile:  testOutputFile,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processEntry(tt.args.entry, tt.args.writeToFile, tt.args.OutputFile)
		})
	}
}

func Test_foundMonday(t *testing.T) {
	type args struct {
		startTime time.Time
		endTime   time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Found Monday testing",
			args: args{
				startTime: time.Now(),
				endTime:   time.Now().AddDate(0, 0, 8),
			},
			want: true,
		},
		{
			name: "Monday not found",
			args: args{
				startTime: time.Date(2020, 11, 16, 0, 0, 0, 0, time.UTC),
				endTime:   time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := foundMonday(tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("foundMonday() = %v, want %v", got, tt.want)
			}
		})
	}
}
