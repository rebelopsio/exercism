package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var a float32
	if balance < 0 {
		a = 3.213
		return a
	}
	if balance < 1000 {
		a = 0.5
		return a
	}
	if balance < 5000 {
		a = 1.621
		return a
	}
	a = 2.475
	return a
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	a := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		a++
	}
	return a
}
