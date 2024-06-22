package bi

type DataProcessor interface {
	Handle(name EventType, properties []byte) error
}
