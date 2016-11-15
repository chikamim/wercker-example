package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL       = "http://www.apple.com/"
	productPrefix = "/jp/shop/product/"
)

type Product struct {
	Name              string
	URL               string
	Price             float64
	Description       string
	IsLanguageVariant bool
}

func GetSpecialMacProducts() ([]*Product, error) {
	specialDealsUrls := []string{
		"http://www.apple.com/jp/shop/browse/home/specialdeals/mac/macbook_pro/15",
	}

	products := []*Product{}
	for _, url := range specialDealsUrls {
		deals, err := GetSpecialDeals(url)
		if err != nil {
			return nil, err
		}
		for _, deal := range deals {
			product, err := GetProduct(deal)
			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	}

	return products, nil
}

func GetSpecialDeals(url string) (urls []string, err error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	set := map[string]bool{}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if strings.HasPrefix(href, productPrefix) {
			set[baseURL+href] = true
		}
	})

	for k, _ := range set {
		urls = append(urls, k)
	}

	return urls, nil
}

func GetProduct(url string) (*Product, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	product := &Product{}
	product.URL = url

	doc.Find("title").Each(func(_ int, s *goquery.Selection) {
		product.Name = s.Text()
	})

	doc.Find("meta[name='description']").Each(func(_ int, s *goquery.Selection) {
		product.Description, _ = s.Attr("content")
	})

	doc.Find("span[itemprop='price']").Each(func(_ int, s *goquery.Selection) {
		product.Price, _ = stringToFloat(s.Text())
	})

	doc.Find("li.as-pdp-prodvariations").Each(func(_ int, s *goquery.Selection) {
		if len(strings.TrimSpace(s.Text())) > 0 {
			product.IsLanguageVariant = true
		}
	})

	return product, nil
}

func stringToFloat(s string) (float64, error) {
	m := regexp.MustCompile("[0-9]+").FindAllString(s, -1)
	return strconv.ParseFloat(strings.Join(m, ""), 64)
}
