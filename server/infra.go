package server

const timeFormat = "2006-01-02"

type Push interface {
	Send(string) error
}
