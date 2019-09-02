package routes

import (
	"net/http"

	"github.com/celest1al/blog-auth-system/utils"
)

func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	shouldReturn := utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
}
