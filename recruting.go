package strife

import (
	"encoding/json"
	"io/ioutil"
)

type Recruting struct {
	Phases *Phases `json:"phases"`
}

func ReadJSON(jsonPath string) (*Recruting, error) {
	bytes, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	var recruting Recruting
	if err := json.Unmarshal(bytes, &recruting); err != nil {
		return nil, err
	}
	return &recruting, nil
}
