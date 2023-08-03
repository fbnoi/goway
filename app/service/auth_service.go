package service

import (
	"flynoob/goway/internal"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) bool {
	return true
}
func AttachUserInfo(client *internal.Client) {
	token := client.GetToken()

}
