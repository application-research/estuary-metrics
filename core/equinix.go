package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DeviceUsage struct {
	Usages []struct {
		Quantity float64 `json:"quantity"`
		Unit     string  `json:"unit"`
		Price    float64 `json:"price"`
		Total    float64 `json:"total"`
	} `json:"usages"`
}

func (m Metrics) GetDeviceInfo(deviceUUID string, createdDate string, createdBefore string) (DeviceUsage, error) {

	deviceUsage := DeviceUsage{}
	url := EquinixEndpoint +
		deviceUUID +
		"?created%5Bafter%5D=" + createdDate + "&created%5Bbefore%5D=" + createdBefore

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return DeviceUsage{}, err
	}
	req.Header.Add("X-Auth-Token", EquinixAuthToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return DeviceUsage{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return DeviceUsage{}, err
	}
	fmt.Println(string(body))
	json.Unmarshal(body, &deviceUsage)
	return deviceUsage, nil
}
