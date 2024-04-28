package caltrain

func GetTripTypeFromRouteId(routeId string) string {
	var tripType string
	switch routeId {
	case "L2":
		tripType = "Weekend Local"
	case "L1":
		tripType = "Weekday Local"
	case "L5":
		tripType = "Weekday Limited"
	case "L3":
		tripType = "Weekday Limited"
	case "L4":
		tripType = "Weekday Limited"
	case "B7":
		tripType = "Weekday Bullet"
	}
	return tripType
}
