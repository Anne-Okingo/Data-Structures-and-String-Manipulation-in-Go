package main

import (
"fmt"
)

func main(){
	// declaring a map
	// var m map[string]int
	// creating a map
	// m := make(map[string]int)
	// initializing a map
	// m := map[string]int{
	// 	"Alice" : 30,
	// 	"Bob": 25,
	// }
	// the above for small known fixed data

	// m := make(map[string]int)

	// m["alice"]= 30
	// m["bob"]= 25
	// m["alice"]= 32

	// the above for large data,of initioal unknown values
	mapp := make(map[string]int)

	models := []string{"Apple","Mac","Ipad","Android"}
	prices :=[]int{100,70,40,10}

	for i:= range models{
		// dynamically adding key-value pair
		mapp[models[i]] = prices[i]
	}
	mapp["Apple"] = 120
	fmt.Println(mapp)
}
