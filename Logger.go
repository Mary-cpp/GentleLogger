package main

import (
	"fmt"
	"time"
)

type level string

const (
	Info    level = "INFO"
	Debug   level = "DEBUG"
	Warning level = "WARNING"
	Error   level = "ERROR"
)

type message struct {
	timestamp time.Time
	level     level
	message   string
}

func (m message) format() string {
	return fmt.Sprintf("[ %s ] - %s - %s", m.timestamp.Format(time.DateTime), m.level, m.message)
}

func log(level level, msg string) {
	str := message{time.Now(), level, msg}
	fmt.Println(str.format())
}

func LogError(data string) {
	log(Error, data)
}

func LogInfo(data string) {
	log(Info, data)
}

func LogDebug(data string) {
	log(Debug, data)
}

func LogWarning(data string) {
	log(Warning, data)
}
