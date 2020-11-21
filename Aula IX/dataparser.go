package dataparser

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Product struct {
	Name     string
	Category string
	Desc     string
}

func badReadFile(filename string) (Product, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Product{}, err
	}
	rows := strings.Split(string(data), "\n")
	var p Product
	for i, row := range rows {
		switch i {
		case 0:
			p.Name = row
		case 1:
			p.Category = row
		case 2:
			p.Desc = row
		default:
			p.Desc = fmt.Sprintf("%s\n%s", p.Desc, row)
		}
	}
	return p, nil
}

func ReadFile(filename string) (Product, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Product{}, err
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
