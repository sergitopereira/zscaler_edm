package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

// csv records
var records [][]string
var add_row bool

func TerminalHelper(args []string) {
	if len(args) != 2 {
		fmt.Println("->Error: Program requires file path ")
		log.Fatal(">Error: Program requires file path")
	}
}

func csv_read(file string) *csv.Reader {
	//function to open file and read csv
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("->Error: Not possible to open file " + file)
		log.Fatal(err)
	}
	csvReader := csv.NewReader(f)

	return csvReader
}

func csv_write(records [][]string) {
	//function to write file to csv
	f, err := os.Create("updated_file.csv")

	if err != nil {

		log.Fatalln("->Error: Failed to open file name updated_file.csv ", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("->Error: writing record to file", err)
		}
	}
	log.Printf("End index script")

}

func main() {
	args := os.Args
	// Regex Any combination of letters (a-z, A-Z). Including hyphens and underscores
	reg, err := regexp.Compile("[^a-zA-Z0-9_-]+")
	if err != nil {
		log.Fatal(err)
	}
    // Matcghes any character that is not 0-9 space or pyhon
	reg_num, err := regexp.Compile("[^0-9 -]")
	if err != nil {
		log.Fatal(err)
	}

	log_file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(log_file)
	log.Printf("Start index script v1.0")
	defer log_file.Close()

	// open csv
	TerminalHelper(args)
	csvReader := csv_read(args[1])
    row :=0 
	for {
		rec, err := csvReader.Read()
		if row == 0{
			// headers
			records = append(records, rec)
			row +=1
			continue

		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("->Error: while reading the csv file. Review %#v\n", rec)
			log.Fatal(err)
		}
		// Check Column 1  is Alphanumeric
		e := reg.MatchString(rec[1])
		if e {
			log.Print("->Warning: Special Characters found in a alphanumeric field-> " + rec[1])
			rec[1] = reg.ReplaceAllString(rec[1], "")
			log.Print("->Warning: Special Characters replaced-> " + rec[1])
		}
        // Check Column 0 is numeric
        e1 := reg_num.MatchString(rec[0])
        if e1 {
			log.Print("->Warning: Special Characters found in a numeric fielf-> " + rec[0])
			rec[0] = reg_num.ReplaceAllString(rec[0], "")
			log.Print("->Warning: Special Characters replaced-> " + rec[0])

		}
		// check each row is at least 3 characters log. if row is higher than 24 discard characters
		add_row = true
		for i := 0; i < len(rec); i++ {
			if len(rec[i]) < 3 {
				log.Print("->Error: Field with less than 3 characters. Discarting row " + rec[i])
				add_row = false
			}
			if len(rec[i]) >24 {
				rec[i]=rec[i][0:23]
				log.Print("->Warning: Field with more than 24 characters.Removing extra characters " + rec[i])

			}

		}
		if add_row {
			records = append(records, rec)
		}

		//

	}
	csv_write(records)

}
