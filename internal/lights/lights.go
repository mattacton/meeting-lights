// Package lights handles manipulating physical lights.
//
// This package exclusively deals with calling Philips Hue
// APIs unless I find need to manipulate other types of lights
package lights

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type responseState struct {
	State LightState
}

type LightState struct {
	On  bool      `json:"on"`
	Bri int       `json:"bri"`
	Hue int       `json:"hue"`
	Sat int       `json:"sat"`
	XY  []float32 `json:"xy"`
}

type Lights struct {
	OriginalStates map[string]LightState
	Host           string
	APIKey         string
	LightIDs       []string
	client         *http.Client
}

func NewLights(host string, apiKey string, lightIds []string) *Lights {
	lights := new(Lights)
	lights.Host = host
	lights.APIKey = apiKey
	lights.LightIDs = lightIds
	lights.client = &http.Client{}
	return lights
}

func (lights Lights) baseAPIPath() string {
	return fmt.Sprintf("%s/api/%s/lights", lights.Host, lights.APIKey)
}

func getBytes(lightState LightState) []byte {
	json, err := json.Marshal(lightState)
	if err != nil {
		fmt.Printf("Problem converting state to json: %v\n", lightState)
		return []byte{}
	}

	return json
}

func (lights Lights) ResetLight() {
	for _, id := range lights.LightIDs {
		fmt.Printf("Resetting light %s to state %v\n", id, lights.OriginalStates[id])
		// No go routine as I don't need the reset to be quick and I don't feel like
		// implementing the means to have main wait for all of these routines to finish
		lights.setState(lights.OriginalStates[id], id)
	}
}

func (lights Lights) TurnRed(keys string) {
	redState := LightState{
		On:  true,
		Bri: 175,
		Hue: 0,
		Sat: 254,
		XY:  []float32{0.6915, 0.3083},
	}

	for _, id := range lights.LightIDs {
		go lights.setState(redState, id)
	}
}

func (lights Lights) TurnBlue(keys string) {
	redState := LightState{
		On:  true,
		Bri: 163,
		Hue: 45055,
		Sat: 233,
		XY:  []float32{0.1734, 0.1318},
	}

	for _, id := range lights.LightIDs {
		go lights.setState(redState, id)
	}
}

func (lights Lights) TurnGreen(keys string) {
	redState := LightState{
		On:  true,
		Bri: 175,
		Hue: 27159,
		Sat: 245,
		XY:  []float32{0.1821, 0.6353},
	}

	for _, id := range lights.LightIDs {
		go lights.setState(redState, id)
	}
}

func (lights Lights) TurnNormal(keys string) {
	redState := LightState{
		On:  true,
		Bri: 80,
		Hue: 8418,
		Sat: 140,
		XY:  []float32{0.4573, 0.4100},
	}

	for _, id := range lights.LightIDs {
		go lights.setState(redState, id)
	}
}

func (lights Lights) setState(state LightState, lightId string) {
	req, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("%s/%s/state", lights.baseAPIPath(), lightId),
		bytes.NewBuffer(getBytes(state)))
	if err != nil {
		fmt.Println("Problem creating request", err.Error())
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := lights.client.Do(req)
	if err != nil {
		fmt.Println("Problem with request", err.Error())
	}
	if response.StatusCode != 200 {
		fmt.Printf("Response was not 200 for changing state \n %+v \n %+v", state, response)
	}
}

func (lights Lights) GetCurrentState() map[string]LightState {
	originalStates := make(map[string]LightState)
	for _, id := range lights.LightIDs {
		response, err := http.Get(fmt.Sprintf("%s/%s", lights.baseAPIPath(), id))
		if err != nil {
			fmt.Println("Problem getting state", err.Error())
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Problem reading state body", err.Error())
		}

		responseData := responseState{}
		jsonErr := json.Unmarshal(data, &responseData)
		if jsonErr != nil {
			fmt.Println("Problem unmarshalling json", jsonErr.Error())
		}

		originalStates[id] = responseData.State
	}
	return originalStates
}
