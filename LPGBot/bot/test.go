package bot

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"LPG-Bot/LPGBot/config"
	"net/url"
)



type PublicKey struct {
	Fact string `json:"fact"`
}

type KeysResponse struct {
	Collection []PublicKey
}


func testJoke() (err error) {
	res, err := http.Get("https://www.chucknorrisfacts.fr/api/get?data=tri:alea;type:img;nb:1")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	k := PublicKey{}
	keys := make([]PublicKey,0)
	json.Unmarshal([]byte(body), &keys)
	fmt.Println(keys[0])

	resp, err := http.PostForm(config.Webhook, url.Values{"content": {joke}, "tts": {"false"}})
	fmt.Println(resp)
}

