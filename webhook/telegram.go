package webhook

type Telegram struct {
	Name string
}

func (t Telegram) HandleMessage(message string) (string, error) {

	return "This is Telegram!!", nil
}

func NewTelegram() Platform {
	return Telegram{}

}
