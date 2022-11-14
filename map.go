package main

import "fmt"

func main() {
	mp1 := map[string]int{
		"James":      32,
		"Miss Money": 27,
	}
	fmt.Println(mp1)
	fmt.Println(mp1["James"])
	fmt.Println(mp1["YoYo"])

	v, ok := mp1["YoYo"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := mp1["YoYo"]; ok {
		fmt.Println(v)
	}
	if v, ok := mp1["James"]; ok {
		fmt.Println(v)
	}

	mp1["Todd"] = 33
	for k, v := range mp1 {
		fmt.Println(k, v)
	}
	delete(mp1, "James")
	fmt.Println(mp1)

	delete(mp1, "YoYo")
	fmt.Println(mp1)

	if v, ok := mp1["Todd"]; ok {
		fmt.Println("value: ", v)
		delete(mp1, "Todd")
	}
	fmt.Println(mp1)
}
