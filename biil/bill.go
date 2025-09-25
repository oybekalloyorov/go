package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	//add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs
}

func (b *bill) updateTip(tip float64){
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64){
	b.items[name] = price
}

// save bill
func (b *bill) save(){
	data := []byte(b.format())

	// filename := strings.ReplaceAll(b.name, " ", "_") + ".txt"
	// err := os.WriteFile(filename, data, 0644)
	err := os.WriteFile("bill/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)		
	}
	fmt.Println("bill was saved to file")
}
