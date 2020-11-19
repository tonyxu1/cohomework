package model

import (
	"testing"
	"time"
)

var (
	tl = map[string]WeeklyTransactionEntry{
		"1234": WeeklyTransactionEntry{
			ID: map[string]struct{}{
				"10001": struct{}{},
				"10002": struct{}{},
			},
			Transaction: map[time.Weekday]DailyTotal{
				time.Monday: DailyTotal{
					Amount: 4999.99,
					Count:  3,
				},
				time.Tuesday: DailyTotal{
					Amount: 4999.99,
					Count:  3,
				},
				time.Wednesday: DailyTotal{
					Amount: 4999.99,
					Count:  3,
				},
				time.Thursday: DailyTotal{
					Amount: 2000.99,
					Count:  3,
				},
				time.Friday: DailyTotal{
					Amount: 2.99,
					Count:  3,
				},
				time.Saturday: DailyTotal{
					Amount: 2.99,
					Count:  3,
				},
			},
		},
		"1235": WeeklyTransactionEntry{
			ID: map[string]struct{}{
				"20001": struct{}{},
				"20002": struct{}{},
			},
			Transaction: map[time.Weekday]DailyTotal{
				time.Monday: DailyTotal{
					Amount: 5000.01,
					Count:  1,
				},
				time.Tuesday: DailyTotal{
					Amount: 4999.99,
					Count:  3,
				},
				time.Friday: DailyTotal{
					Amount: 4999.99,
					Count:  3,
				},
			},
		},
	}
	dailyMaxAmount  float32 = 5000.00
	dailyMaxCount   int8    = 3
	weeklyMaxAmount float32 = 20000.00
)

func TestTransactionList_IsDupTransaction(t *testing.T) {
	type args struct {
		transactionID string
		customerID    string
	}

	tests := []struct {
		name string
		tl   TransactionList
		args args
		want bool
	}{
		// TODO: Add test cases.

		{
			name: "Not a duplicate load",
			tl:   tl,
			args: args{
				transactionID: "100003",
				customerID:    "1234",
			},
			want: false,
		},
		{
			name: "A duplicate load",
			tl:   tl,
			args: args{
				transactionID: "20001",
				customerID:    "1235",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tl.IsDupTransaction(tt.args.transactionID, tt.args.customerID); got != tt.want {
				t.Errorf("TransactionList.IsDupTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionList_canAccept(t *testing.T) {
	type args struct {
		customerID   string
		dailyAmount  float32
		weeklyAmount float32
		dailyCount   int8
	}
	tests := []struct {
		name string
		tl   TransactionList
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "List is valid still good for next load",
			tl:   tl,
			args: args{
				customerID:   "1234",
				dailyAmount:  dailyMaxAmount,
				weeklyAmount: weeklyMaxAmount,
				dailyCount:   dailyMaxCount,
			},
			want: true,
		},
		{
			name: "Customer cannot load any money",
			tl:   tl,
			args: args{
				customerID:   "1235",
				dailyAmount:  dailyMaxAmount,
				weeklyAmount: weeklyMaxAmount,
				dailyCount:   dailyMaxCount,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tl.canAccept(tt.args.customerID, tt.args.dailyAmount, tt.args.weeklyAmount, tt.args.dailyCount); got != tt.want {
				t.Errorf("TransactionList.canAccept() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionList_Update(t *testing.T) {
	type args struct {
		ae AccountEntry
	}
	tests := []struct {
		name string
		tl   TransactionList
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Amount update for weekday of today",
			tl:   tl,
			args: args{
				ae: AccountEntry{
					ID:         "10003",
					CustomerID: "1234",
					LoadAmount: 1000.00,
					LoadTime:   time.Now(),
				},
			},
		},
		{
			name: "Amount update for weekday of tomorrow",
			tl:   tl,
			args: args{
				ae: AccountEntry{
					ID:         "10003",
					CustomerID: "1234",
					LoadAmount: 1000.00,
					LoadTime:   time.Now().AddDate(0, 0, 1),
				},
			},
		},
		{
			name: "Amount update weekday of yesterday",
			tl:   tl,
			args: args{
				ae: AccountEntry{
					ID:         "10003",
					CustomerID: "1234",
					LoadAmount: 1000.00,
					LoadTime:   time.Now().AddDate(0, 0, -1),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalAmount := tt.tl[tt.args.ae.CustomerID].Transaction[tt.args.ae.LoadTime.Weekday()].Amount
			tt.tl.Update(tt.args.ae)
			updatedAmount := tt.tl[tt.args.ae.CustomerID].Transaction[tt.args.ae.LoadTime.Weekday()].Amount

			if (updatedAmount - originalAmount) != tt.args.ae.LoadAmount {
				t.Errorf("Load Amount update error: 'updatedAmount - originalAmount =  %.2f ' and not equals to the given %.2f ",
					updatedAmount-originalAmount, tt.args.ae.LoadAmount)
			}
		})
	}
}

func TestTransactionList_adjustAmount(t *testing.T) {
	type args struct {
		ae         AccountEntry
		dailyAmt   float32
		weeklyAmt  float32
		dailyCount int8
	}
	tests := []struct {
		name string
		tl   TransactionList
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Need to be adjusted back to original value",
			tl:   tl,
			args: args{
				ae: AccountEntry{
					ID:         "10003",
					CustomerID: "1234",
					LoadAmount: dailyMaxAmount + 1.00,
					LoadTime:   time.Now().AddDate(0, 0, -1),
				},
				dailyAmt:   dailyMaxAmount,
				dailyCount: dailyMaxCount,
				weeklyAmt:  weeklyMaxAmount,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalAmount := tt.tl[tt.args.ae.CustomerID].Transaction[tt.args.ae.LoadTime.Weekday()].Amount
			tt.tl.Update(tt.args.ae)
			tt.tl.adjustAmount(tt.args.ae, tt.args.dailyAmt, tt.args.weeklyAmt, tt.args.dailyCount)
			adjustedValue := tt.tl[tt.args.ae.CustomerID].Transaction[tt.args.ae.LoadTime.Weekday()].Amount
			if adjustedValue != originalAmount {
				t.Errorf("Load Amount adjustment error: 'adjusted value %.2f 'not equals to the orginal amount %.2f ",
					adjustedValue, originalAmount)
			}
		})
	}
}

func TestTransactionList_CreateOutput(t *testing.T) {
	type args struct {
		ae           AccountEntry
		dailyAmount  float32
		weeklyAmount float32
		dailyCount   int8
	}
	tests := []struct {
		name string
		tl   TransactionList
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Generate output based on set criteria : not accepted",
			tl:   tl,
			args: args{
				ae: AccountEntry{
					ID:         "10003",
					CustomerID: "1234",
					LoadAmount: dailyMaxAmount + 1.00,
					LoadTime:   time.Now().AddDate(0, 0, -1),
				},
				dailyAmount:  dailyMaxAmount,
				dailyCount:   dailyMaxCount,
				weeklyAmount: weeklyMaxAmount,
			},
			want: `{"id":"10003","customer_id":"1234","accepted":false}`,
		},
		{
			name: "Generate output based on set criteria : accepted",
			tl:   tl,
			args: args{
				ae: AccountEntry{
					ID:         "10004",
					CustomerID: "8888",
					LoadAmount: 1.00,
					LoadTime:   time.Now(),
				},
				dailyAmount:  dailyMaxAmount,
				dailyCount:   dailyMaxCount,
				weeklyAmount: weeklyMaxAmount,
			},
			want: `{"id":"10004","customer_id":"8888","accepted":true}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tl.Update(tt.args.ae)
			if got := tt.tl.CreateOutput(tt.args.ae, tt.args.dailyAmount, tt.args.weeklyAmount, tt.args.dailyCount); got != tt.want {
				t.Errorf("TransactionList.CreateOutput() = %v, want %v", got, tt.want)
			}
			tt.tl.adjustAmount(tt.args.ae, dailyMaxAmount, weeklyMaxAmount, dailyMaxCount)
		})
	}
}

func TestTransactionList_Reset(t *testing.T) {
	tests := []struct {
		name string
		tl   TransactionList
	}{
		// TODO: Add test cases.
		{
			name: "Reset amount and counts by deleting all items in the Transaction",
			tl:   tl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tl.Reset()
			for _, v := range tt.tl {
				if len(v.Transaction) != 0 {
					t.Error("Transaction list reset failed")
				}
			}
		})
	}
}
