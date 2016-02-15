package bruxism

import (
	"errors"
	"io"
)

type MessageType string

const (
	MessageTypeCreate MessageType = "create"
	MessageTypeUpdate             = "update"
	MessageTypeDelete             = "delete"
)

// Message is a message interface, wraps a single message from a service.
type Message interface {
	Channel() string
	UserName() string
	UserID() string
	UserAvatar() string
	Message() string
	RawMessage() string
	MessageID() string
	IsModerator() bool
	Type() MessageType
}

// ErrAlreadyJoined is an error dispatched on Join if the bot is already joined to the request.
var ErrAlreadyJoined = errors.New("Already joined.")

// Service is a service interface, wraps a single service such as YouTube or Discord.
type Service interface {
	Name() string
	UserName() string
	Open() (<-chan Message, error)
	IsMe(message Message) bool
	SendMessage(channel, message string) error
	DeleteMessage(channel, messageID string) error
	SendFile(channel, name string, r io.Reader) error
	BanUser(channel, userID string, duration int) error
	UnbanUser(channel, userID string) error
	SetPlaying(game string) error
	Join(join string) error
	Typing(channel string) error
	PrivateMessage(userID, messageID string) error
	IsPrivate(message Message) bool
	SupportsMultiline() bool
	CommandPrefix() string
	ChannelCount() int
	SupportsMessageHistory() bool
	MessageHistory(chanel string) []Message
}

// LoadFunc is the function signature for a load handler.
type LoadFunc func(*Bot, Service, []byte) error

// SaveFunc is the function signature for a save handler.
type SaveFunc func() ([]byte, error)

// HelpFunc is the function signature for a help handler.
type HelpFunc func(*Bot, Service, bool) []string

// MessageFunc is the function signature for a message handler.
type MessageFunc func(*Bot, Service, Message)

// Plugin is a plugin interface, supports loading and saving to a byte array and has help and message handlers.
type Plugin interface {
	Name() string
	Load(*Bot, Service, []byte) error
	Save() ([]byte, error)
	Help(*Bot, Service, bool) []string
	Message(*Bot, Service, Message)
}
