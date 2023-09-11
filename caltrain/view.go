package caltrain

import (
	"fmt"
	"time"
)

func GetRealTime() RealTimeStatus {
	status := RealTimeStatus{}
	timestamp := time.Now().UnixMilli()
	caltrainUrl := fmt.Sprintf("https://www.caltrain.com/files/rt/tripupdates/CT.json?time=%d", timestamp)
	requestStatus := GetHandler(caltrainUrl)
	if requestStatus.Status == 1 {
		status.ErrorMsg = requestStatus.ErrorMsg
		status.Status = requestStatus.Status
		return status
	}

	decodeStatus := RealTimeDecoder(requestStatus.Data)
	if decodeStatus.Status == 1 {
		status.ErrorMsg = decodeStatus.ErrorMsg
		status.Status = requestStatus.Status
		return status
	}

	status.RealTime = decodeStatus.RealTime
	return status

}
