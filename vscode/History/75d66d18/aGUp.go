package urlHandler

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(m map[string]string, mux *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := m[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		mux.ServeHTTP(w, r)
	}
}

func YAMLHandler(b []byte, mapHandler *http.Handler) (http.HandlerFunc, error) {
	mm:=make(map[string]string)
	err:=yaml.Unmarshal(b, &mm)
	if err!=nil {
		return nil,err
	}
	return mapHandler(mm,),nil

}
