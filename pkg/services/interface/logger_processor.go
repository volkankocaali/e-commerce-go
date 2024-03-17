package _interface

import "time"

type LogMessage struct {
	Request         string    `json:"request" bson:"request"`
	RequestHeaders  string    `json:"requestHeaders" bson:"requestHeaders"`
	Response        string    `json:"response" bson:"response"`
	ResponseHeaders string    `json:"responseHeaders" bson:"responseHeaders"`
	CreatedAt       time.Time `json:"createdAt" bson:"createdAt"`
}

type LoggerProcessor interface {
	Process(message LogMessage) error
}
