package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type flagsConfig struct {
	season int
	data   string
	top    int
}

func main() {
	var cnf flagsConfig
	var latestyear int
	const oldestyear int = 1950
	t := time.Now()
	latestyear = t.Year()
	path := "data/jolpica"
	flag.IntVar(&cnf.season, "season", 2022, "input season as 4 digit year")
	flag.StringVar(&cnf.data, "data", path, "folder with the jolpica CSVs ")
	flag.IntVar(&cnf.top, "top", 5, "Top n drivers to show ")
	flag.Parse()
	//season check

	if cnf.season < oldestyear {
		log.Fatalf("%d out of range", cnf.season)
	}
	if cnf.season > latestyear {
		log.Fatalf("%d out of range", cnf.season)
	}
	//top check
	if cnf.top < 1 {
		log.Fatalf("top values should be positive greater than zero")
	}
	if cnf.top > 22 {
		log.Fatalf("there are no more than 22 driver at a single time in a season ")
	}

	fmt.Println("The season is ", cnf.season)
	fmt.Println("The data is stored in  ", cnf.data)
	fmt.Println("The top driver you are looking for", cnf.top)
}
