package auth

import (
	"assets/internal/api"
	"assets/internal/entity"
	serviceErrors "assets/internal/errors"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
)

type APIAuth interface {
	Auth(http.ResponseWriter, *http.Request)
}

type API struct {
	service ServiceAuth
}

func NewAPI(service ServiceAuth) APIAuth {
	return &API{service: service}
}

func (us *API) Auth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var authData entity.AuthData
		err := json.NewDecoder(r.Body).Decode(&authData)
		if err != nil {
			api.JSONResponseError(w, err)
			return
		}

		ip := getIP(r)

		var token entity.Token
		token.Token, err = us.service.Auth(r.Context(), &authData, ip)
		if err != nil {
			api.JSONResponseError(w, err)
			return
		}

		data, err := json.Marshal(&token)
		if err != nil {
			api.JSONResponseError(w, err)
			return
		}

		api.JSONResponse(w, data)
	default:
		api.JSONResponseError(w, serviceErrors.ErrorMethodNotAllowed)
		return
	}
}

func getIP(r *http.Request) net.IP {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")
	if len(splitIps) > 0 && splitIps[len(splitIps)-1] != "" {
		ipS := splitIps[len(splitIps)-1]
		if ipS == "::1" {
			ipS = "127.0.0.1"
		}
		return net.ParseIP(ipS)
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("SplitHostPort error: %s\n", err.Error())
		return nil
	}

	netIP := net.ParseIP(ip)
	if netIP != nil && netIP.String() == "::1" {
		ip = "127.0.0.1"
	}

	return net.ParseIP(ip)

}
