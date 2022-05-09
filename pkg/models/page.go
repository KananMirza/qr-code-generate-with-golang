package models

type Page struct {
	IsAlert      bool
	AlertTitle   string
	AlertClass   string
	AlertContent interface{}
}
