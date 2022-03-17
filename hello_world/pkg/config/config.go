package config

import "text/template"

//AppConfig holds the application's config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
