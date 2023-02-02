package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeRate struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	var Divisa, opcion string
	var cantidad float64
	for {
		fmt.Println("indique las siglas de la divisa de la que quiere obtener el cambio, USD, GBP, EUR, CAD, EGP")
		fmt.Scanln(&Divisa)
		fmt.Println("indique la cantidad que tiene en esa divisa")
		fmt.Scanln(&cantidad)

		//response, err := http.Get("https://api.exchangerate-api.com/v4/latest/%v", DIVISA)
		url := fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/%s", Divisa)
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			var data ExchangeRate
			json.NewDecoder(response.Body).Decode(&data)

			fmt.Printf("El tipo de cambio a %v es:, ")
			fmt.Println(data.Rates["EUR"])
			cambio := data.Rates["EUR"]
			resultado := cambio * cantidad
			fmt.Printf("%v %v al cambio de %v son %v Euros\n", cantidad, Divisa, cambio, resultado)
		}
		fmt.Println("Quiere hacer más cambios de divisa?")
		fmt.Scanln(&opcion)
		if opcion == "NO" {
			break
		}
	}
}
