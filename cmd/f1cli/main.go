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
	flag.IntVar(&cnf.season, "season", 2022, "Driver season / year")
	flag.StringVar(&cnf.data, "data", path, "this is path to data set")
	flag.IntVar(&cnf.top, "top", 5, "check out the season leaders")
	flag.Parse()
	fmt.Println("The season is ", cnf.season)
	fmt.Println("The data is stored in  ", cnf.data)
	fmt.Println("The top driver you are looking for", cnf.top)
}
