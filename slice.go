package main

import "fmt"

// a SLICE allows you to group together VALUES of the same TYPE

func main() {
	// x := type{values} // composite literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for k, v := range x {
		fmt.Println(k, v)
	}

	// Slicing a slice
	fmt.Println(x[2:4])

	// Append
	x = append(x, 77, 88, 99, 114, 1014)
	fmt.Println(x)

	y := []int{234, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	// Deleting
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// built-in function "make"
	z := make([]int, 10, 12)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	z[0] = 42
	z[9] = 999
	fmt.Println(z)
	//z[10] = 1000
	z = append(z, 10000)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z = append(z, 10000)
	z = append(z, 10000)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	// Multidimensional slice
	jb := []string{"a", "b", "c", "d"}
	mp := []string{"e", "f", "g", "h"}
	xp := [][]string{jb, mp}
	fmt.Println(xp)
	fmt.Println(len(xp))

}
