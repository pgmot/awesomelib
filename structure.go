package awesomelib


type Hoge struct {
	Name string
}

func NewHoge(name string) Hoge {
	return Hoge{
		Name: name,
	}
}
