package transport

func NewNSQTransport() Transport {
	return &NSQPublisher{}
}

type NSQPublisher struct{}

func (p *NSQPublisher) MultiPublish(body [][]byte) error {
	return nil
}
