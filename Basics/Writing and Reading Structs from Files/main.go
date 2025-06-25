//You have seen how to marshal a struct to JSON, and the next listing shows you how to
//expand on this and write the contents of a struct and an embedded struct to a file.
//Note that the code would be quite similar when marshalling via HTTP or sockets

package main

import (
	"encoding/json"
	"io/ioutil"
)

type Salary struct {
	Basic float64
}
type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func main() {

	data := Employee{
		FirstName: "Ibrahim",
		LastName:  "El Othmani",
		Email:     "ibrahim@gmail.com",
		Age:       26,
		MonthlySalary: []Salary{{Basic: 15000.00}, {Basic: 16000.00},
			{Basic: 17000.00}},
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("my_salary.json", file, 0644)
}
