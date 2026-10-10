package store 

import "testing"

func TestSetThenGet(t *testing.T) {
	key := "xander"
	value := 5
    Set(key, value)
	get, ok := Get(key)

	if !ok || get != value {
        t.Errorf("Wanted %d from GET on %s, but got %d instead", value, key, get)
    }
}