package webhook

type Platform interface {
	HandleMessage(message string) (string, error)
}

var Platforms = map[string]Platform{
	"facebook": Facebook{Name: "this is facebook"},
	"telegram": Telegram{Name: "this is telegram"},
}
