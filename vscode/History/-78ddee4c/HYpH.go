package urlshort

import "net/http"
func MapHandler(m map[string]string,mux http.ServeMux)http.HandlerFunc{
	
}


func YAMLHandler(b []byte, fallback http.Handler)http.HandlerFunc{
	return nil
}

