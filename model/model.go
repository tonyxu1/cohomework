package model

import (
	"encoding/json"
	"time"
)

// AccountEntry : AccountEntry model holds record from input.txt
type AccountEntry struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	LoadAmount float32   `json:"load_amount"`
	LoadTime   time.Time `json:"time"`
}

// OutputEntry : Json data that is for output.txt
type OutputEntry struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}

// WeeklyTransactionEntry : Holds total load amount of each weekday and transaction count of the customer and transaction id.
type WeeklyTransactionEntry struct {
	ID          string
	Transaction map[time.Weekday]DailyTotal
}

// DailyTotal : Holds dail load total amount and number of transactions.
type DailyTotal struct {
	Amount float32
	Count  int8
}

// TransactionList : Transaction information for all input customers
type TransactionList map[string]WeeklyTransactionEntry

// Init : Function to initalize the TransactionList
func (tl TransactionList) Init() {
	tl = make(map[string]WeeklyTransactionEntry)
}

// IsDupTransaction : Function to check if a load ID is observed more than once for a particular user, all but the first instance can be ignored
func (tl TransactionList) IsDupTransaction(transactionID, customerID string) bool {
	if v, ok := tl[customerID]; ok {
		if v.ID == transactionID {
			return true
		}
	}
	return false
}

/*- A maximum of $5,000 can be loaded per day
- A maximum of $20,000 can be loaded per week
- A maximum of 3 loads can be performed per day, regardless of amount
*/
func (tl TransactionList) canAccept(customerID string, dailyAmount, weeklyAmunt float32, dailyCount int8) bool {
	if v, ok := tl[customerID]; ok {
		t := v.Transaction
		var weeklyAmoutTotal float32 = 0.0
		for _, dt := range t {
			if dt.Amount > dailyAmount {
				return false
			}
			if dt.Count > dailyCount {
				return false
			}
			weeklyAmoutTotal = weeklyAmoutTotal + dt.Amount
		}
		if weeklyAmoutTotal > weeklyAmunt {
			return false
		}

	}
	return true

}

//Update : Function to update the Transaction List map
func (tl TransactionList) Update(ae AccountEntry) {

	if tl.IsDupTransaction(ae.ID, ae.CustomerID) {
		return
	}
	v, _ := tl[ae.CustomerID]
	newAmt := v.Transaction[ae.LoadTime.Weekday()].Amount + ae.LoadAmount
	newCount := v.Transaction[ae.LoadTime.Weekday()].Count + 1

	v.Transaction[ae.LoadTime.Weekday()] = DailyTotal{
		Amount: newAmt,
		Count:  newCount,
	}
}

// CreateOutput : Create output
func (tl TransactionList) CreateOutput(ae AccountEntry, dailyAmount, weeklyAmount float32, dailyCount int8) string {
	oe := OutputEntry{
		ID:         ae.ID,
		CustomerID: ae.CustomerID,
		Accepted:   tl.canAccept(ae.CustomerID, dailyAmount, weeklyAmount, dailyCount),
	}

	result, _ := json.Marshal(oe)
	return string(result)
}
