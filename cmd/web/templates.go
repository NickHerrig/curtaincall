package main

import "curtaincall.tech/pkg/models"

type templateData struct {
    Theater  *models.Theater
    Theaters []*models.Theater
}
