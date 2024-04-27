package main

func monthlyInstallmemt(price, interestRate, year float64) float64 {
	interest := price * (interestRate / 100) * year
	return (price + interest) / (year * 12)
}
