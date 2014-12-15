package errmux

import (
	"fmt"
	"net/http"
)

//DefaultRouter is the DefaultRouter used by Error.
//
//It is not thread safe to change this after the server has been started.
var DefaultRouter = &Router{}

//Error is shorthand for DefaultRouter.Error.
func Error(w http.ResponseWriter, r *http.Request, code int) {
	DefaultRouter.Error(w, r, code)
}

//DefaultHandler gets the message from DefaultStatusCodes
//and outputs it similarly to http.Error.
//
//There is no reason to call this directly.
//It is exposed only for documentation purposes.
func DefaultHandler(w http.ResponseWriter, r *http.Request, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	code, msg := DefaultStatusCodes.Message(code)
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}
