package interfaces

type Transformator interface {
	Transformate(body []byte) string
}