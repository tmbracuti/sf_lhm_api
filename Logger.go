package main

import (
	//"log"
	"net/http"
	"time"
)

//the Logger function wraps the provided http.Handler in
// http.HandlerFunc instance (an adaptor) so that we can
// "decorate" the original handler call with logging statments
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			inner.ServeHTTP(w, r)

			logger.Printf("%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)

		})

}
