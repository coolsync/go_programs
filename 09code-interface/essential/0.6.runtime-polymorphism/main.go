// calculator different countries drop tax income

package main

import "fmt"

type taxSystem interface {
	calculatorTax() int
}

type usTax struct {
	taxPercentage int
	income        int
}

func (u usTax) calculatorTax() int {
	tax := u.income * u.taxPercentage / 100
	return tax
}

type chinaTax struct {
	taxPercentage int
	income        int
}

func (c chinaTax) calculatorTax() int {
	tax := c.income * c.taxPercentage / 100
	return tax
}

func main() {
	us := &usTax{
		taxPercentage: 20,
		income:        3000,
	}

	china := &chinaTax{
		taxPercentage: 30,
		income:        2000,
	}

	taxSystems := []taxSystem{us, china}

	totalTax := calculatorTotalTax(taxSystems)

	fmt.Printf("total tax is: $%d\n", totalTax)
}

func calculatorTotalTax(taxSystems []taxSystem) int {
	totalTax := 0

	for _, t := range taxSystems {
		totalTax += t.calculatorTax()
	}

	return totalTax
}
