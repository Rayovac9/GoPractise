// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/webmin7761/go-school/homework/final/internal/conf"
	"github.com/webmin7761/go-school/homework/final/internal/server/travel"
	"github.com/webmin7761/go-school/homework/final/internal/service"
	"go.opentelemetry.io/otel/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(server *conf.Server, traceTracerProvider trace.TracerProvider, logger log.Logger) (*kratos.App, func(), error) {
	travelService := service.NewTravelService(logger)
	httpServer := travel.NewHTTPServer(server, traceTracerProvider, travelService)
	grpcServer := travel.NewGRPCServer(server, traceTracerProvider, travelService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
	}, nil
}
