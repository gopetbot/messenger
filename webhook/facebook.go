package webhook

type Facebook struct {
	Name string
}

func (f Facebook) HandleMessage(message string) (string, error) {

	return "handled for facebook", nil
}
func NewFacebook() Platform {
	return Facebook{}
}
