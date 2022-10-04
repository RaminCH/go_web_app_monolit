package render

import (
	"net/http"
	"testing"

	"github.com/RaminCH/bookings/internal/models"
)


func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash result of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{}) 
	if err != nil {
		t.Error("error writing error to browser")
	}

	err = RenderTemplate(&ww, r, "non-existing.page.tmpl", &models.TemplateData{}) 
	if err == nil {
		t.Error("rendered template that not exist")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil 

}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}


// xxx@xxx:~/go/src/Web_Part1/bookings/internal/render(master)$ go test -v
// === RUN   TestAddDefaultData
// --- PASS: TestAddDefaultData (0.00s)
// === RUN   TestRenderTemplate
// --- PASS: TestRenderTemplate (0.01s)
// === RUN   TestNewTemplates
// --- PASS: TestNewTemplates (0.00s)
// === RUN   TestCreateTemplateCache
// --- PASS: TestCreateTemplateCache (0.00s)
// PASS
// ok      github.com/RaminCH/bookings/internal/render     0.018s