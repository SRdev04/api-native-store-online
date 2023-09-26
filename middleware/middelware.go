package middleware

import (
	"fmt"
	"net/http"
	"store-online-restfulapi/execption"
	"store-online-restfulapi/helper"
	"store-online-restfulapi/restcallapi"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}

}

func (auth *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// buat orang dalem
	//saya pakai apikey agar mudah tesnya
	//jika sebenarnya saya mencoba menggunakan oauth2 yang saya coba pahami alurnya
	test := r.Header.Get("X-ApiKey-X")
	urlRequets := r.URL.Path

	fmt.Println(test)
	if test == "x" {
		auth.Handler.ServeHTTP(w, r)

	} else if urlRequets == "/api/user/register" {
		auth.Handler.ServeHTTP(w, r)

	} else if urlRequets == "/api/user" {
		auth.Handler.ServeHTTP(w, r)
	} else {
		//disini akan restcall api untuk mencocokan apakah token tersebut ada/valid
		//dan ditangkap lewat headernya
		token := r.Header.Get("AUTHORIZATION")
		url := "http://localhost:9090/gettoken"

		found, err := restcallapi.RestCallApiChechk(url, token)

		if err != nil {
			panic(execption.NewNotFound(err.Error()))
		}

		if found {

			auth.Handler.ServeHTTP(w, r)

		} else {
			helper.Unauthorize(w)

		}

	}

}
