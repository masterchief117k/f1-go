package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type flagsConfig struct {
	season int
	data   string
	top    int
}

func checkTop(t int) error {
	if t > 0 && t <= 22 {
		return nil
	}
	return errors.New("top is out of range")
}
func checkSeason(s int) error {
	if s >= 1950 && s <= 2025 {
		return nil
	}
	return errors.New("season is out of bound")
}

func checkData(d string) error {
	file, err := os.Open(d)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func main() {
	var cnf flagsConfig
	path := filepath.Join("data", "jolpica")
	defaultfile := "season.csv"
	flag.IntVar(&cnf.season, "season", 2022, "input season as 4 digit year")
	flag.StringVar(&cnf.data, "data", defaultfile, "folder with the jolpica CSVs ")
	flag.IntVar(&cnf.top, "top", 5, "Top n drivers to show ")
	flag.Parse()
	data := filepath.Join(path, cnf.data)
	err := errors.Join(checkTop(cnf.top), checkSeason(cnf.season), checkData(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The season is ", cnf.season)
	fmt.Println("The data is stored in  ", cnf.data)
	fmt.Println("The top driver you are looking for", cnf.top)
}
