package transport

type Transport interface {
	MultiPublish(body [][]byte) error
}
