package main

import "github.com/ArunArivanandam/snippet/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
