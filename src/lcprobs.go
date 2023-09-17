package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
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
	// Solved     bool      `json:"solved"` /* solved problems instantly have LOW priority, but should remain until deleted? */
}

var lcproblems []Lcproblem

func write(file *os.File) error {
	outstr, err := json.Marshal(lcproblems)
	if err != nil {
		return err
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.WriteAt(outstr, 0)

	return err
}

func read(file *os.File) error {
	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &lcproblems)

	if err != nil {
		return err
	}

	sort.Slice(lcproblems, func(i, j int) bool {
		return (lcproblems[i].Added.Nanosecond() * (lcproblems[i].Difficulty + 1)) > (lcproblems[j].Added.Nanosecond() * (lcproblems[j].Difficulty + 1))
	})

	return nil
}

func remove(index int) {
	lcproblems[index] = lcproblems[len(lcproblems)-1]
	lcproblems = lcproblems[:len(lcproblems)-1]
}

func parseAdd() {
	var problem Lcproblem
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("problem name: ")
	reader.Scan()
	problem.Name = reader.Text()
	easy := color.GreenString("Easy")
	medium := color.YellowString("Medium")
	hard := color.RedString("Hard")
	fmt.Printf("problem difficulty (%v, %v, %v): ", easy, medium, hard)
out:
	for {
		reader.Scan()
		switch strings.ToLower(reader.Text()) {
		case "easy":
			problem.Difficulty = Easy
			break out
		case "medium":
			problem.Difficulty = Medium
			break out
		case "hard":
			problem.Difficulty = Hard
			break out
		default:
			fmt.Println("invalid problem difficulty (Easy, Medium, Hard)")
			fmt.Printf("problem difficulty (%v, %v, %v): ", easy, medium, hard)
			continue
		}
	}

	problem.Added = time.Now()

	lcproblems = append(lcproblems, problem)
}

func printList() {
	fmt.Println("Problem Name            Added On      Difficulty")
	for _, problem := range lcproblems {
		var name string
		if len(problem.Name) > 15 {
			name = problem.Name[:15]
			name = name + "[...]"
		} else {
			name = problem.Name
		}

		var difficulty string
		switch problem.Difficulty {
		case Easy:
			difficulty = color.GreenString("Easy")
		case Medium:
			difficulty = color.YellowString("Medium")
		case Hard:
			difficulty = color.RedString("Hard")
		default:
			difficulty = ""
		}

		dateFormatted := problem.Added.Format("Jan 02, 2006")

		fmt.Printf("%-20.20v   %v   %v\n", name, dateFormatted, difficulty)
	}
}
