package main

type validate func(token *Token) bool

func UserValidator(userId int64) validate {
	return func(token *Token) bool {
		return userId == token.User
	}
}

func UserScenarioValidator(userId int64, scenario string) validate {
	return func(token *Token) bool {
		return userId == token.User && scenario == token.Scenario
	}
}
