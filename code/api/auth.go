/* vim: foldmethod=marker foldmarker={{{,}}} foldlevel=2
*/
package main

import(
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"time"
	"errors"
	"net/http"
	utils "github.com/ondrax/sympinator-be/code/utils"
)

type TokenClaims struct {
	UserName string
	jwt.StandardClaims
}

var signKey []byte = []byte("")

// {{{ Authentication router middleware -- requires sending of token on each access request
func JWTAuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		exceptions := []string{"/","/login"}
		path := r.URL.Path

		for _,ex := range exceptions { // don't check exceptions for auth tokens -- we want to be able to log a user in or create one (maybe)
			if ex == path {
				log.Println("loading page with authentication exception")
				next.ServeHTTP(w,r)
				return
			}
		}

		tokenString := r.Header.Get("AuthToken")

		if tokenString == "" {
			response := AuthResponse{false,"no token provided"}
			utils.WriteJSONResponse(w,response,500) // TODO: look up non-arbitrary server error numbers
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			_,isCorrectMethod := token.Method.(*jwt.SigningMethodHMAC)
			if (isCorrectMethod != true) {
				return nil, errors.New("improperly signed or completely malformed token")
			}
			return signKey, nil
		})

		if (err != nil) {

			response := AuthResponse{false,"authentication token malformed"}
				utils.WriteJSONResponse(w,response,500)
			// TODO: error responses inside routing middleware should probably be handled
			// using http.Error ? -- but we need this for the api, so we want to return JSON
			// maybe there is no need for AuthResponse struct and we can use default http.Error() and check for error code
		}

		if token.Valid {
			response := AuthResponse{true,"authentication token valid"}
			utils.WriteJSONResponse(w,response,200)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				response := AuthResponse{false,"authentication token malformed"}
				utils.WriteJSONResponse(w,response,500)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				response := AuthResponse{false,"authentication token not in use anymore"}
				utils.WriteJSONResponse(w,response,500)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				response := AuthResponse{false,"authentication token not in use yet"}
				utils.WriteJSONResponse(w,response,500)
			} else {
                response := AuthResponse{false,"uhnhandled error while handling token"+err.Error()}
				utils.WriteJSONResponse(w,response,500)
			}
		} else {
			response := AuthResponse{false,"unexpected error while handling token"+err.Error()}
				utils.WriteJSONResponse(w,response,500)
		}
		// check if authentication token is present

		log.Println("auth BEFORE")
		next.ServeHTTP(w,r)
		log.Println("auth AFTER")
	})
}
// }}}


// {{{
func makeTokenForUser(uname string, expiryHours int) (string,error) {
	claims := TokenClaims {
		uname,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiryHours) * time.Hour).Unix(),
			Issuer:    "com.sympinator.app",
		} }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	signedstring, err := token.SignedString(signKey)

	return signedstring, err
}

type AuthResponse struct {
	Success bool
	Message string
}
// }}}
