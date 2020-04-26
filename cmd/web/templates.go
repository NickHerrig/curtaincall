package main

import (
    "html/template"
    "path/filepath"

    "curtaincall.tech/pkg/forms"
    "curtaincall.tech/pkg/models"
)

type templateData struct {
    CurrentYear     int
    Flash           string
    Form            *forms.Form
    IsAuthenticated bool
    Theater         *models.Theater
    Theaters        []*models.Theater
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
    cache := map[string]*template.Template{}

    pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
    if err != nil {
        return nil, err
    }

    for _, page := range pages {
        name := filepath.Base(page)

        ts, err := template.ParseFiles(page)
        if err != nil {
            return nil, err
        }

        ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
        if err != nil {
            return nil, err
        }

        ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
        if err != nil {
            return nil, err
        }
        cache[name] = ts
    }

    return cache, nil

}
