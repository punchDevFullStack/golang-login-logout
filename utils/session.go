package utils

import (
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)

var globalSession *session.Manager

func init() {
	globalSession, _ = session.NewManager("redis", &session.ManagerConfig{
		CookieName:      "sessions",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		CookieLifeTime:  3600,
		ProviderConfig:  "127.0.0.1:6379",
	})

	go globalSession.GC()

}

func GetSession() *session.Manager {
	return globalSession
}
