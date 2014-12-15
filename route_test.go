package errmux

import (
	"fmt"
	"net/http"
)

func static(msg string) Handler {
	return func(http.ResponseWriter, *http.Request, int) {
		fmt.Println(msg)
	}
}

func ExampleRouter_specific() {
	//NB. static creates a handler that prints the string given to it when called.
	r := &Router{
		Specific: Routes{
			418: static("418 specific handler"),
			502: static("502 specific handler"),
		},
		Universal: static("universal handler"),
	}

	r.Error(nil, nil, 418)
	r.Error(nil, nil, 502)
	r.Error(nil, nil, 500)
	//Output:
	//418 specific handler
	//502 specific handler
	//universal handler
}

func ExampleRouter_range() {
	//NB. static creates a handler that prints the string given to it when called.
	r := &Router{
		Client: static("client"),
		Server: static("server"),
	}

	r.Error(nil, nil, 404)
	r.Error(nil, nil, 502)
	r.Error(nil, nil, 418)
	r.Error(nil, nil, 500)
	//Output:
	//client
	//server
	//client
	//server
}

func ExampleRouter_universal() {
	r := &Router{
		Universal: func(w http.ResponseWriter, r *http.Request, code int) {
			fmt.Println(DefaultStatusCodes.Message(code))
		},
	}

	r.Error(nil, nil, 404)
	r.Error(nil, nil, 418)
	r.Error(nil, nil, 502)
	//Output:
	//404 Not Found
	//418 I'm a teapot
	//502 Bad Gateway
}

func ExampleRouter_parent() {
	//NB. static creates a handler that prints the string given to it when called.
	parent := &Router{
		Universal: static("parent"),
	}

	child := &Router{
		Specific: Routes{
			404: static("overridden not found"),
		},
		Parent: parent,
	}

	child.Error(nil, nil, 404)
	child.Error(nil, nil, 403)
	//Output:
	//overridden not found
	//parent
}
