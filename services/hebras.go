package services

import (
	"fmt"
	"github.com/gocolly/colly"
	"hebras-scrapping/constants"
	"hebras-scrapping/models"
	"strconv"
	"strings"
	"sync"
)

func ScrapeHebras(teaBlendsUrls []string) (teaHebras []models.HebrasTea) {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	for _, url := range teaBlendsUrls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			c := colly.NewCollector()
			scrapeTeaBlends(&teaHebras, c, mutex)
			c.OnRequest(func(r *colly.Request) {
				r.Headers.Set("User-Agent", constants.USER_AGENT)
			})
			c.Visit(url)
			c.Wait() //Este wait va a ser individual para cada gorooutine
		}(url)

	}

	wg.Wait()
	//go func() {
	//	//TODO: Investigar diferencia entre poner en una goroutine y no ponerlo en goroutine embembido el wgWait
	//	wg.Wait()
	//}()

	return teaHebras
}

// ----------------------------HELPERS--------------------------------

// *esto mismo para el link: teablendsforyou.com.ar
func scrapeTeaBlends(teaHebras *[]models.HebrasTea, c *colly.Collector, mutex *sync.Mutex) {
	c.OnHTML("li[data-hook=product-list-grid-item]", func(e *colly.HTMLElement) {
		tea := new(models.HebrasTea)
		aContainer := e.DOM.Nodes[0].FirstChild.FirstChild.FirstChild
		imgContainer := e.DOM.Nodes[0].FirstChild.FirstChild.FirstChild.FirstChild.FirstChild.FirstChild.FirstChild.FirstChild.Attr[0]
		if imgContainer.Key == "src" && imgContainer.Val != "" {
			tea.Img = imgContainer.Val
		}
		if aContainer.Data == "a" && aContainer.Attr[0].Key == "href" {
			tea.Link = aContainer.Attr[0].Val
		}
		if e.ChildText("h3[data-hook=product-item-name]") != "" {
			tea.Name = e.ChildText("h3[data-hook=product-item-name]")
		}
		if e.ChildText("span[data-hook=price-range-from]") != "" {
			tea.Price = e.ChildText("span[data-hook=price-range-from]")
		} else {
			tea.Price = e.ChildText("span[data-hook=product-item-price-to-pay]")
		}

		if tea.Price != "" {
			rawPrice, err := formatTeaBlendPrice(tea.Price)
			if err != nil {
				fmt.Print("Error al parsear precio: ", err.Error())
				return
			}
			tea.RawPrice = rawPrice
		}
		if tea != nil {
			tea.From = "TeaBlends"
			tea.Grams = "40g"
			mutex.Lock() //TODO: Usar channels para evitar mutexLock
			*teaHebras = append(*teaHebras, *tea)
			mutex.Unlock()
		}
	})
}

//TODO: MAs urls, fijarme imgs tema
//func scrape

// Example input: "Desde $8.500,00"
func formatTeaBlendPrice(text string) (int, error) {
	text = strings.Split(text, ",")[0] //Omito decimales
	text = strings.ReplaceAll(text, "Desde", "")
	text = strings.ReplaceAll(text, "$", "")
	text = strings.ReplaceAll(text, " ", "") //Remuevo espacios
	text = strings.ReplaceAll(text, ".", "") //Remuevo puntos

	price, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	return price, nil
}
