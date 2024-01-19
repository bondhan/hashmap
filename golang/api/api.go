package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func init() {

}

func Put(h http.ResponseWriter, r *http.Request) {
	_, err := h.Write([]byte("Put"))
	if err != nil {
		logrus.Error(err)
		return
	}
}

func Get(h http.ResponseWriter, r *http.Request) {
	_, err := h.Write([]byte("Get"))
	if err != nil {
		logrus.Error(err)
		return
	}
}

func Remove(h http.ResponseWriter, r *http.Request) {
	_, err := h.Write([]byte("Remove"))
	if err != nil {
		logrus.Error(err)
		return
	}
}
