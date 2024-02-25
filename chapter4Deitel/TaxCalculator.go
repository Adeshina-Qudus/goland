package main

type TaxCalculator struct {
	name     string
	earnings float64
}

func (tax *TaxCalculator) setName(name string) {
	tax.name = name
}

func (tax *TaxCalculator) getName() string {
	return tax.name
}

func (tax *TaxCalculator) setEarnings(earning float64) {
	if earning < 0 {
		tax.earnings = 0
	} else {
		tax.earnings = earning
	}
}

func (tax *TaxCalculator) getEarning() float64 {
	return tax.earnings
}

func (tax *TaxCalculator) calculateTaxRate() float64 {
	var taxRate float64
	if tax.earnings <= 30_000 {
		taxRate = 0.15 * tax.getEarning()
	} else {
		taxRate = 0.20 * tax.getEarning()
	}
	return taxRate
}
