package cache

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func Cache(key string, d ...time.Duration) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		v, ok := db.InstanceGet(pluginName)
		if !ok {
			return db
		}

		plugin := v.(*Caches)
		if len(d) > 0 {
			plugin.tmp.Dur = d[0]
		}
		plugin.tmp.Key = key
		fmt.Println(plugin)

		return db
	}
}
