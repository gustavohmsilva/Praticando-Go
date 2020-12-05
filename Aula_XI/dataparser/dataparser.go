package dataparser

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Product struct {
	Name     string
	Category string
	Desc     string
}

func Read(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func File(f *os.File) (Product, error) {
	if f == nil {
		return Product{}, fmt.Errorf("No file supplied")
	}
	defer f.Close()
	lines := bufio.NewScanner(f)
	var p Product
	row := 0
	for lines.Scan() {
		switch row {
		case 0:
			p.Name = lines.Text()
		case 1:
			p.Category = lines.Text()
		case 2:
			p.Desc = lines.Text()
		default:
			p.Desc = fmt.Sprintf("%s\n%s", p.Desc, lines.Text())
		}
		row++
	}
	return p, nil
}

func CSV(f *os.File, hasHeader bool) ([]Product, error) {
	if f == nil {
		return []Product{}, fmt.Errorf("No file supplied")
	}
	defer f.Close()
	rows := csv.NewReader(f)
	var ps []Product
	for {
		row, err := rows.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return ps, err
		}
		if hasHeader {
			hasHeader = false
			continue
		}
		ps = append(ps, Product{row[0], row[1], row[2]})
	}
	return ps, nil
}
