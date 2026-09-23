package booking

import "errors"

var (
	ErrSeatAlreadyBooked = errors.New("Assento já está ocupado")
)

type Booking struct {
	ID         string
	MovieID    string
	SeatNumber string
	UserID     string
	Status     string
}

type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}
