package authentication

import (
	"appengine"
	"appengine/memcache"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"time"

	"github.com/Orbittman/timeoff/dto"
)

func CreateAuthCookie(w http.ResponseWriter, r *http.Request, loginRequest dto.LoginRequest) {
	u4, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	cookie, err := setCookie(w, r, "auth_token", u4.String(), false)
	createSession(r, cookie.Value, loginRequest)
}

func setCookie(w http.ResponseWriter, r *http.Request, name string, value string, encrypt bool) (*http.Cookie, error) {
	var hashKey = []byte("4d167c3f54fe46918412465c00ad81a7")
	var blockKey = []byte("3cec3deeecc74abba40333bc85507eca")
	var s = securecookie.New(hashKey, blockKey)

	if encrypt {
		cookieValue, err := s.Encode(name, value)
		if err != nil {
			return nil, err
		}

		value = cookieValue
	}

	cookie := &http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
	}

	http.SetCookie(w, cookie)

	return cookie, nil
}

func createSession(r *http.Request, key string, loginRequest dto.LoginRequest) {
	c := appengine.NewContext(r)

	var buf []byte
	b := bytes.NewBuffer(buf)
	enc := gob.NewEncoder(b)
	enc.Encode(Session{Name: loginRequest.UserName})

	item := &memcache.Item{
		Key:        key,
		Expiration: time.Minute * 30,
		Value:      b.Bytes(),
	}

	if err := memcache.Add(c, item); err == memcache.ErrNotStored {
		c.Infof("item with key %q already exists", item.Key)
	} else if err != nil {
		c.Errorf("error adding item: %v", err)
	}

	if err := memcache.Set(c, item); err != nil {
		c.Errorf("error setting item: %v", err)
	}
}

func GetSession(r *http.Request, key string) (Session, error) {
	c := appengine.NewContext(r)

	item, err := memcache.Get(c, key)
	if err == memcache.ErrCacheMiss {
		c.Infof("item not in the cache")
	} else if err != nil {
		c.Errorf("error getting item: %v", key)
	} else {
		c.Infof("name: %s", item.Value)
	}

	var s Session
	reader := bytes.NewReader(item.Value)
	dec := gob.NewDecoder(reader)
	err = dec.Decode(&s)
	return s, err
}

type Session struct {
	Name string
}
