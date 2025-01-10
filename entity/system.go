package entity

type ClientContextKey struct{}

type ClientDevice struct {
	RemoteAddress string
	XForwardedFor string
	ClientIP      string
	UserAgent     string
}

type RequestIdContextKey struct{}
