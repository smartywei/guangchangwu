package functions

import (
	"fmt"
	"net/http"
)

func OutPutPageNotFound(w http.ResponseWriter) {
	fmt.Fprintln(w, "<html><center><h1>404 Page Not Found</h1></center></html>")
}
