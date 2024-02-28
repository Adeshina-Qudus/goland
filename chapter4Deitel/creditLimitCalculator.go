package main

type CreditLimitCalculator struct {
	balance            int
	totalItemCharged   int
	totalCreditApplied int
	newBalance         int
	allowedCreditLimit int
	accountNumber      int
}

func (limit *CreditLimitCalculator) setAccountNumber(accountNumber int) {
	limit.accountNumber = accountNumber
}

func (limit *CreditLimitCalculator) getAccountNumber() int {
	return limit.accountNumber
}
func (limit *CreditLimitCalculator) setBalance(balance int) {
	limit.balance = balance
}

func (limit *CreditLimitCalculator) getBalance() int {
	return limit.balance
}
func (limit *CreditLimitCalculator) setTotalItemCharged(totalItemCharged int) {
	limit.totalItemCharged = totalItemCharged
}
func (limit *CreditLimitCalculator) getTotalItemCharged() int {
	return limit.totalItemCharged
}
func (limit *CreditLimitCalculator) setTotalCreditApplied(totalCreditApplied int) {
	limit.totalCreditApplied = totalCreditApplied
}
func (limit *CreditLimitCalculator) getTotalCreditApplied() int {
	return limit.totalCreditApplied
}
func (limit *CreditLimitCalculator) setAllowedCreditLimit(allowedCreditLimit int) {
	limit.allowedCreditLimit = allowedCreditLimit
}
func (limit *CreditLimitCalculator) getAllowedCreditLimit() int {
	return 10_000
}
func (limit *CreditLimitCalculator) getNewBalance() int {
	limit.newBalance = limit.balance + limit.totalItemCharged - limit.totalCreditApplied
	return limit.newBalance
}
func (limit *CreditLimitCalculator) getCreditLimitExceeded() int {
	if limit.newBalance > limit.allowedCreditLimit {
		println("Credit limit exceeded")
	}
	return limit.allowedCreditLimit
}
