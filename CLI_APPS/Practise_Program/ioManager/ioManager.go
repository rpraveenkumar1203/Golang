package iomanager

type IOmanager interface {
	Readfile() ([]string, error)
	WriteFile(data interface{}) error
}
