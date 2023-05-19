package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username        string   `json:"username"`
	FirstName       string   `json:"firstname"`
	LastName        string   `json:"lastname"`
	Age             int64    `json:"age"`
	Password        string   `json:"-"`
	EnrolledCourses []string `json:"enrolledCourses,omitempty"`
}

func main() {
	encodeJSON()
	decodeJSON()
}

func encodeJSON() {

	fmt.Println("Encoding the json data")
	usersData := []User{
		{"nikhilrana", "Nikhil", "Rana", 23, "asdf@12345", []string{"Python Bootcamp", "Rust Bootcamp"}},
		{"aman", "Aman", "Sharma", 21, "asdf@12345", []string{"Python Bootcamp", "Golang Bootcamp"}},
		{"veenarana", "Veena", "Rana", 19, "asdf@12345", nil},
	}

	// Marshal data gives data in bytes
	// Masrhsal Indent gives data in indented way as provided the indentation required.

	//dataBytes, err := json.Marshal(usersData)

	//var responseString strings.Builder
	dataBytes, err := json.MarshalIndent(usersData, "", "\t")
	if err != nil {
		panic(err)
	}

	//dataLength, err := responseString.Write(dataBytes)

	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s\n", dataBytes)
	// fmt.Println(dataLength)
	// fmt.Println(responseString.String())
	fmt.Printf("%s\n", dataBytes)

}

func decodeJSON() {
	fmt.Println("Decoding the json data")

	jsonDataFromWeb := []byte(`
		{
			"username": "aman",
			"firstname": "Aman",
			"lastname": "Sharma",
			"age": 21,
			"enrolledCourses": [
					"Python Bootcamp",
					"Golang Bootcamp"
			]
		}
	`)

	var usersData User

	isJSON_Valid := json.Valid(jsonDataFromWeb)

	if isJSON_Valid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &usersData)
		fmt.Printf("%#v\n", usersData)
	} else {
		fmt.Println("Json is not valid")
	}

	// sometimes data is not valid which results into store them in map

	//var jsonDataToMap map[string]interface{}
	jsonDataToMap := make(map[string]interface{})
	json.Unmarshal(jsonDataFromWeb, &jsonDataToMap)
	//fmt.Printf("%#v\n", jsonDataToMap)

	for key, value := range jsonDataToMap {
		fmt.Printf("Key is %v and value is %v and type of value is %T\n", key, value, value)
	}

}
