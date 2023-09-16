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
	realTimeStatus := GetRealTime()
	if realTimeStatus.Status == 1 {
		status.ErrorMsg = realTimeStatus.ErrorMsg
		status.Status = realTimeStatus.Status
		return status
	}

	var trips []CalTrainTrip
	for _, entity := range realTimeStatus.RealTime.Entities {
		trip := CalTrainTrip{Id: entity.Id, RouteId: entity.TripUpdate.Trip.RouteId, DirectionId: entity.TripUpdate.Trip.DirectionId}
		var stops []Stop
		for _, stopTimeUpdate := range entity.TripUpdate.StopTimeUpdate {
			stopIdInt, err := strconv.Atoi(stopTimeUpdate.Id)
			if err != nil {
				status.ErrorMsg = err.Error()
				status.Status = 1
				return status
			}
			stop := Stop{Id: stopIdInt, Arrival: stopTimeUpdate.Arrival.Time, Departure: stopTimeUpdate.Departure.Time}
			stops = append(stops, stop)
		}
		trip.Stops = stops
		trips = append(trips, trip)
	}
	status.CalTrainTrips = trips
	return status
}

func ParseCalTrainStop(originN int, originS int) CalTrainVehicleStatus {
	status := CalTrainVehicleStatus{}
	calTrainTripStatus := ParseRealTime()
	if calTrainTripStatus.Status == 1 {
		status.ErrorMsg = calTrainTripStatus.ErrorMsg
		status.Status = calTrainTripStatus.Status
		return status
	}

	var vehicles []CalTrainVehicle
	for _, trip := range calTrainTripStatus.CalTrainTrips {
		var currentStop int
		for j, stop := range trip.Stops {
			if j == 0 {
				currentStop = stop.Id
				continue
			}
			if stop.Id == originN || stop.Id == originS {
				direction := "Northbound"
				if trip.DirectionId == 1 {
					direction = "Southbound"
				}
				stopName := GetStopById(currentStop)
				vehicle := CalTrainVehicle{
					Id:            trip.Id,
					ArrivalTime:   stop.Arrival,
					DepartureTime: stop.Departure,
					StopsLeft:     j,
					CurrentStop:   stopName.Name,
					TripType:      trip.RouteId,
					Direction:     direction,
				}
				vehicles = append(vehicles, vehicle)
				break
			}
		}
	}

	status.CalTrainVehicles = vehicles

	return status
}
