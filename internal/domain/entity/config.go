package entity

// Config enables dynamic updates to configuration by defining read functions for each
// configuration property
type Config struct {
	Token string
	Url   string
}

func (receiver Config) ReadToken() string {
	return receiver.Token
}

func (receiver Config) ReadUrl() string {
	return receiver.Url
}
