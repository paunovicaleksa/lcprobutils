package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	Easy   int = 0
	Medium int = 1
	Hard   int = 2
)

type Lcproblem struct {
	Name       string    `json:"name"`
	Added      time.Time `json:"added"` /* calculate priority */
	Difficulty int       `json:"difficulty"`
	Solved     bool      `json:"solved"` /* solved problems instantly have LOW priority, but should remain until deleted? */
}

var lcproblems []Lcproblem

func write(file *os.File) error {
	outstr, err := json.Marshal(lcproblems)
	if err != nil {
		log.Fatal("error marshalling")
		return err
	}

	file.Truncate(0)
	file.WriteAt(outstr, 0)

	return nil
}

func read(file *os.File) error {
	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	json.Unmarshal(bytes, &lcproblems)

	return nil
}
