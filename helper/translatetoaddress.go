package helper

import (
	"context"
	"fmt"
	"log"
	"os"

	"googlemaps.github.io/maps"
)

func TranslateToAddress(latitude, longitude float64) (string, error) {
	mapsApiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	client, err := maps.NewClient(maps.WithAPIKey(mapsApiKey))
	if err != nil {
		log.Fatalf("Kesalahan pembuatan klien: %v", err)
	}

	// Mengganti latitude dan longitude sesuai dengan koordinat Anda
	latLng := &maps.LatLng{
		Lat: latitude,
		Lng: longitude,
	}

	// Gunakan klien untuk melakukan reverse geocoding
	results, err := client.ReverseGeocode(context.Background(), &maps.GeocodingRequest{
		LatLng: latLng,
	})
	if err != nil {
		log.Fatalf("Kesalahan saat reverse geocoding: %v", err)
	}

	var address string
	// Ambil alamat dari hasil reverse geocoding
	if len(results) > 0 {
		address = results[0].FormattedAddress
		fmt.Printf("Alamat: %s\n", address)
	} else {
		fmt.Println("Tidak ada hasil reverse geocoding yang ditemukan.")
	}
	return address, err
}
