// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"LiteraTest/show/v1/internal"
	"LiteraTest/show/v1/internal/uc"
)

// Injectors from wire.go:

// Initialize method to initialize wire
func Initialize() (*internal.Handler, error) {
	findDoubleBookedEventsUC := uc.NewFindDoubleBookedEventsUC()
	parseEventsToUTCUC := uc.NewParseEventsToUTCUC()
	handler := internal.NewHandler(findDoubleBookedEventsUC, parseEventsToUTCUC)
	return handler, nil
}
