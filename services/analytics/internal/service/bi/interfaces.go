package bi

type DataProcessor interface {
	Handle(name EventType, userID string, distinct string, properties []byte) error
}
