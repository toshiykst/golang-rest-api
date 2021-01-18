package controllers

// Context is interface that contains controller context methods.
type Context interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{}) error
}
