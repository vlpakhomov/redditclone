package auth

import "context"

type sessionKey string

const (
	SessKey sessionKey = "session"
)

type Session struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type SessionID struct {
	AccessToken string
}

func ContextWithSession(ctx context.Context, session Session) context.Context {
	return context.WithValue(ctx, SessKey, session)
}
