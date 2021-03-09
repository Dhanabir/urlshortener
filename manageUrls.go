package main

import (
	"math/rand"
	"io/ioutil"
	"encoding/json"
	"log"
)

func LoadFile(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(data), &urls)
	return nil
}

func SaveJsonUrl(urls map[string]shortUrl, file string) error {
	data, err := json.MarshalIndent(&urls, "", " ")
	if err != nil {
		log.Fatal(err)
		return err
	}
	ioutil.WriteFile(file, data, 0644)
	return nil
}

func SaveAndSendUrl(longUrl string) string {
	uniqueId := int(rand.Uint32())
	base := 62
	existingUrl, exists := urls[longUrl]
	if exists == false {
		newShortUrl := CreateShortUrl(base, base62, uniqueId)
		url := shortUrl{
			uniqueId: uniqueId,
			shortUrl: newShortUrl,
		}
		urls[longUrl] = url
		err := SaveJsonUrl(urls, "urls.txt")
		if err != nil {
			log.Fatal(err)
		}
		return newShortUrl
	} else {
		return existingUrl.shortUrl
	}
}

func CreateShortUrl(base int, base62 []rune, uniqueId int) string{
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