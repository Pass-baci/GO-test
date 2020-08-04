package session

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemorySeesionMgr struct {
	sessionMap map[string]Session
	rwlock sync.RWMutex
}

func NewMemorySeesionMgr() *MemorySeesionMgr {
	m := &MemorySeesionMgr{
		sessionMap: make(map[string]Session,1024),
	}
	return m
}

func (m *MemorySeesionMgr)Init(addr string, options ...string) (err error) {
	return
}

func (m *MemorySeesionMgr)CreateSession() (session *MemorySession, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	sessionId := id.String()
	session = NewMemorySession(sessionId)
	return
}

func (m *MemorySeesionMgr)Get(sessionId string)(session Session,err error) {
	return
}