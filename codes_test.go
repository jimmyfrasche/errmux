package errmux

import (
	"fmt"
	"testing"
)

func TestStatusCodesZeroValue(t *testing.T) {
	var s StatusCodes
	c, m := s.Message(404)
	if c != 500 || m != "Internal Server Error" {
		t.Errorf(`Expected 500 "Internal Server Error" but got %d "%s"`, c, m)
	}
}

func ExampleStatusCodes() {
	codes := StatusCodes{
		404: "It's not here",
		500: "It broke :(",
	}

	fmt.Println(codes.Message(404))
	//There's no message for 418, so Message will return 500
	fmt.Println(codes.Message(418))
	//Output:
	//404 It's not here
	//500 It broke :(
}
