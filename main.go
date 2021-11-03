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
	Secret = "segredoMuitoSecretoDeTestes"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type Params struct {
	AutoCopy  bool
	HideEmail bool
	Subject   string
}

func Generate(params Params) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	if !params.HideEmail {
		claims["unique_name"] = params.Subject
	}
	claims["iss"] = Secret
	now := time.Now()
	claims["iat"] = float64(now.Unix())
	claims["exp"] = float64(now.AddDate(100, 0, 0).Unix())
	claims["sub"] = params.Subject
	claims["aud"] = Secret
	t, err := token.SignedString([]byte(Secret))
	must(err)
	return t

}
func main() {
	copy := flag.Bool("c", true, "Automatically copies the token to clipboard")
	std := flag.Bool("std", false, "Automatically copies the token to clipboard")
	email := flag.String("e", "", "Type (App or Email)")
	hideEmail := flag.Bool("h", false, "Hide unique_name")
	flag.Parse()

	token := Generate(Params{
		AutoCopy:  *copy,
		Subject:   *email,
		HideEmail: *hideEmail,
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
