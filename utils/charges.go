package utils

func CalculateCharge(hours int) int {

	// kondisi dua jam pertama 10$
	if hours <= 2 {
		return 10
	}

	// dua jam setelahnya di tambahkan 10$
	return 10 + (hours-2)*10
}
