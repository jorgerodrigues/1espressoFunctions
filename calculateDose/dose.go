package dose

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"io"
	"net/http"
)

func init() {
	functions.HTTP("Dose", Dose)
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func Dose(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello, World!")

}
