package apperror

import "fmt"

type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
	Cause      error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code, message string, statusCode int) *AppError {
	return &AppError{Code: code, Message: message, StatusCode: statusCode}
}

func (e *AppError) WithMessagef(format string, args ...any) *AppError {
	return &AppError{
		Code:       e.Code,
		Message:    fmt.Sprintf(format, args...),
		StatusCode: e.StatusCode,
		Cause:      e.Cause,
	}
}

func Wrap(err error, code, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Cause:      err,
	}
}

var (
	// Auth and user
	ErrEmailTaken          = New("EMAIL_TAKEN", "Email already registered", 409)
	ErrInvalidCredentials  = New("INVALID_CREDENTIALS", "Invalid email or password", 401)
	ErrInvalidRefreshToken = New("INVALID_REFRESH_TOKEN", "Refresh token is invalid or expired", 401)
	ErrCannotDemoteAdmin   = New("CANNOT_DEMOTE_ADMIN", "Cannot change role of another admin", 400)

	// Common
	ErrNotFound    = New("NOT_FOUND", "Resource not found", 404)
	ErrForbidden   = New("FORBIDDEN", "Access denied", 403)
	ErrInvalidDate = New("INVALID_DATE", "Invalid date value", 400)

	// Booking and tickets
	ErrInsufficientTickets     = New("INSUFFICIENT_TICKETS", "Not enough tickets available", 400)
	ErrEventNotOnSale          = New("EVENT_NOT_ON_SALE", "This event is no longer on sale", 400)
	ErrExceedsMaxPerBooking    = New("EXCEEDS_MAX_PER_BOOKING", "Exceeds maximum tickets per booking", 400)
	ErrBookingAlreadyPaid      = New("BOOKING_ALREADY_PAID", "Booking has already been paid", 400)
	ErrBookingAlreadyCancelled = New("BOOKING_ALREADY_CANCELLED", "Booking is already cancelled", 400)
	ErrBookingExpired          = New("BOOKING_EXPIRED", "Booking has expired", 400)
	ErrPaymentAlreadyInitiated = New("PAYMENT_ALREADY_INITIATED", "Payment already initiated for this booking", 400)
	ErrEventHasPaidBookings    = New("EVENT_HAS_PAID_BOOKINGS", "Cannot delete event with paid bookings", 400)
	ErrQuantityBelowSold       = New("QUANTITY_BELOW_SOLD", "Quantity cannot be less than sold count", 400)
	ErrTicketTypeNameTaken     = New("TICKET_TYPE_NAME_TAKEN", "Ticket type name already exists for this event", 409)
	ErrTicketTypeHasBookings   = New("TICKET_TYPE_HAS_BOOKINGS", "Cannot delete ticket type with active bookings", 400)
)
