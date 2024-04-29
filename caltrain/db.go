package caltrain

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const createTripsTable string = `
	CREATE TABLE IF NOT EXISTS trips (
		id TEXT NOT NULL PRIMARY KEY,
		route_id TEXT NOT NULL,
		direction_id INTEGER NOT NULL
	);`

const createTripStopsTable string = `
	CREATE TABLE IF NOT EXISTS trip_stops (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		trip_id TEXT NOT NULL,
		stop_id INTEGER NOT NULL,
		arrival INTEGER NOT NULL,
		departure INTEGER NOT NULL
	);`

func getDB() *sql.DB {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	return sqliteDatabase
}

func createTable(db *sql.DB, createStatement string) {
	statement, err := db.Prepare(createStatement)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func setupDB() {
	os.Remove("sqlite-database.db")
	log.Printf("Creating sqlite-database.db...\n")
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Printf("sqlite-database.db created\n")
	db := getDB()
	createTable(db, createTripsTable)
	createTable(db, createTripStopsTable)
	db.Close()
}

func insertIntoTripsTable(id string, routeId string, DirectionId int) {
	db := getDB()
	insertTripSQL := `
	INSERT INTO trips(
		id,
		route_id,
		direction_id)
	VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertTripSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec(id, routeId, DirectionId)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Close()
}

func insertIntoTripStopsTable(tripId string, stopId int, arrival int64, departure int64) {
	db := getDB()
	insertTripStopsSQL := `
	INSERT INTO trip_stops(
		trip_id,
		stop_id,
		arrival,
		departure)
	VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertTripStopsSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec(tripId, stopId, arrival, departure)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Close()
}

func getTrainsArrivingAtStops(stopNId int, stopSId int) []CalTrainVehicle {
	var vehicles []CalTrainVehicle
	db := getDB()
	selectTrainsSQL := `
	WITH trips_id_arrivals AS (
		SELECT trip_stops.arrival,
			   trip_stops.departure,
			   trips.id AS trip_id,
			   trips.route_id,
			   trips.direction_id
		  FROM trip_stops
			   JOIN
			   trips ON trip_stops.trip_id = trips.id
		 WHERE stop_id IN (70221, 70222)
	),
	current_stops AS (
		SELECT trip_id,
			   stop_id
		  FROM trip_stops
		 GROUP BY trip_id
		 ORDER BY arrival
	),
	stops_remaining AS (
		SELECT trip_stops.trip_id,
			   COUNT(trip_stops.stop_id) - 1 AS stops_remaining
		  FROM trip_stops
			   JOIN
			   trips_id_arrivals ON trip_stops.trip_id = trips_id_arrivals.trip_id
		 WHERE trip_stops.arrival <= trips_id_arrivals.arrival
		 GROUP BY trip_stops.trip_id
	)
	SELECT current_stops.stop_id AS current_stop_id,
		   stops_remaining.stops_remaining,
		   trips_id_arrivals.arrival,
		   trips_id_arrivals.departure,
		   trips_id_arrivals.trip_id,
		   trips_id_arrivals.route_id,
		   trips_id_arrivals.direction_id
	  FROM trips_id_arrivals
		   LEFT JOIN
		   current_stops USING (
			   trip_id
		   )
		   LEFT JOIN
		   stops_remaining USING (
			   trip_id
		   )`
	row, err := db.Query(selectTrainsSQL, stopNId, stopSId)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer row.Close()
	for row.Next() {
		var currentStopId int
		var stopsLeft int
		var arrival int64
		var departure int64
		var tripId string
		var routeId string
		var directionId int
		row.Scan(&currentStopId, &stopsLeft, &arrival, &departure, &tripId, &routeId, &directionId)
		direction := "Northbound"
		if directionId == 1 {
			direction = "Southbound"
		}
		stopName := GetStopById(currentStopId)
		tripType := GetTripTypeFromRouteId(routeId)
		vehicle := CalTrainVehicle{
			Id:            tripId,
			ArrivalTime:   arrival,
			DepartureTime: departure,
			StopsLeft:     stopsLeft,
			CurrentStop:   stopName.Name,
			TripType:      tripType,
			Direction:     direction,
		}
		vehicles = append(vehicles, vehicle)
	}
	return vehicles
}
