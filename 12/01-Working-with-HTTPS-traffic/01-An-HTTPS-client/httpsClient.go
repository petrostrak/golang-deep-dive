package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// httpsClient.go might fail in some cases, depending on the server certificate
// The solution to this problem is the use of the InsecureSkipVerify: true option in the
// initialization of http.Transport.
func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL := os.Args[1]

	// The http.Transport structure is configured for TLS using
	// TLSClientConfig which holds another structure named tls.Config
	// that at this point uses its default values.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}

	client := &http.Client{Transport: tr}
	response, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	s := strings.TrimSpace(string(content))

	fmt.Println(s)

}
