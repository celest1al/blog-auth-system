package routes

import (
	"net/http"

	"github.com/celest1al/blog-auth-system/utils"
)

func login(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	shouldReturn := utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
}
