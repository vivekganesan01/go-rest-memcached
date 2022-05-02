package util

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Client struct {
	client *memcache.Client
}

func init() {
	log.Println("mc init().")
}

func NewMemCached() (*Client, error) {
	client := memcache.New(os.Getenv("MEMCACHED"))
	if err := client.Ping(); err != nil {
		return nil, err
	}
	client.Timeout = 100 * time.Millisecond
	client.MaxIdleConns = 100

	return &Client{
		client: client,
	}, nil
}

func (c *Client) Get(val string) (Name, error) {
	i, e := c.client.Get(val)
	if e != nil {
		return Name{}, e
	}
	var res Name
	if err := gob.NewDecoder(bytes.NewReader(i.Value)).Decode(&res); err != nil {
		return Name{}, err
	}
	return res, nil
}

func (c *Client) Set(val Name) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(val); err != nil {
		return err
	}

	return c.client.Set(&memcache.Item{
		Key:        val.NConst,
		Value:      b.Bytes(),
		Expiration: int32(time.Now().Add(25 * time.Second).Unix()),
	})
}
