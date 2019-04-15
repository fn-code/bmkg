package bmkg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"time"

	"net/http"
)

// Infogempa get data infogempa from bmkg
type Infogempa struct {
	Gempa []Gempa `xml:"gempa"`
}

// Gempa get gempa data from bmkg
type Gempa struct {
	Tanggal   string     `xml:"Tanggal"`
	Jam       string     `xml:"Jam"`
	Point     Coordinate `xml:"point"`
	Lintang   string     `xml:"Lintang"`
	Bujur     string     `xml:"Bujur"`
	Magnitude string     `xml:"Magnitude"`
	Kedalaman string     `xml:"Kedalaman"`
	Symbol    string     `xml:"_symbol"`
	Wilayah   string     `xml:"Wilayah"`
}

// Coordinate get coordinate data or location form bmkg
type Coordinate struct {
	Coordinates string `xml:"coordinates"`
}

// GempaTerkini return gempa data from bmkg
func GempaTerkini(url string) (*Infogempa, error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error http bad request: %v", http.StatusBadRequest)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var info Infogempa
	err = xml.Unmarshal(buf, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
