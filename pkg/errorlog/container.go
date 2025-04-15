package errorlog

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ErrorContainer struct {
	Errors []*ErrorLog
}

func NewErrorContainer() *ErrorContainer {
	return &ErrorContainer{
		Errors: []*ErrorLog{},
	}
}

// AddError adds a new error to the container
func (ec *ErrorContainer) AddError(errorLog *ErrorLog) {
	if errorLog != nil {
		ec.Errors = append(ec.Errors, errorLog)
	}
}

// HasError returns true if the container has any errors
func (ec *ErrorContainer) HasError() bool {
	return len(ec.Errors) > 0
}

// Json returns the errors in JSON format (for backward compatibility)
func (ec *ErrorContainer) Json() string {
	return ec.Format(&JSONFormatter{})
}

// Text returns the errors in text format (for backward compatibility)
func (ec *ErrorContainer) Text() string {
	return ec.Format(&TXTFormatter{})
}

// Format returns the errors in the specified format
func (ec *ErrorContainer) Format(formatter ErrorFormatter) string {
	return formatter.Format(ec)
}

// ErrorFormatter is the interface that defines how errors should be formatted
type ErrorFormatter interface {
	Format(container *ErrorContainer) string
}

// JSONFormatter implements ErrorFormatter for JSON format
type JSONFormatter struct{}

func (f *JSONFormatter) Format(container *ErrorContainer) string {
	jsonBytes, _ := json.Marshal(container)
	return string(jsonBytes)
}

// TXTFormatter implements ErrorFormatter for text format
type TXTFormatter struct{}

func (f *TXTFormatter) Format(container *ErrorContainer) string {
	var sb strings.Builder
	for i, err := range container.Errors {
		sb.WriteString(fmt.Sprintf("Error %d:\n", i+1))
		sb.WriteString(fmt.Sprintf("  Module: %s\n", err.Module))
		sb.WriteString(fmt.Sprintf("  Layer: %s\n", err.Layer))
		sb.WriteString(fmt.Sprintf("  Package: %s\n", err.Package))
		sb.WriteString(fmt.Sprintf("  Message: %s\n", err.Error))
	}
	return sb.String()
}

func (ec *ErrorContainer) Error() string {
	if !ec.HasError() {
		return ""
	}
	errMsg := "Errors: \n"
	for _, e := range ec.Errors {
		errMsg += e.Err().Error() + "\n"
	}
	return errMsg
}
