package errmux

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

//Handler is http.HandlerFunc with an extra code field.
//
//It is the Handler's responsibility to call ResponseWriter.WriteHeader.
//If you do not need to set any headers, use Wrap or WrapFunc.
type Handler func(w http.ResponseWriter, r *http.Request, code int)

//Wrap takes an http.Handler, h, and returns a handler that writes the header
//with the given code and calls h.ServeHTTP.
func Wrap(h http.Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request, code int) {
		w.WriteHeader(code)
		h.ServeHTTP(w, r)
	}
}

//WrapFunc is Wrap for an http.HandlerFunc.
func WrapFunc(h http.HandlerFunc) Handler {
	return Wrap(h)
}

//UseTemplate is a helper to create a Handler that uses a template.
//
//The data parameter will always be
//	map[string]interface{}{
//		"Code": code,
//		"Message": message,
//	}
//where code and message are the (int, string) returned
//by StatusCodes.Message.
//
//The template rendering is buffered and if it returns an error,
//DefaultHandler will be called with 500
//and then the template's error message is appended to the output.
func UseTemplate(t interface {
	Execute(io.Writer, interface{}) error
}) Handler {
	return func(w http.ResponseWriter, r *http.Request, code int) {
		code, msg := DefaultStatusCodes.Message(code)

		//try to render t
		var buf bytes.Buffer
		err := t.Execute(&buf, map[string]interface{}{
			"Code":    code,
			"Message": msg,
		})

		//rendering t failed, use default handler then add error message
		if err != nil {
			//no status code for ironic failure, so go with 500
			DefaultHandler(w, r, 500)
			//would write this to http.Server.ErrorLog if we knew which server to use
			fmt.Fprintln(w, err)
			return
		}

		//otherwise, dump the buffer into w
		w.WriteHeader(code)
		io.Copy(w, &buf)
	}
}
