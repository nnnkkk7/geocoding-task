package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	geojson "github.com/paulmach/go.geojson"

	_ "github.com/lib/pq"
)

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

	fc, err := geojson.UnmarshalFeatureCollection(raw)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	const school = "P29_005"
	for i, ft := range fc.Features {
		f, err := ft.Geometry.MarshalJSON()
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
		_, err = db.Exec(fmt.Sprintf("insert into points (id, data, g) values('%d','%s', ST_GeomFromGeoJSON('%s'))", i, ft.Properties[school], f))
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	}

}
