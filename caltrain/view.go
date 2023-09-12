package caltrain

import (
	"fmt"
	"strconv"
	"time"
)

func GetRealTime() RealTimeStatus {
	status := RealTimeStatus{}
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		status.ErrorMsg = "Could not load timezone"
		status.Status = 1
		return status
	}

	timestamp := time.Now().In(location).UnixMilli()
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

func ParseRealTime() CalTrainTripStatus {
	status := CalTrainTripStatus{}
	real_time_status := GetRealTime()
	if real_time_status.Status == 1 {
		status.ErrorMsg = real_time_status.ErrorMsg
		status.Status = real_time_status.Status
		return status
	}
	entityCount := len(real_time_status.RealTime.Entities)
	var calTrainTrips []CalTrainTrip
	for i := 0; i < entityCount; i++ {
		entity := real_time_status.RealTime.Entities[i]
		calTrainTrip := CalTrainTrip{Id: entity.Id, RouteId: entity.TripUpdate.Trip.RouteId, DirectionId: entity.TripUpdate.Trip.DirectionId}
		stopCount := len(entity.TripUpdate.StopTimeUpdate)
		var stops []Stop
		for j := 0; j < stopCount; j++ {
			stopTimeUpdate := entity.TripUpdate.StopTimeUpdate[j]
			stopIdInt, err := strconv.Atoi(stopTimeUpdate.StopId)
			if err != nil {
				status.ErrorMsg = err.Error()
				status.Status = 1
				return status
			}
			stop := Stop{Id: stopIdInt, Arrival: stopTimeUpdate.Arrival.Time, Departure: stopTimeUpdate.Departure.Time}
			stops = append(stops, stop)
		}
		calTrainTrip.Stops = stops
		calTrainTrips = append(calTrainTrips, calTrainTrip)
	}
	status.CalTrainTrips = calTrainTrips
	return status
}

func ParseCalTrainStop(origin int) CalTrainVehicleStatus {
	status := CalTrainVehicleStatus{}
	calTrainTripStatus := ParseRealTime()
	if calTrainTripStatus.Status == 1 {
		status.ErrorMsg = calTrainTripStatus.ErrorMsg
		status.Status = calTrainTripStatus.Status
		return status
	}

	var calTrainVehicles []CalTrainVehicle
	for i := 0; i < len(calTrainTripStatus.CalTrainTrips); i++ {
		calTrainTrip := calTrainTripStatus.CalTrainTrips[i]
		var currentStop int
		for j := 0; j < len(calTrainTrip.Stops); j++ {
			calTrainTripStop := calTrainTrip.Stops[j]
			if j == 0 {
				currentStop = calTrainTripStop.Id
				continue
			}
			if calTrainTripStop.Id == origin {
				calTrainVehicle := CalTrainVehicle{
					TrainId:       calTrainTrip.Id,
					ArrivalTime:   calTrainTripStop.Arrival,
					DepartureTime: calTrainTripStop.Departure,
					StopsLeft:     j,
					CurrentStop:   currentStop,
					TripType:      calTrainTrip.RouteId,
					Direction:     calTrainTrip.DirectionId,
				}
				calTrainVehicles = append(calTrainVehicles, calTrainVehicle)
				break
			}
		}
	}

	status.CalTrainVehicles = calTrainVehicles

	return status
}
