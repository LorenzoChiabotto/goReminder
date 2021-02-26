package main

import (
	"./authentication"
)

//RequestToken represents a token request
type RequestToken struct {
	authentication.User
}
//ResponseToken represents a token response
type ResponseToken struct {
	authentication.Token
}
