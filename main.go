package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://m.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	printCityList(all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}

const text = "My email is sq0502@126.com"

func getEmailURL() {
	re := regexp.MustCompile(`.+@.+\..+`)
	match := re.FindString(text)
	fmt.Println(match)

}

func printCityList(contents []byte) {
	regStr := `<a href="(http://m.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	re := regexp.MustCompile(regStr)
	matchs := re.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		fmt.Printf("city:%s,URL:%s\n", m[2], m[1])
		fmt.Println()
	}
	fmt.Printf("Matches :%d", len(matchs))

}
