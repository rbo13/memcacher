package memcachier_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/whaangbuu/memcacher/examples/memcachier"
)

const key = "myKey"

type UserContainer struct {
	Name     string
	Age      int
	Address  string
	Birthday time.Time
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

	var userContainer UserContainer

	memcachier := memcachier.NewMemcachier("localhost:11211", "", "")

	out, err := json.Marshal(user)

	if err != nil {
		t.Errorf("Error Marshalling: %v", err)
		t.Fail()
	}

	ok, err := memcachier.Set(key, string(out))

	if err != nil && ok == false {
		t.Errorf("Error: %v", err)
		t.Fail()
	}

	val, err := memcachier.Get(key)

	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fail()
	}

	err = json.Unmarshal([]byte(val.(string)), &userContainer)

	if err != nil {
		t.Errorf("Error Unmarshalling: %v", err)
		t.Fail()
	}

	t.Log(userContainer.Name)

}
