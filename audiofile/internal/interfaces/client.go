package interfaces

import "net/http"

// interfaces allow for polymorphism and decoupling of code.
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}