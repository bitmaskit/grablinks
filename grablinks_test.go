package grablinks

import (
	"fmt"
	"strings"
	"testing"
)

const testHTML = `
<html lang="en">
<head>
	<title>Title</title>
</head>
<body><div><p>
	<a href="http://www.example.org/#home">First url</a>
	<a href='http://www.example.net'>Second Url</a>
	<a style=\"\" href=www.example.com>Third Url</a></p></div>
	<a href="http://www.example.org/#main">First url</a>
	<a href="http://www.example.org/">First url</a>
	<a href='http://www.example.net'>Second Url</a>
	<a style=\"\" href=www.example.com>Third Url</a></p></div>
</body>
</html>
`

var expectedURLs = [3]string{"http://www.example.org/", "http://www.example.net", "www.example.com"}

func TestAll(t *testing.T) {
	reader := strings.NewReader(testHTML)

	links := All(reader)

	fmt.Println(links)

	if len(links) != 3 {
		t.Errorf("Wrong number of links: Got %d expected %d\n", len(links), 3)
	}

	for i, link := range expectedURLs {
		if links[i] != link {
			t.Errorf("Wrong link at index: %d. Expected %s, got %s", i, link, links[i])
		}
	}
}
