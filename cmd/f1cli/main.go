package main

import (
	"fmt"
	//	"os"
	"flag"
)

type flagsConfig struct {
	season int
	data   string
	top    int
}

func main() {
	//fmt.Println("hello world ")
	var cnf flagsConfig
	path := "data/jolpica"
	flag.IntVar(&cnf.season, "season", 2022, "input season as 4 digit year")
	flag.StringVar(&cnf.data, "data", path, "folder with the jolpica CSVs ")
	flag.IntVar(&cnf.top, "top", 5, "Top n drivers to show ")
	flag.Parse()
	fmt.Println("The season is ", cnf.season)
	fmt.Println("The data is stored in  ", cnf.data)
	fmt.Println("The top driver you are looking for", cnf.top)
}
