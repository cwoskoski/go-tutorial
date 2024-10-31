package main

import (
	"fmt"
	"regexp"
	s "strings"
)

type VersionMatch struct {
	Version string
	Matched bool
}

func MatchAcceptVersion(headerName string, headerValue string) VersionMatch {
	if headerName != "Accept" {
		return VersionMatch{}
	}

	r, _ := regexp.Compile("application/vnd\\.mycvtapi\\.([Vv]\\d+)")
	match := r.FindStringSubmatch(headerValue)
	if match != nil {
		fmt.Println("matched: " + match[1])
		return VersionMatch{match[1], true}
	}

	return VersionMatch{}
}

func main() {

	path := "/subscribers/ping"
	headers := map[string]string{
		"test":   "abc",
		"Accept": "application/vnd.mycvtapi.v1",
	}

	fmt.Printf("Incomming path: %s\n\n", path)

	var newPath string = path

	for k, v := range headers {
		fmt.Printf("checking: %s: %s\n", k, v)

		match := MatchAcceptVersion(k, v)
		if match.Matched {
			newPath = "/" + s.ToLower(match.Version) + newPath
		}
	}

	fmt.Printf("New path: %s\n\n", newPath)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
