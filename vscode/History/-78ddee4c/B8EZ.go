package URLSHORTNER

import "net/http"
func MapHandler(m map[string]string,mux http.ServeMux)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		path:=r.URL.Path
		if dest,ok:= m[path];ok{
			
		}
	}
}


func YAMLHandler(b []byte, fallback http.Handler)http.HandlerFunc{
	return nil
}

