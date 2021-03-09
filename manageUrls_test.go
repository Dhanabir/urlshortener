package main 

import (
	"testing"
)

func TestCreateShortUrl(t *testing.T) {
	base62 := []rune{
		'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z',
		'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z',
	}

	tests := []struct {
		base int
		base62 []rune
		id int
		expected string
	}{
		{62,base62,0,"0"},
		{62,base62,14,"e"},
		{62,base62,62,"01"},
		{62,base62,7876554324,"MAd3B8"},
		{62,base62,562736544222098,"Yz578iNz2"},
	}

	for _, test := range tests {
		if output := CreateShortUrl(test.base, test.base62, test.id); output != test.expected {
			t.Error("String expected: ", output, test.expected)
		}
	}
}

func TestSaveAndSendUrl(t *testing.T) {
	tests := []struct {
		url string
		expected string
	}{
		{"https://www.google.com",""},
		{"jhttps://www.facebook.com",""},
	}

	for _, test := range tests {
		output := SaveAndSendUrl(test.url)
		if SaveAndSendUrl(test.url) != output {
			t.Error("String expected: ", output)
		}
	}
}

func TestLoadFile(t *testing.T) {
	tests := []string {
		"urls.txt",
		"abc.txt",
	}

	for _, test := range tests {
		if output:=LoadFile(test); output != nil {
			t.Error("nil expected")
		}	
	}
}

func TestSaveJsonUrl(t *testing.T) {
	urls := map[string]shortUrl {
		"https://www.google.com" : shortUrl{
			uniqueId: 908867566,
			shortUrl: "gfT5resD",
		},
		"https://www.facebook.com" : shortUrl{
			uniqueId: 1679898989,
			shortUrl: "sall4yTgja",
		},
	}

	tests := []string {
		"urls.txt",
		"abc.txt",
	}

	for _, test := range tests {
		if output:=SaveJsonUrl(urls, test); output != nil {
			t.Error("nil expected")
		}	
	}
}