package main

const (
	_Url        = "http://localhost:5555"
	_RequestSubject    = "videoProcessing"
	_ResponseSubject    = "processResult"
	_ChanSize   = 64
	_ResultPath = "../output"
)

type requestType string

const (
	processRequest requestType = "processRequest"
)

type Data struct {
	Type    requestType
	Payload []byte
}
