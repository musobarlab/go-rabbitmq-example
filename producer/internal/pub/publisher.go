package pub

//Publisher interface
type Publisher interface {
	Publish(string, []byte) error
}