package consumer

import (
	"encoding/json"
	"errors"
	. "fio"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GenderAPI struct {
	Count  int     `json:"count"`
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Prob   float64 `json:"probability"`
}

type AgeAPI struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type NationalityAPI struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryId string `json:"country_id"`
	} `json:"country"`
}

var httpClient *http.Client

func GetNationality(name string, n *NationalityAPI) error {
	url := "https://api.nationalize.io/?name=" + name
	return GetJSON(url, n)
}
func GetAge(name string, a *AgeAPI) error {
	url := "https://api.agify.io/?name=" + name
	return GetJSON(url, a)
}

func GetGender(name string, g *GenderAPI) error {
	url := "https://api.genderize.io/?name=" + name
	return GetJSON(url, g)
}

func GetJSON(url string, inter interface{}) error {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	log.Println(resp.Body)

	return json.NewDecoder(resp.Body).Decode(inter)
}

func (c *Consumer) Enrich(client Client) (Client, error) {
	var nation NationalityAPI
	var gender GenderAPI
	var age AgeAPI

	httpClient = &http.Client{
		Timeout: time.Second * 5,
	}

	if err := GetAge(client.Name, &age); err != nil {
		log.Print(err.Error())
		return client, err
	}
	if age.Age == 0 {
		return client, errors.New("invalid name")
	}

	if err := GetGender(client.Name, &gender); err != nil {
		log.Print(err.Error())
		return client, err
	}
	if gender.Gender == "" {
		return client, errors.New("invalid name")
	}

	if err := GetNationality(client.Name, &nation); err != nil {
		log.Print(err.Error())
		return client, err
	}
	if nation.Country == nil {
		return client, errors.New("invalid name")
	}

	enrichedClient := Client{
		Name:       client.Name,
		Surname:    client.Surname,
		Patronymic: client.Patronymic,
		Age:        age.Age,
		Gender:     gender.Gender,
		CountryId:  nation.Country[0].CountryId,
	}

	fmt.Printf("\nName: %s,\nSurname: %s,\nAge: %d,\nGender: %s,\nCountry: %s\n",
		enrichedClient.Name, enrichedClient.Surname, enrichedClient.Age, enrichedClient.Gender, enrichedClient.CountryId)

	return enrichedClient, nil
}
