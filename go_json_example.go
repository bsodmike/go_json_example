/*
 * Michael de Silva <michael@mwdesilva.com>
 * CEO @ http://omakaselabs.com/ / https://github.com/bsodmike
 */

package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "bytes"
)

// GH user type matching response JSON object.
type GithubUser struct {
  Login string
  Name string
  Email string
  Company string
  Blog string
}

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

func GetGithubUser(username string) (*GithubUser, error) {
  var buffer bytes.Buffer
  url := "https://api.github.com/users/"
  buffer.WriteString(url)
  buffer.WriteString(username)
  payload, _ := getContent(buffer.String())

  var err error
  var record GithubUser
  err = json.Unmarshal(payload, &record)
  if err != nil {
    return nil, err
  }
  return &record, err
}

func main() {
  username := "bsodmike"
  record, _ := GetGithubUser(username)
  fmt.Printf("Output\n")
  fmt.Printf("Info for Github user '%s':\n%v\n", username, record)
}

