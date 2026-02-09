package main_test

import (
	"testing"

	go_specs_greet "github.com/jdgomez/go-specs-greet"
	"github.com/jdgomez/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	driver := go_specs_greet.Driver{BaseUrl: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
