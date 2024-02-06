package main

import (
	"fmt"
	"log"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 300, true)

	//cust.WrOffDebt()

	fmt.Printf("%v\n\n", cust)

	disc, err := cust.CalcDiscount()
	fmt.Printf("Скидка: %+v\n", disc)

	price, err := internal.CalcPrice(cust, 1000)
	if err != nil {
		log.Fatalf("расчёт цены: %v", err)
		return
	}

	fmt.Println(price)

}
