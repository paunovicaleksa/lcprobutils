package main

import (
	"flag"
	"log"
	"os"
)

func main() {

	addPtr := flag.Bool("add", false, "add a problem to the list/file")
	listPtr := flag.Bool("list", false, "list all problems, sorted by urgency")
	removePtr := flag.Int("remove", -1, "remove a problem with given index, as seen in list")
	filePtr := flag.String("file", "problems.json", "path to file, default: problems.json")
	/* toggle interactive mode? */
	interactivePtr := flag.Bool("interactive", false, "toggle interactive mode")

	flag.Parse()

	file, err := os.OpenFile(*filePtr, os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		log.Fatal("Error opening file")
	}

	defer file.Close()
	defer Write(file)
	/* cli parsing here? */
	Read(file)

	/* only one at the same time? */
	if *addPtr {
		ParseAdd()
	} else if *listPtr {
		PrintList()
	} else if isFlagPassed("remove") {
		Remove(*removePtr)
	} else if *interactivePtr {
		/* TODO: tui */
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}
