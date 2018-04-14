package bot

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../config"
)


type JokeBody struct {
	Id	int	`json:"id"`
	Joke	string	`json:"joke"`
	Date	int	`json:"date"`
	Vote	int	`json:"vote"`
	Points	int `json:"points"`
}

//type Joke []JokeBody

//Send joke
func sendJoke() (err error) {
	// Get a Chuck Norris joke
	joke, err := getJoke()
	if err != nil {
		return err
	}
	// Send
	resp, err := http.PostForm(config.Webhook, url.Values{"content": {joke}, "tts": {"false"}})
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Couldn't send message")
		fmt.Println(err)
		return err
	} else {
		fmt.Println(resp)
		return err
	}

	return nil
}

//Fetch Chuck Norris Joke
func getJoke() (string, error) {
	resp, err := http.Get("http://www.chucknorrisfacts.fr/api/get?data=tri:alea;type:txt;nb:1")
	if err != nil {
		fmt.Println("Could not fetch joke")
		return "nil", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unknown response body")
		return "nil", err
	}

	var jokeResp JokeBody
	json.Unmarshal(body, &jokeResp)
	fmt.Println(jokeResp)
	return jokeResp.Joke, nil
}