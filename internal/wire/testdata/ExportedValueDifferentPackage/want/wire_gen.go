// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/go-orb/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"os"
)

// Injectors from wire.go:

func injectedFile() *os.File {
	file := _wireFileValue
	return file
}

var (
	_wireFileValue = os.Stdout
)
