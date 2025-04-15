package errorlog

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrLayerEmpty   = errors.New("layer should no be empty")
	ErrPackageEmpty = errors.New("package should no be empty")
	ErrMessageEmpty = errors.New("error should no be empty")
)

type ErrorLog struct {
	Module  string // module that originated the error i.e: "core"
	Layer   string // layer that originated the error i.e: "core"
	Package string // package that originated the error i.e: "domain"
	Error   string // error message
}

func NewErrorLog(module, layer, pkg, error string) *ErrorLog {
	errorLog := &ErrorLog{
		Module:  module,
		Layer:   layer,
		Package: pkg,
		Error:   error,
	}
	if err := errorLog.Validate(); err != nil {
		return nil
	}
	return errorLog
}

func (e *ErrorLog) Validate() error {
	if strings.TrimSpace(e.Layer) == "" {
		return ErrLayerEmpty
	}
	if strings.TrimSpace(e.Package) == "" {
		return ErrPackageEmpty
	}
	if strings.TrimSpace(e.Error) == "" {
		return ErrMessageEmpty
	}
	return nil
}

// converts the ErrorLog to a JSON string
func (e *ErrorLog) Json() string {
	jsonBytes, _ := json.Marshal(e)
	return string(jsonBytes)
}

// converts the ErrorLog to a formatted text string
func (e *ErrorLog) Text() string {
	return fmt.Sprintf("Module: %s, Layer: %s, Package: %s, Error: %s", e.Module, e.Layer, e.Package, e.Error)
}

// converts the ErrorLog to a Go error
func (e *ErrorLog) Err() error {
	err := fmt.Sprintf("Module: %s, Layer: %s, Package: %s, Error: %s", e.Module, e.Layer, e.Package, e.Error)
	return errors.New(err)
}
