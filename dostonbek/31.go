package main

import "fmt"

func main(){
	// 31 Map yaratish
	m := map[string]int{"Alex": 20, "Bob": 25}
	fmt.Println("Boshlangich:",m)
	
	//32 Qiymatni olish va yangilash
	m["Alex"] = 21
	fmt.Println("Boshlangich:",m)
	fmt.Println("Alex yashi:", m["Alex"])

	//33 Kalit mavjudligini tekshirish
	val, ok := m["Kate"]
	if ok {
		fmt.Println("Topildi", val)
	}else{
		fmt.Println("Topilmadi")
	}

	// 34 Kalitni o‘chirish
	delete(m, "Alex")
	fmt.Println(m)

	// 35 Map bo‘ylab yurish (iteratsiya)

	newMap := map[string]int{"a": 1, "b": 2}

	for k, v := range newMap{
		fmt.Println(k ,"->",v)
	}
}