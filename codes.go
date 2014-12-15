package errmux

//StatusCodes maps status codes to messages.
type StatusCodes map[int]string

//Message returns the message for code.
//
//This should be preferred to s[code] as it handles cases such as no such code,
//by returning a 500 and the 500 message.
//If no 500 message is set, "Internal Server Error" will be used.
func (s StatusCodes) Message(code int) (int, string) {
	//if we're nil, no idea how to process code, so 500
	if s == nil {
		return 500, "Internal Server Error"
	}

	msg, ok := s[code]
	if !ok {
		//if there is no message, consider that an internal server error
		code = 500
		msg, ok = s[code]
		//if the 500 text was deleted for some reason, fall back
		if !ok {
			msg = "Internal Server Error"
		}
	}
	return code, msg
}

//DefaultStatusCodes is all standard http codes of the 4xx and 5xx series.
//
//This is used by DefaultHandler and UseTemplate.
//
//It is not safe to modify this after the server is started,
//but can be changed at any time before that.
var DefaultStatusCodes = StatusCodes{
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Request Entity Too Large",
	414: "Request-URI Too Long",
	415: "Unsupported Media Type",
	416: "Requested Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot", //it is of critical importance that this be handled properly
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	510: "Not Extended",
}
