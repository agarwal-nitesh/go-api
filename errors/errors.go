package errors

import (
	"fmt"
	"log"
)

type (
	CustomError struct {
		ErrorMessage 	string 		`json:"errorMessage"`
		ErrorCode    	int  		`json:"errorCodes"`
	}
)

func (this *CustomError) Error() string {
	if this == nil {
		return fmt.Sprintf("Error is nil, but was tried to be accessed")
	}
	return fmt.Sprintf("%s - %s", this.ErrorMessage, this.ErrorCode)
}

func NewCustomError(code int, message string) error {
	return &CustomError{message, code}
}

func (this *CustomError) Log() {
	if this == nil {
		log.Println("Error is nil, but was tried to be accessed")
	}
	log.Println(fmt.Sprintf("Error: %s - %s", this.ErrorCode, this.ErrorMessage))
}


var (
	ResourceInitializationError = NewCustomError(21000, "Error in initializing the resource")
	DomainInitializationError   = NewCustomError(21001, "Error in initializing the domain")
	ResourceNotFound            = NewCustomError(21002, "Resource not found")
	NotificationNotFound        = NewCustomError(21003, "Notification not found")
	RequestMapperTypeCastError  = NewCustomError(21004, "RequestMapper type and casting type are different!")
	BadRequestError				= NewCustomError(400, "Bad Request")
)

