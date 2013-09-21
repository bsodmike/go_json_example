/*
 * Michael de Silva <michael@mwdesilva.com>
 * CEO @ http://omakaselabs.com/ / https://github.com/bsodmike
 */

package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func getContent(url string) ([]byte, error) {
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }

	// Send the request via a client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Defer the closing of the body
	defer resp.Body.Close()

	// Read the content into a byte array
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// At this point we're done - simply return the bytes
	return body, nil
}

func main() {
  result, _ := getContent("http://google.com")
	fmt.Printf("Output\n")
  fmt.Printf(string(result))
}

