package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr-esr-scanr-publications-scientifiques&q=plasma&facet=type_de_publication&facet=numero_national_de_structure_de_recherche&facet=date_de_publication&facet=type_de_la_source&q="
)

type Dataset struct {
	NHits   int        `json:"nhits"`
	Records []Document `json:"records"`
}

type Document struct {
	DatasetID       string `json:"datasetid"`
	RecordID        string `json:"recordid"`
	RecordTimestamp string `json:"record_timestamp"`
	Field           Fields `json:"fields"`
}

type Fields struct {
	TypeDePublication                    string `json:"type_de_publication"`
	Thematiques                          string `json:"thematiques"`
	PrenomsDesAuteurs                    string `json:"prenoms_des_auteurs"`
	DateDePublication                    string `json:"date_de_publication"`
	NumeroNationalDeStructureDeRecherche string `json:"numero_national_de_structure_de_recherche"`
	Lien                                 string `json:"lien"`
	ReferenceHAL                         string `json:"reference_hal"`
	TypeDeLaSource                       string `json:"type_de_la_source"`
	TitreDeLaSource                      string `json:"titre_de_la_source"`
	Titre                                string `json:"titre"`
	NomsDesAuteurs                       string `json:"noms_des_auteurs"`
	ReferencesArchivesOAI                string `json:"references_archives_oai"`
	Resume                               string `json:"resume"`
}

func getPublications(name string) Dataset {
	url := baseURL + name
	req, err := http.NewRequest("GET", url, nil)
	check(err)

	// Handle the request
	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	var dataset Dataset
	json.Unmarshal(body, &dataset)
	return dataset
}

type Publications interface {
	getPublications(name string) Dataset
}

func DisplayPublications(name string) {
	name = cleanQuotes(name)
	fmt.Printf("Getting publications: %s\n", name)
	dataset := getPublications(name)
	for _, document := range dataset.Records {
		fmt.Println()
		fmt.Println(`*************************** Publication ***************************`)
		fmt.Println(`Date:                `, document.Field.DateDePublication)
		fmt.Println(`Auteurs:             `, document.Field.NomsDesAuteurs)
		fmt.Println(`ReferenceHAL:        `, document.Field.ReferenceHAL)
		fmt.Println(`Thematiques:         `, document.Field.Thematiques)
		fmt.Println(`Titre:               `, document.Field.Titre)
		fmt.Println(`Resume:              `, cleanTags(document.Field.Resume))
		fmt.Println(`Numero national de structure de recherche:         `, document.Field.NumeroNationalDeStructureDeRecherche)
		fmt.Println(`References archives OAI:                           `, document.Field.ReferencesArchivesOAI)
	}
}
