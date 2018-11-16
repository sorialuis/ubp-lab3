package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

type APIResponseSite []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type APIResponseCat struct {
	ID                       string        `json:"id"`
	Name                     string        `json:"name"`
	Picture                  string        `json:"picture"`
	Permalink                string        `json:"permalink"`
	TotalItemsInThisCategory int           `json:"total_items_in_this_category"`
	PathFromRoot             []interface{} `json:"path_from_root"`
	ChildrenCategories       []struct {
		ID                       string `json:"id"`
		Name                     string `json:"name"`
		TotalItemsInThisCategory int    `json:"total_items_in_this_category"`
	} `json:"children_categories"`
	AttributeTypes string `json:"attribute_types"`
	Settings       struct {
	} `json:"settings"`
	MetaCategID  interface{} `json:"meta_categ_id"`
	Attributable bool        `json:"attributable"`
}

const (
	baseURL = "https://api.mercadolibre.com"
)

func main() {
	fmt.Print("\nIngresa un sitio: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	site := scanner.Text()

	url := GetUrlCat(strings.ToUpper(site))

	siteData, err := GetCategories(url)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	i := len(*siteData)

	wg.Add(i)

	for i, tt := range *siteData {
		idCat := tt.ID
		url2 := GetUrlSubC(idCat)

		categories, err := GetCategoriesInfo(url2)
		if err != nil {
			panic(err)
		}

		imageUrl := categories.Picture

		go SaveImage(imageUrl, fmt.Sprintf("file %d.png", i), &wg)

	}

	wg.Wait()
}

func GetUrlCat(id string) string {
	var buff bytes.Buffer
	buff.WriteString(baseURL)
	buff.WriteString("/sites/")
	buff.WriteString(id)
	buff.WriteString("/categories")

	return buff.String()
}

func GetUrlSubC(id string) string {
	var buff bytes.Buffer
	buff.WriteString(baseURL)
	buff.WriteString("/categories/")
	buff.WriteString(id)

	return buff.String()
}

func GetCategories(url string) (*APIResponseSite, error) {
	resp, err := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponseSite
	err = json.Unmarshal(data, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

func GetCategoriesInfo(url string) (*APIResponseCat, error) {
	resp, err := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponseCat
	err = json.Unmarshal(data, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

func SaveImage(imageUrl string, filename string, wg *sync.WaitGroup) error {
	resp, err := http.Get(imageUrl)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	wg.Done()
	return nil
}
