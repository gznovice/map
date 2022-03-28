package main

import(
	"fmt"
)

func  print(my *map[string]int){
	fmt.Println(my)
	fmt.Println(*my)
	fmt.Println("len:", len(*my))
}

func main(){
	//make 
	var map1 = make(map[string]int)
	map1["apple"] = 2
	map1["banana"] = 4
	print(&map1)
	//literal
	map2 := map[string]int{
		"tom":2,
		"ben":1,
		"hit":3}
	print(&map2)
	//iterate
	for k, va := range(map1){
		map1[k] = va + 1
	}
	//exist
	if  num, exists := map1["apple"]; exists{
		fmt.Println(num, " apple(s) exist(s)")
	}else{
		fmt.Println("no apple exists")
	}
	//delete
	delete(map1, "apple")
	
	
	//exist
	if  _, exists := map1["apple"]; !exists{		
		fmt.Println("no apple exists")
	}
}