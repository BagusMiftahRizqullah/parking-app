package models

import (
	"errors"
	"fmt"
)

// storage sementara parking lot
type ParkingLot struct {
	Capacity int
	Slots    []*ParkingSlot
}

// storage sementara mobil
type ParkingSlot struct {
	Number int
	Car    *Car
}

func NewParkingLot(capacity int) *ParkingLot {
	// proses membuat parking lot dengan kapasitas sesuai input
	slots := make([]*ParkingSlot, capacity)
	for i := 0; i < capacity; i++ {
		slots[i] = &ParkingSlot{Number: i + 1, Car: nil}
	}
	return &ParkingLot{
		Capacity: capacity,
		Slots:    slots,
	}
}

func (p *ParkingLot) ParkCar(registrationNumber string) (int, error) {
	// proses parkir mobil sesuai urutan parking lot yang di pilih
	for _, slot := range p.Slots {
		if slot.Car == nil { // slot is available
			slot.Car = &Car{RegistrationNumber: registrationNumber}
			return slot.Number, nil
		}
	}
	return 0, errors.New("parking lot is full")
}

func (p *ParkingLot) LeaveCar(registrationNumber string) (int, error) {
	// proses menghapus mobil pada parking lot yang di pilih

	for _, slot := range p.Slots {
		if slot.Car != nil && slot.Car.RegistrationNumber == registrationNumber {
			slot.Car = nil
			return slot.Number, nil
		}
	}
	return 0, errors.New("car not found")
}

func (p *ParkingLot) PrintStatus() {
	// proses menampilkan status keseluruhan parking lot dan mobilnya
	fmt.Println("Slot No. Registration No.")
	for _, slot := range p.Slots {
		if slot.Car != nil {
			fmt.Printf("%d %s\n", slot.Number, slot.Car.RegistrationNumber)
		}
	}
}
