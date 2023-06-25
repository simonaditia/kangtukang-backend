// main_done_but_not_with_astarmanual
package models

import (
	"context"
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"

	"googlemaps.github.io/maps"
)

// GG dah bisa, tapi ini gk menerapkan algortima astar sendiri
// tetapi pakai library google maps yang menentukan jarak tercepat/terpendek
func FindTukangWithAStarGmaps(start, end *maps.LatLng) (float64, error) {
	// Inisialisasi klien Google Maps
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyDI-XccW0sqI-nyBZz97iRprnSu01oLMXA"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Lokasi awal dan tujuan
	// start := &maps.LatLng{Lat: -6.2345036, Lng: 106.9840831}
	// end := &maps.LatLng{Lat: -6.2376927, Lng: 106.9727969}
	// end := &maps.LatLng{Lat: -6.2507407, Lng: 106.9640273}

	numWaypoints := 3

	// Menentukan rentang latitude dan longitude untuk waypoint acak
	minLat := math.Min(start.Lat, end.Lat)
	maxLat := math.Max(start.Lat, end.Lat)
	minLng := math.Min(start.Lng, end.Lng)
	maxLng := math.Max(start.Lng, end.Lng)

	// Menentukan waypoint acak
	waypoints := make([]maps.LatLng, numWaypoints)
	for i := 0; i < numWaypoints; i++ {
		lat := rand.Float64()*(maxLat-minLat) + minLat
		lng := rand.Float64()*(maxLng-minLng) + minLng
		waypoints[i] = maps.LatLng{Lat: lat, Lng: lng}
	}

	// Mengonversi waypoints menjadi tipe []string
	waypointStrings := make([]string, len(waypoints))
	for i, waypoint := range waypoints {
		waypointStrings[i] = waypoint.String()
	}

	// Membuat permintaan rute menggunakan algoritma A*
	r := &maps.DirectionsRequest{
		Origin:        start.String(),
		Destination:   end.String(),
		Mode:          maps.TravelModeDriving,
		Alternatives:  true,                       // Meminta beberapa rute
		DepartureTime: "now",                      // Menggunakan waktu saat ini
		TrafficModel:  maps.TrafficModelBestGuess, // Estimasi berdasarkan situasi lalu lintas saat ini
	}

	// Mengirim permintaan rute ke API Google Maps
	routes, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("Failed to request directions: %v", err)
	}

	// Menyimpan rute-rute dalam struktur data untuk perbandingan
	var routeDurations []struct {
		Route    *maps.Route
		Duration int
	}

	// Menyimpan simpul-simpul koordinat dalam map
	coordinatesMap := make(map[*maps.LatLng]int)

	// Menghitung estimasi waktu tempuh masing-masing rute dan menyimpan simpul-simpul koordinat dalam map
	for i, route := range routes {
		legs := route.Legs
		duration := 0

		for _, leg := range legs {
			duration += int(leg.DurationInTraffic.Seconds())

			// Menyimpan simpul koordinat dalam map
			coordinates := &leg.StartLocation
			coordinatesMap[coordinates] = i

			steps := leg.Steps
			for _, step := range steps {
				coordinates = &step.StartLocation
				coordinatesMap[coordinates] = i
				coordinates = &step.EndLocation
				coordinatesMap[coordinates] = i
			}

			coordinates = &leg.EndLocation
			coordinatesMap[coordinates] = i
		}

		// Menyimpan durasi estimasi waktu tempuh dalam routeDurations
		routeDurations = append(routeDurations, struct {
			Route    *maps.Route
			Duration int
		}{Route: &route, Duration: duration})
	}

	// Mengurutkan rute berdasarkan estimasi waktu tempuh terpendek
	sort.Slice(routeDurations, func(i, j int) bool {
		return routeDurations[i].Duration < routeDurations[j].Duration
	})

	// Mendapatkan rute terpendek
	shortestRoute := routes[0]
	shortestDistance := shortestRoute.Legs[0].Distance.Meters

	for i := 1; i < len(routes); i++ {
		distance := routes[i].Legs[0].Distance.Meters
		if distance < shortestDistance {
			shortestRoute = routes[i]
			shortestDistance = distance
		}
	}

	// Menampilkan hasil
	/*fmt.Println("Rute Terdekat ada", len(routes), "rute:")
	for i, route := range routes {
		fmt.Println("Rute", i+1, ":")
		fmt.Println("- Jarak:", route.Legs[0].Distance.HumanReadable)
		fmt.Println("- Start Address:", route.Legs[0].StartAddress)
		for _, step := range route.Legs[0].Steps {
			fmt.Println("  -", step.String(), "", step.StartLocation.Lat, "", step.StartLocation.Lng)
		}
		fmt.Println("- End Address:", route.Legs[0].EndAddress)
	}

	fmt.Println()
	fmt.Println("Rute paling terdekat adalah rute", shortestRoute.Summary, "dengan jarak", shortestRoute.Legs[0].Distance.HumanReadable)
	*/

	distance := shortestRoute.Legs[0].Distance.HumanReadable

	return convertDistanceToFloat64(distance), err
}

func convertDistanceToFloat64(distance string) float64 {
	// Menghapus karakter non-digit dari string jarak
	distance = strings.ReplaceAll(distance, " km", "")
	distance = strings.ReplaceAll(distance, ",", "")

	// Mengonversi string menjadi int
	distanceFloat64, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		return 0
	}

	return distanceFloat64
}
