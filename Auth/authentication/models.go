package authentication

//Token represents one token
type Token struct {
	Token        string
	ExpiresIn    int64
}

type User struct {
	Name     string
	Password string
}

