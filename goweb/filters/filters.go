package filters

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterSession() gin.HandlerFunc {
	store := sessions.NewCookieStore([]byte("secret"))
	// store, _ := sessions.NewRedisStore(
	// 	10,
	// 	"tcp",
	// 	config.GetEnv().REDIS_IP+":"+config.GetEnv().REDIS_PORT,
	// 	config.GetEnv().REDIS_PASSWORD,
	// 	[]byte("secret"))
	return sessions.Sessions("mysession", store)
}

func RegisterCache() gin.HandlerFunc {
	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewInMemoryStore(time.Minute)
	return cache.Cache(&cacheStore)
}
