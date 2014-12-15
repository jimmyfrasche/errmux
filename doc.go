//Package errmux is a router for http errors.
//
//In its simplest case it can replace
//	http.Error(w, 500, "Internal Server Error")
//with
//	errmux.Error(w, r, 500)
//where the "Internal Server Error" string comes from DefaultStatusCodes.
//
//Like http.Error, errmux.Error must be called
//before writing to the http.ResposneWriter.
//
//Unlike http.Error, handlers can be set for specific error codes,
//ranges of error codes, and the default handler.
//Additionally, many custom error routers can be created.
package errmux
