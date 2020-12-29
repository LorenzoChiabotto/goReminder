package authentication

type Token struct {
	token         string
	refresh_token string
	expires_in    int
}

type User struct {
	Name     string
	Password string
}
