package URLSHORTNER

import "net/http"
func MapHandler(m map[string]string,mux http.ServeMux)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}


func YAMLHandler(b []byte, fallback http.Handler)http.HandlerFunc{
	return nil
}

