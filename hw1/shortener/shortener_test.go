package shortener

import (
	"testing"
)

func TestShortener(t *testing.T) {
	url1 := "https://example.org/path?foo=bar"
	url2 := "http://otus.ru/some-long-link"

	shortenUrl1 := GetInstance().Shorten(url1)
	shortenUrl2 := GetInstance().Shorten(url2)

	if resolveUrl1 := GetInstance().Resolve(shortenUrl1); resolveUrl1 != url1 {
		t.Errorf("Fail: %s -> %s -> %s", url1, shortenUrl1, resolveUrl1)
	}

	if resolveUrl2 := GetInstance().Resolve(shortenUrl2); resolveUrl2 != url2 {
		t.Errorf("Fail: %s -> %s -> %s", url2, shortenUrl2, resolveUrl2)
	}

}
