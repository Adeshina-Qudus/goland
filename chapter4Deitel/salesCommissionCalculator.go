package main

type SalesCommission struct {
	priceForItemSold float64
}

func getSalary() float64 {
	return 200
}

func getPercentage() float64 {
	return 0.09
}

func (sales *SalesCommission) setPriceForItemSold(priceForItemSold float64) {
	if priceForItemSold < 0 {
		sales.priceForItemSold = 0
	} else {
		sales.priceForItemSold = priceForItemSold
	}
}

func (sales *SalesCommission) getPriceForItemSold() float64 {
	return sales.priceForItemSold
}

func calculateCommission() float64 {
	var sale *SalesCommission
	return getPercentage() * sale.getPriceForItemSold()

}

func calculateTotalSalary() float64 {
	return calculateCommission() + getSalary()
}
