package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// https://github.com/2dust/v2rayN/wiki/%E5%88%86%E4%BA%AB%E9%93%BE%E6%8E%A5%E6%A0%BC%E5%BC%8F%E8%AF%B4%E6%98%8E(ver-2)
type Data struct {
	Version    string `json:"v"`
	PostScript string `json:"ps"`
	Port       string `json:"port"`
	Address    string `json:"add"`
	Id         string `json:"id"`
	AlterId    string `json:"aid"`
	Scy        string `json:"scy"`
	Net        string `json:"net"`
	Type       string `json:"type"`
	Host       string `json:"host"`
	Path       string `json:"path"`
	Tls        string `json:"tls"`
	Sni        string `json:"sni"`
}

// Output struct as json
func (data *Data) String() string {
	b, err := json.Marshal(*data)
	if err != nil {
		return fmt.Sprintf("%+v", *data)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *data)
	}
	return out.String()
}

func main() {
	resp, _ := http.Get("https://ednovas.tech/api/v1/client/subscribe?token=9a466af050d85aa1356771b4d566075b")
	contentBytes, _ := ioutil.ReadAll(resp.Body)

	// get content bytes from file
	// file, err := os.Open("text")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// contentBytes, err := ioutil.ReadAll(file)

	contentString := string(contentBytes)
	decodedContentBytes, _ := base64.StdEncoding.DecodeString(contentString)
	decodedContentString := string(decodedContentBytes)
	items := strings.Split(decodedContentString, "\n")
	for _, value := range items {
		// create an empty map, use bultin make
		elements := strings.Split(value, "://")
		detailBytes, _ := base64.StdEncoding.DecodeString(elements[1])
		// detailString := string(detailBytes)
		data := Data{}
		err := json.Unmarshal(detailBytes, &data)
		if err != nil {
			// panic(err)
		}
		fmt.Println(data.String())
	}
}
