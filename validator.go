package main

type validate func(token *Token) bool

func UserValidator(userId int64) validate {
	return func(token *Token) bool {
		return userId == token.User
	}
}
