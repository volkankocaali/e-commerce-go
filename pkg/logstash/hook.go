package logstash

import (
	"net"
)

type LogstashHook struct {
	conn net.Conn
}

func NewLogstashHook(address string) (*LogstashHook, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &LogstashHook{conn: conn}, nil
}

// Fire sends the log entry to Logstash
func (hook *LogstashHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.conn.Write([]byte(line))
	return err
}

// Levels returns the log levels that this hook is interested in
func (hook *LogstashHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
