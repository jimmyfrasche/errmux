package errmux

import "net/http"

//Routes maps specific status codes to a specific Handler.
type Routes map[int]Handler

//A Router finds the appropriate Handler for a given error code.
//
//The zero value of a Router will always use the DefaultHandler.
//
//Resolution is handled by the For method and follows the following steps:
//
//First, if Specific is set, it is checked for the specific error code.
//
//Second, if the code is in the 4xx range, Client is returned, if set.
//If the code is in the 5xx range, Server is returned if set.
//If both are set, resolution ends here.
//
//Third, if the Universal handler is set, it's returned.
//
//Fourth, if Parent is set, the Parent router follows the above steps.
//
//Otherwise, the DefaultHandler is used.
type Router struct {
	//Specific contains Handlers for specific codes.
	Specific Routes
	//Client is the handler for all 4xx codes.
	Client Handler
	//Server is the handler for all 5xx codes.
	Server Handler
	//Universal is the handler for all error codes that do not match any of the above.
	Universal Handler
	//Parent allows lookup to continue in another Router.
	//This is useful to override some error handling in specific circumstances.
	Parent *Router
}

//For returns the Handler for a given status code.
//
//If code is not in the 4xx or 5xx range, the code is changed to 500.
func (r *Router) For(code int) (h Handler) {
	//not an error code is an error
	if code < 400 || code >= 600 {
		code = 500
	}

	//see if a specific handler is set
	if r.Specific != nil {
		h = r.Specific[code]
	}

	//if not, see if the general client or server handler is set, respectively
	if h == nil {
		if code >= 400 && code < 500 {
			h = r.Client
		} else {
			h = r.Server
		}
	}

	//if the appropriate general handler is not set, try the universal handler
	if h == nil {
		h = r.Universal
	}

	if h == nil {
		//if we have a parent, let the parent handle it
		if r.Parent != nil {
			return r.Parent.For(code)
		}
		//otherwise return the default handler
		h = DefaultHandler
	}

	return h
}

//Error gets the Handler for code and invokes it with w and req
//
//It is the caller's responsibility to ensure
//that nothing has been written to w before calling Error.
func (r *Router) Error(w http.ResponseWriter, req *http.Request, code int) {
	r.For(code)(w, req, code)
}
