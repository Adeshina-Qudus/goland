package main

type AutoPolicy struct {
	accountNumber int
	makeAndModel  string
	state         string
}

func (policy *AutoPolicy) setAccountNumber(accountNumber int) {
	policy.accountNumber = accountNumber
}
func (policy *AutoPolicy) getAccountNumber() int {
	return policy.accountNumber
}
func (policy *AutoPolicy) setMakeAndModel(makeAmdModel string) {
	policy.makeAndModel = makeAmdModel
}
func (policy *AutoPolicy) getMakeAndModel() string {
	return policy.makeAndModel
}
func (policy *AutoPolicy) setState(state string) {
	policy.state = state
}
func (policy *AutoPolicy) getState() string {
	return policy.state
}

//
//func (policy *AutoPolicy) isNoFaultState() bool{
//	return switch policy.getState() {
//	case "CT", "MA", "NH","NJ","NY","PA","VT":
//	default:
//	}
//	return false
//}
