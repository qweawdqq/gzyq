package onecache

import (
	"sync"

	"time"
)

/** 
 * @Description:
 * @author : 贾亮
 * @date : 2017-06-14 08:48
 */


//var ruleCache *CacheTable
//func init() {
//	ruleCache := Cache("RULE_CACHE")
//	fmt.Println("创建了缓存RULE")
//}

//func GetRuleCache()*CacheTable  {
//	return ruleCache
//}

var (
	CACHE_TIME time.Duration = 15 //缓存时间
	cache = make(map[string]*CacheTable)
	mutex sync.RWMutex
)

// Cache returns the existing cache table with given name or creates a new one
// if the table does not exist yet.
func Cache(table string) *CacheTable {
	mutex.RLock()
	t, ok := cache[table]
	mutex.RUnlock()

	if !ok {
		t = &CacheTable{
			name:  table,
			items: make(map[interface{}]*CacheItem),
		}

		mutex.Lock()
		cache[table] = t
		mutex.Unlock()
	}

	return t
}