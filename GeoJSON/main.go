package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometory  Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	Name string `json:"P29_005"`
	// ID      string `json:"P17_002"`
	// Type    int    `json:"P29_001"`
	// Address string `json:"P29_006"`
}

type Geometry struct {
	// Type        string   `json:"type"`
	Coordinates []LatLng `json:"coordinates"`
}

type LatLng float64

func main() {
	db, err := sql.Open("postgres", "dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	defer db.Close()

	raw, err := ioutil.ReadFile("./schools.json")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	var fc FeatureCollection

	err = json.Unmarshal(raw, &fc)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	for i, ft := range fc.Features {
		_, err = db.Exec(fmt.Sprintf("insert into points (id, data, g) values('%d','%s', ST_GeomFromText('POINT(%f %f)',4326))", i, ft.Properties.Name, ft.Geometory.Coordinates[0], ft.Geometory.Coordinates[1]))
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	}

}
