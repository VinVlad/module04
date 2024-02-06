package internal

import "errors"

const defaultDiscount = 500

type Customer struct {
	Name     string
	Age      int
	Balance  int
	Debt     int
	discount bool
}

func (c *Customer) WrOffDebt() error {
	if c.Debt >= c.Balance {
		return errors.New("Нет доступных средств для списания")
	}

	c.Balance -= c.Debt
	c.Debt = 0

	return nil
}

func (c Customer) CalcDiscount() (int, error) {
	if !c.discount {
		return 0, errors.New("discount not available")
	}
	result := defaultDiscount - c.Debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		discount: discount,
	}
}
