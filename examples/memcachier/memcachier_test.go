package memcachier_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/rbo13/memcacher/examples/memcachier"
)

const key = "myKey"

var c *memcachier.Memcachier

type UserContainer struct {
	Name     string
	Age      int
	Address  string
	Birthday time.Time
}

func init() {
	c = memcachier.NewMemcachier(memcachier.Config{Server: "localhost:11211", Username: "", Password: ""})
}

func TestSetMemcachier(t *testing.T) {

	user := struct {
		Name     string
		Age      int
		Address  string
		Birthday time.Time
	}{
		"Richard",
		12,
		"Cebu City",
		time.Now(),
	}

	out, err := json.Marshal(user)

	if err != nil {
		t.Errorf("Error Marshalling: %v", err)
		t.Fail()
	}

	ok, err := c.Set(key, string(out))

	if err != nil && ok == false {
		t.Errorf("Error: %v", err)
		t.Fail()
	}

}

func TestGetMemcachier(t *testing.T) {
	var userContainer UserContainer
	val, err := c.Get(key)

	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fail()
	}

	err = json.Unmarshal([]byte(val.(string)), &userContainer)

	if err != nil {
		t.Errorf("Error Unmarshalling: %v", err)
		t.Fail()
	}

	t.Log(userContainer)
}

func TestDeleteMemcachier(t *testing.T) {
	ok, err := c.Delete(key)

	if err != nil {
		t.Errorf("Error due to: %v", err)
	}

	if !ok {
		t.Error("Failed to Delete item from cache")
	}

}
