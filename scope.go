package cache

import (
	"time"

	"gorm.io/gorm"
)

// db.Where(maps).Scopes(cache.Cache("xxx", 10)).....
// 缓存的一个scope。默认是不需要缓存的
func Cache(key string, d ...time.Duration) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		v, ok := db.Plugins[pluginName]

		if !ok {
			return db
		}
		plugin := v.(*Caches)
		if len(d) > 0 {
			plugin.tmp.Dur = d[0]
		}
		plugin.tmp.Key = key
		return db
	}
}
