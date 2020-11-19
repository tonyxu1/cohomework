package model

import (
	"encoding/json"
	"io/ioutil"
)

// Config : Get configuration information
type Config struct {
	DailyMaxAmout   float32 `json:"dailymaxamount"`
	DailyMaxCount   int8    `json:"dailymaxcount"`
	WeeklyMaxAmount float32 `json:"weeklymaxamount"`
	OutputFile      string  `json:"outputfile"`
}

// GetValue : read values from config.json  to this instance
func (conf *Config) GetValue() error {
	c, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(c, &conf)
	if err != nil {
		return err
	}
	return nil
}
