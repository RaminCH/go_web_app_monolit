package main

import "testing"

func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("failed run()")
	}
}

//go test -v
//go/src/Web_Part1/bookings/cmd/web(master)$ go test -cover
//go test -coverprofile=coverage.out && go  tool cover -html=coverage.out
