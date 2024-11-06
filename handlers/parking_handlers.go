package handlers

import (
	"fmt"
	"parking-app/models"
	"parking-app/utils"
)

var parkingLot *models.ParkingLot

func CreateParkingLot(capacity int) {
	// menambahkan data pada  ParkingLot
	parkingLot = models.NewParkingLot(capacity)
	fmt.Printf("Created parking lot with %d slots\n", capacity)
}

func Park(carNumber string) {

	// validation jika parking lot tidak ada
	if parkingLot == nil {
		fmt.Println("Parking lot not created")
		return
	}

	// Fungsi untuk parkir mobil sesuai urutan parking Lot
	slot, err := parkingLot.ParkCar(carNumber)
	if err != nil {
		fmt.Println("Sorry, parking lot is full")
	} else {
		fmt.Printf("Allocated slot number: %d\n", slot)
	}
}

func Leave(carNumber string, hours int) {

	// validation jika parking lot tidak ada
	if parkingLot == nil {
		fmt.Println("Parking lot not created")
		return
	}

	// fungsi untuk menghapus mobil pada parking lot
	slot, err := parkingLot.LeaveCar(carNumber)
	if err != nil {
		fmt.Printf("Registration number %s not found\n", carNumber)
	} else {
		// menghitung charge parking untuk mobil yang keluar / jam
		charge := utils.CalculateCharge(hours)
		fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", carNumber, slot, charge)
	}
}

func Status() {
	if parkingLot == nil {
		fmt.Println("Parking lot not created")
		return
	}

	// Menampilkan status keseluruhan parking lot dan mobilnya
	parkingLot.PrintStatus()
}
