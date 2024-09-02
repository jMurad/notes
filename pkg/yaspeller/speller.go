package yaspeller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Result []struct {
	Code int      `json:"code,omitempty"`
	Pos  int      `json:"pos,omitempty"`
	Row  int      `json:"row,omitempty"`
	Col  int      `json:"col,omitempty"`
	Len  int      `json:"len,omitempty"`
	Word string   `json:"word,omitempty"`
	S    []string `json:"s,omitempty"`
}

const uri = "https://speller.yandex.net/services/spellservice.json/checkText"

func CheckText(text string) (*Result, error) {
	data := url.Values{}
	data.Set("text", text)
	data.Set("options", "518")
	body := data.Encode()

	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body)))

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	res := new(Result)
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (res *Result) IsCorrect() bool {
	// res, _ := CheckText(text)

	if len(*res) == 0 {
		return true
	} else {
		return false
	}
}

func (res *Result) RightText(text string) string {
	// res, err := CheckText(text)
	// if err != nil {
	// 	return ""
	// }
	uc := []rune(text)
	for i := len(*res) - 1; i >= 0; i-- {
		r := (*res)[i]
		begin := string(uc[:r.Pos])
		middle := strings.Replace(string(uc[r.Pos:r.Pos+r.Len]), r.Word, r.S[0], 1)
		end := string(uc[r.Pos+r.Len:])

		uc = []rune(begin + middle + end)
	}

	return string(uc)
}
