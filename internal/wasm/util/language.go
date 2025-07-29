package util

type (
	Language struct {
		code string
		name string
	}
)

func (l *Language) Code() string {
	return l.code
}

func (l *Language) Name() string {
	return l.name
}

var (
	english = &Language{
		code: "en",
		name: "English",
	}
	spanish = &Language{
		code: "es",
		name: "Español",
	}
	portuguese = &Language{
		code: "pt",
		name: "Português",
	}
)
