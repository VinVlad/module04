package internal

import "fmt"

func CalcPrice(cust Customer, price int) (int, error) {
	res, err := cust.CalcDiscount()
	if err != nil {
		return 0, fmt.Errorf("расчёт скидки: %w", err)
	}
	return price - res, nil
}
