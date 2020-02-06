package provider

// Driver ...
type Driver interface {
	Send(mobile, content string) (string, error)
	String() string
}
