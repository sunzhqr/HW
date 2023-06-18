package in_memory_cash

type Cache struct {
	m map[string]interface{}
}

func New() *Cache {
	return &Cache{
		m: make(map[string]interface{}),
	}
}

func (c Cache) Get(k string) interface{} {
	for key, value := range c.m {
		if key == k {
			return value
		}
	}
	return nil
}

func (c *Cache) Set(k string, v interface{}) {
	for key, _ := range c.m {
		if key == k {
			c.m[key] = v
		}
	}
}

func (c *Cache) Delete(k string) {
	delete(c.m, k)
}
