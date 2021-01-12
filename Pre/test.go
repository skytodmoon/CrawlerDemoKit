package Pre

import (
	"fmt"
	"regexp"
)

const text = "My email is sq0502@126.com"

func GetEmailUrl() {
	re := regexp.MustCompile("sq0502@126.com")
	match := re.FindString(text)
	fmt.Println(match)

}
