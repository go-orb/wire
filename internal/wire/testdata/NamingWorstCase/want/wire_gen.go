// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/go-orb/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	context2 "context"
)

// Injectors from wire.go:

func inject(context3 context2.Context, err2 struct{}) (context, error) {
	mainContext, err := provide(context3)
	if err != nil {
		return context{}, err
	}
	return mainContext, nil
}
