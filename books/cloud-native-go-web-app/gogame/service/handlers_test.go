package service

import (
	"bytes"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

func TestCreateMatch(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(http.HandlerFunc(createMatchHandler(formatter)))
	defer server.Close()

	body := []byte(
		`{
  "gridsize": 19,
  "players": [
    {
      "color": "white",
      "name": "bob"
    },
    {
      "color": "black",
      "name": "alfred"
    }
  ]
}`)
	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST request for createMatchHandler: %v", err)
	}
	req.Header.Add("Countet-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST createMatchHandler: %v", err)
	}

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected response status 201, received %s", res.Status)
	}

	if _, headerOk := res.Header["Location"]; !headerOk {
		t.Error("Location header is not set")
	}

}
