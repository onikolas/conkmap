package conkmap

import "errors"
import "sync"

// map from anything to anything, protected by a mutex, fancy stuff
type ConkMap struct {
	lock sync.RWMutex
	data map[interface{}]interface{}
}

// Make a new ConkMap 
func New() ConkMap {

	c := ConkMap{}
	c.data = make(map[interface{}]interface{})
	return c
}

// Get data at key
func (c *ConkMap) Get(key interface{}) interface{} {

	c.lock.RLock()
	ret:=c.data[key]
	c.lock.RUnlock()
	return ret
}


// Set key-data pair
func (c *ConkMap) Set(key interface{}, data interface{}) error {

	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = data
	_, ok := c.data[key]
	
	
	if !ok {
		return errors.New("Could not insert element")
	} else {
		return nil
	}
}

// Is the map initiated
func (c *ConkMap) Initiated() bool {
	if c.data == nil {
		return true 
	} else {
		return false
	}
}

