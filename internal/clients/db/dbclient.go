package db

type ClientDB struct {
	Token string
}

type TokenGetter interface {
	TokenDB() string
}

func NewDBClient(tokenGetter TokenGetter) (*ClientDB, error) {
	return &ClientDB{
		Token: tokenGetter.TokenDB(),
	}, nil
}
