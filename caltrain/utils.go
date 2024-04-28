package caltrain

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// force timeout on requests
var netClient = &http.Client{
	Timeout: time.Second * 60,
}

// GetHandler returns the response body of a get request
func GetHandler(url string) RequestStatus {
	requestStatus := RequestStatus{}
	response, requestError := netClient.Get(url)
	if requestError != nil {
		requestStatus.ErrorMsg = "Error requesting data"
		requestStatus.Status = 1
		return requestStatus
	}

	defer response.Body.Close()
	responseBody, responseError := io.ReadAll(response.Body)

	if responseError != nil {
		requestStatus.ErrorMsg = "Error reading response body"
		requestStatus.Status = 1
		return requestStatus
	}

	requestStatus.Status = 0
	requestStatus.Data = responseBody

	return requestStatus
}

func RealTimeDecoder(responseBody []byte) CalTrainDecodeStatus {
	CalTrainDecodeStatus := CalTrainDecodeStatus{}

	var RealTime = RealTime{}

	jsonDecodeError := json.Unmarshal(responseBody, &RealTime)

	if jsonDecodeError != nil {
		CalTrainDecodeStatus.ErrorMsg = "Error decoding response body"
		CalTrainDecodeStatus.Status = 1
		CalTrainDecodeStatus.RealTime = RealTime
		return CalTrainDecodeStatus
	}

	CalTrainDecodeStatus.ErrorMsg = ""
	CalTrainDecodeStatus.Status = 0
	CalTrainDecodeStatus.RealTime = RealTime

	return CalTrainDecodeStatus
}
