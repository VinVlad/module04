package main

import (
	"errors"
	"fmt"
	"log"
	"myapp/internal"
)

const DEFAULT_DISCOUNT = 500

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 0, false)

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	fmt.Printf("%+v\n", cust)

	price, err := internal.CalcPrice(*cust, 1000)
	if err != nil {
		log.Fatalf("расчёт цены: %v", err)
		return
	}

	fmt.Println(price)

}
