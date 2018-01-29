package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "os"
  "bytes"
)

func check(e error) {
  if e != nil {
    fmt.Printf("%s", e)
    os.Exit(1)
  }
}

func writeToFile (m map[string]int) {
  b := new(bytes.Buffer)
  for key, value := range m {
    fmt.Fprintf(b, "%s:%d \n", key, value)
  }
  file, err := os.Create("result.txt")
  check(err)
  fmt.Fprintf(file, b.String())
}

func getData(url string) string {
  var data string
  response, err := http.Get(url)
  check(err)
  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)
  check(err)
  data = string(contents)
  fmt.Printf("%s\n", data)
  pattern := strings.NewReplacer(",", " ", ";", " ", "(", " ", ")", " ", "{", " ", "}", " ", ".", " ", "!", " ")
  processedData := pattern.Replace(data)
  return processedData
}

func wordCount(str string) {
  m := make(map[string]int)
  for _, f := range strings.Fields(str) {
    m[f] = m[f]+1
  }
  fmt.Printf("%v\n", m)
  writeToFile(m)
}

func main() {
  data := getData("http://www.gutenberg.org/files/15/text/moby-000.txt")
  wordCount(data)
}
