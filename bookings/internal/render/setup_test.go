package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/RaminCH/bookings/internal/config"
	"github.com/RaminCH/bookings/internal/models"
	"github.com/alexedwards/scs/v2"
)

var session *scs.Session
var testApp config.AppConfig

func TestMain(m *testing.M) {

	// what am I going to put in the session
	gob.Register(models.Reservation{})

	//session
	//change it to true when in production mode
	testApp.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session
	//end session

	app = &testApp

	os.Exit(m.Run())
}
