package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bondhan/hashmap/hashmap"
	"github.com/sirupsen/logrus"
)

var hm *hashmap.HashMap

func init() {
	hm = new(hashmap.HashMap)
}

func Root(h http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		logrus.Error(err)
		return
	}

	_, err = h.Write([]byte(fmt.Sprint("hi, from ", hostname)))
	if err != nil {
		logrus.Error(err)
		return
	}
}

func Init(h http.ResponseWriter, r *http.Request) {
	_, err := h.Write([]byte("init"))
	if err != nil {
		logrus.Error(err)
		return
	}
}

func Put(h http.ResponseWriter, r *http.Request) {
	_, err := h.Write([]byte("put"))
	if err != nil {
		logrus.Error(err)
		return
	}
}

func Get(h http.ResponseWriter, r *http.Request) {
	_, err := h.Write([]byte("get"))
	if err != nil {
		logrus.Error(err)
		return
	}
}

func Remove(h http.ResponseWriter, r *http.Request) {
	_, err := h.Write([]byte("remove"))
	if err != nil {
		logrus.Error(err)
		return
	}
}
