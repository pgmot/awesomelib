package awesomelib


type Hoge struct {
	Username string
}

func NewHoge(username string) Hoge {
	return Hoge{
		Username: username,
	}
}
