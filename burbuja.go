package main

import(
  "fmt"
)

func Burbuja(s []int64) []int64 {
	var temp int64 = 0

	for i := 0; i < len(s) - 1; i++{
		for j := 0; j < len(s) - 1; j++{
			if(s[j] > s[j+1]){
				temp = s[j]
				s[j] = s[j+1]
				s[j+1] = temp 
			}
		}
	}

	return s
}

func main()  {
	slice := []int64{3, 16, 0, 9, 2, 24, 1}
	slice2 := []int64{0, 1, 2, 3, 9, 16, 24} 
	new_slice := Burbuja(slice)

	for i, _ := range slice2{
		if (new_slice[i] != slice2[i]){
			fmt.Println("fallaste")
			break
		}
	}
}