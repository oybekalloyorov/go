package main 

import "fmt"

func main(){
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//int as map key type
	phonebook := map[int]string{
		21351515: "mari",
		54546546: "luigi",
		21564845: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[21351515])
	
	phonebook[54546546] = "bowser"
	fmt.Println(phonebook)

	name := "Oybek"

	copyName := &name
	fmt.Println(*copyName)
}