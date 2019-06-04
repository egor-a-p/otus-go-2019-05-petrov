package shortener

import (
	"log"
	net "net/url"
	"strconv"
	"strings"
	"sync"
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type shortenerImpl struct {
	lock   sync.RWMutex
	urlMap map[int]string
}

func (s *shortenerImpl) Shorten(url string) string {
	u, err := net.Parse(url)
	if err != nil {
		log.Fatal(err)
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	i := len(s.urlMap)
	s.urlMap[i] = url

	return strings.Replace(url, u.RequestURI(), "/"+strconv.Itoa(i), -1)
}

func (s *shortenerImpl) Resolve(url string) string {
	u, err := net.Parse(url)
	if err != nil {
		log.Fatal(err)
	}

	i, err := strconv.Atoi(strings.ReplaceAll(u.RequestURI(), "/", ""))

	if err != nil {
		return ""
	}

	s.lock.RLock()
	defer s.lock.RUnlock()

	result, ok := s.urlMap[i]

	if ok {
		return result
	} else {
		return ""
	}
}

var instance Shortener
var once sync.Once

func GetInstance() Shortener {
	once.Do(func() {
		instance = &shortenerImpl{urlMap: make(map[int]string)}
	})

	return instance
}
