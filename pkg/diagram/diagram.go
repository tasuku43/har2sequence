package diagram

import (
	"fmt"
	"strings"
)

type SequenceDiagram struct {
	Participants    []string
	Messages        []Message
	MessagePrefixes map[string]string
}

type Message struct {
	From     string
	To       string
	Request  Request
	Response Response
}

type Request struct {
	Method string
	Path   string
}

type Response struct {
	Status int
}

func (s SequenceDiagram) Render() string {
	var sb strings.Builder
	sb.WriteString("sequenceDiagram\n")
	for _, participant := range s.Participants {
		sb.WriteString(fmt.Sprintf("    participant %s\n", participant))
	}
	for _, message := range s.Messages {
		prefix, ok := s.MessagePrefixes[message.Request.Method]
		if !ok {
			prefix = ""
		}
		sb.WriteString(fmt.Sprintf("    %s->>%s: %s%s %s\n", message.From, message.To, prefix, message.Request.Method, message.Request.Path))
		sb.WriteString(fmt.Sprintf("    activate %s\n", message.To))
		sb.WriteString(fmt.Sprintf("    %s-->>%s: %d\n", message.To, message.From, message.Response.Status))
		sb.WriteString(fmt.Sprintf("    deactivate %s\n", message.To))
	}
	return sb.String()
}
