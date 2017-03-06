package main

// go:generate sqlboiler postgres

import (
	"github.com/labstack/echo"
	"bytes"
	"github.com/rohanthewiz/echo_basic_exercise/template"
	log "gopkg.in/inconshreveable/log15.v2"
	_ "github.com/lib/pq"
	//"net/http"
	"database/sql"
	"github.com/rohanthewiz/echo_basic_exercise/models"
	//. "github.com/vattle/sqlboiler/queries/qm"
	"fmt"
	"encoding/json"
)

func main() {
	var pilotsList = []string {
		"Alice",
		"Tiger",
		"Tom",
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) (err error) {
		buf := new(bytes.Buffer)
		template.UserList(pilotsList, buf)
		c.HTMLBlob(200, buf.Bytes())
		
		//resp := c.Response()
		//resp.Write([]byte("Hello World"))
		//fmt.Fprint(c.Response(), "Hello World!")
		return
	})

	e.GET("/pilots", func(c echo.Context) (error) {
		var err error
		out_str := "Blank page"

		db, err := sql.Open("postgres", "dbname=pilots_development user=devuser password=devword port=32768 sslmode=disable")
		if err != nil {
			log.Error("Failed to open database", "Err:", err)
			return err
		}
		log.Info("Connected to database")

		p := &models.Pilot{
			Name: "John",
		}
		err = p.Insert(db)
		if err != nil {
			log.Error("Failed inserting first pilot into db", "Err", err)
			return err
		}
		log.Info("Inserted first pilot into db", "id", p.ID)

		p2 := &models.Pilot{
			Name: "George",
		}
		err = p2.Insert(db)
		if err != nil {
			log.Error("Failed to insert second pilot into db", "Err", err)
			return err
		}
		log.Info("Inserted second pilot into db", "id", p2.ID)

		pilots_count, err := models.Pilots(db).Count()
		if err != nil {
			log.Error("Could not get a count of pilots")
			return err
		}
		out_str = fmt.Sprintf("%d Pilots found in db", pilots_count)

		pilots, err := models.Pilots(db).All()
		if err != nil {
			log.Error("Unable to retrieve all pilots")
			return err
		}

		out_str += "\n"
		for _, pilot := range pilots {
			jstr, err := json.Marshal(pilot)
			if err != nil {
				log.Error("Unable to json encode pilot"); return err
			}
			out_str += string(jstr) + "\n"
		}

		c.String(200, out_str)
		log.Info(out_str)
		return err
	})


	e.Start(":3000")
}
