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
	"net/http"
	"time"
)

func main() {
	var pilotsList = []string {
		"Alice",
		"Tiger",
		"Tom",
	}

	e := echo.New()

	admin := e.Group("/admin", RoAuth)
//	or: admin.Use(RoAuth)
	admin.GET("/", func(c echo.Context) error{
		c.String(200, "Hello Admin!")
		return nil
	})

	e.GET("/", func(c echo.Context) (err error) {
		buf := new(bytes.Buffer)
		template.UserList(pilotsList, buf)
		c.HTMLBlob(200, buf.Bytes())
		
		//resp := c.Response()
		//resp.Write([]byte("Hello World"))
		//fmt.Fprint(c.Response(), "Hello World!")
		return
	})
	e.GET("/login", func(c echo.Context) error {
		c.String(http.StatusOK, "We would login here")

		return nil
	})

	// Simulate a good login
	e.GET("/auth", func(c echo.Context) error {
		fmt.Println("In /auth...")
		sess, err := c.Cookie("sess_id")
		if err != nil || sess.Value == "" {
			cookie := new(http.Cookie)
			cookie.Name = "sess_id"
			cookie.Value = "xY78jq"
			cookie.Expires = time.Now().Add(24 * time.Hour)
			c.SetCookie(cookie)
			fmt.Println("Set auth cookie")

			// todo create a key in Redis with a TTL (https://github.com/go-redis/redis)
			// Perhaps use a redis hash
			// key: { user_id: "nnn", first_name: "john", last_name: "Smith" }
		}
		// todo for normal login: https://godoc.org/golang.org/x/crypto/scrypt
		return nil
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
			Hobbies: []string{ "flying", "reading"},
		}
		err = p.Insert(db)
		if err != nil {
			log.Error("Failed inserting first pilot into db", "Err", err)
			return err
		}
		log.Info("Inserted first pilot into db", "id", p.ID)

		p2 := &models.Pilot{
			Name: "George",
			Hobbies: []string{ "singing", "dancing" },
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

// Auth middleware
func RoAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get session_id from cookie
		sess, err := c.Cookie("sess_id")
		if err != nil {
			redirect_to_login(c)
			return nil
		}
		if sess.Value == "" { redirect_to_login(c) }

		// Retrieve session from store
		session, err := get_session(sess.Value)
		if err != nil {
			redirect_to_login(c)
			return nil
		}
		// If session.blank?: redirect to login
		extend_session(session, 3600)
		return next(c)
	}
}

// todo
func get_session(session_id string) (map[string]string, error) {
	var err error
	fmt.Println("Getting session for:", session_id)
	session := make(map[string]string)  // bogus

	return session, err
}

// todo
func extend_session(session map[string]string, ext int64) {
	fmt.Println("Extending session:", session, "by", ext, "seconds")
	// todo extend TTL of key in Redis
}

func redirect_to_login(c echo.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/login")
}
