package webhook

type Platform interface {
	HandleMessage(message string) (string, error)
}