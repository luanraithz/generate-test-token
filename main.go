package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/golang-jwt/jwt/v4"
)

type AppType string

const (
	Secret         = "segredoMuitoSecretoDeTestes"
	Email  AppType = "Email"
	App    AppType = "App"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type Params struct {
	AutoCopy bool
	Subject  string
	Type     string
}

func Generate(params Params) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	if params.Type == string(App) {
		claims["appid"] = params.Subject
	} else {
		claims["unique_name"] = params.Subject
	}
	claims["iss"] = Secret
	now := time.Now()
	claims["at"] = float64(now.Unix())
	expiration := float64(now.AddDate(100, 0, 0).Unix())
	claims["xp"] = expiration
	claims["exp"] = expiration
	claims["ud"] = Secret
	claims["sub"] = "1234567890"
	claims["aud"] = Secret
	t, err := token.SignedString([]byte(Secret))
	must(err)
	return t

}
func main() {
	copy := flag.Bool("c", true, "Automatically copies the token to clipboard")
	std := flag.Bool("std", false, "Automatically copies the token to clipboard")
	t := flag.String("t", string(Email), "Type (App or Email)")
	email := flag.String("e", "", "Type (App or Email)")
	flag.Parse()

	token := Generate(Params{
		AutoCopy: *copy,
		Subject:  *email,
		Type:     *t,
	})
	if *std {
		fmt.Print(token)
		return
	}
	if !*copy {
		print("Copy to clipboard?")
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
	}
	clipboard.WriteAll(token)
	print("Copied to clipboard")
}
