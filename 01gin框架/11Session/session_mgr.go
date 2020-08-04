package session

type SessionMgr interface {
	Init(add string, options ...string) (err error)
	CreateSession() (session Session, err error)
	Get(sessionId string) (session Session, err error)
}
