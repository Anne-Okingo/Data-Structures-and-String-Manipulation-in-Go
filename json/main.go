package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name     string
    Age      int
    IsStudent bool
    Courses  []string
}

// func main() {
//     jsonString := `{
//         "name": "John Doe",
//         "age": 30,
//         "isStudent": false,
//         "courses": ["math", "history", "chemistry"]
//     }`

//     var p Person
//     err := json.Unmarshal([]byte(jsonString), &p)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Printf("%+v\n", p)
// }

func main(){
	p := Person{
		Name: "John Doe",
		Age: 30,
		IsStudent:false,
		Courses: []string{"math","history","chemistry"},
	}
	jsonData,err := json.Marshal(p)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println(string(jsonData))
}