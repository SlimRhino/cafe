package responds

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
)

//GzipJSON todo
func GzipJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(status)
	gz := gzip.NewWriter(w)
	json.NewEncoder(gz).Encode(data)
	gz.Close()
}

//JSON response in common format
func JSON(w http.ResponseWriter, status int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Del("Content-Length")
	w.WriteHeader(status)
	w.Write(response)
}
