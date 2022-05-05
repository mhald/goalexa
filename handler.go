package goalexa

import "github.com/aivahealth/goalexa/alexaapi"

type Handler interface {
	CanHandle(*alexaapi.RequestRoot) bool
	Handle(*alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error)
}
