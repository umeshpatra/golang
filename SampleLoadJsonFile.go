package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Catlog : TODO
type Catlog struct {
	Product_id string `json:"product_id"`
	Quantity   int32  `json:"quantity"`
}

// CatlogList : TODO
type CatlogList struct {
	Catlogs []Catlog `json:"catlogs"`
}

func ReadJSonFile(data CatlogList) {
	file, err := ioutil.ReadFile("catalog.json")

	if err != nil {
		fmt.Println(err)
	}

	//data := CatlogList{}

	json.Unmarshal([]byte(file), &data)
}

func main() {
	//FindTheBiggestNumber()
	data := CatlogList{}
	ReadJSonFile(data)

	for i := 0; i < len(data.Catlogs); i++ {
		fmt.Println("Product Id :", data.Catlogs[i].Product_id)
		fmt.Println("Product Quantity :", data.Catlogs[i].Quantity)
	}

}
