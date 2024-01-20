package api

import (
	"github.com/bondhan/hashmap/hashmap"
	"github.com/sirupsen/logrus"
	"net/http"
)

var hm *hashmap.HashMap

func init() {
	hm = new(hashmap.HashMap)
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
