package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"log"
)

var (
	longUrls = map[string]int{}
	shortUrls = map[int]string{}
	base62 = []rune{
	'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z',
	'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z',
	}
)

func CreateShortUrl(base int, base62 []rune, uniqueId int) string {
	q := uniqueId
	var r int
	var shortUrl []rune
	for q > 0 {
		r = q % base
		shortUrl = append(shortUrl, base62[r])
		q = q / base
	}
	if uniqueId == 0 {
		return string(base62[0])
	}
	return string(shortUrl)
}

func SaveAndSendUrl(url string) string {
	uniqueId := int(rand.Uint32())
	base := 62
	UrlId, exists := longUrls[url]
	if exists == false {
		shortUrl := CreateShortUrl(base, base62, uniqueId)
		shortUrls[uniqueId] = shortUrl
		longUrls[url] = uniqueId
		return shortUrl
	} else {
		return shortUrls[UrlId]
	}
}

func GetURL(w http.ResponseWriter, r *http.Request) {
	longUrl := r.URL.Path[1:]
	if len(longUrl) == 0 {
		return
	}
	shortUrl := SaveAndSendUrl(longUrl)
	fmt.Fprint(w, shortUrl)
}

func main() {
	http.HandleFunc("/", GetURL)
	log.Fatal(http.ListenAndServe(":8080", nil))
}