package datastructures

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// ServeHTTP function
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

// constructor function
// uses a convention regarding the naming 'newUserController'
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`), // backtick can be used to define a raw string literal (equally to @ in c#)
	}
}
