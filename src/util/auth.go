package auth

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

//Auth auth
func Auth(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Firebase SDK のセットアップ
		test := os.Getenv("FIREBASE_KEY")
		data, _ := base64.StdEncoding.DecodeString(test)
		credentials, err := google.CredentialsFromJSON(c, data)
		if err != nil {
			log.Printf("error credentials from json: %v\n", err)
			os.Exit(1)
		}
		opt := option.WithCredentials(credentials)
		app, err := firebase.NewApp(c, nil, opt)
		if err != nil {
			log.Printf("error initializing app: %v\n", err)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		// クライアントから送られてきた JWT 取得
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// JWT が無効なら Handler に進まず別処理
			fmt.Printf("error verifying ID token: %v\n", err)
			c.String(http.StatusUnauthorized, "error verifying ID token\n")
			return
		}
		log.Printf("Verified ID token: %v\n", token)
		next(c)
	}
}
