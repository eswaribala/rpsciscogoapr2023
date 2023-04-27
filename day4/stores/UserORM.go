package stores

import (
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"net/http"
)

// GetUsers godoc
// @Summary Get details of all users
// @Description Get details of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Create a Resty Client
	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		Get("https://jsonplaceholder.typicode.com/users")
	fmt.Fprintf(w, resp.String())
}
