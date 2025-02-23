// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/go-orb/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectBaz() (Baz, func(), error) {
	foo, cleanup := provideFoo()
	bar, cleanup2, err := provideBar(foo)
	if err != nil {
		cleanup()
		return 0, nil, err
	}
	baz, err := provideBaz(bar)
	if err != nil {
		cleanup2()
		cleanup()
		return 0, nil, err
	}
	return baz, func() {
		cleanup2()
		cleanup()
	}, nil
}
