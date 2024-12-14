package services

import (
	"github.com/gocolly/colly"
	"google.golang.org/appengine/log"
	"hebras-scrapping/constants"
	"hebras-scrapping/models"
	"sync"
)

type HebrasService struct {
	Utils *HebrasUtils
}

func NewHebrasService() *HebrasService {
	return &HebrasService{
		Utils: NewHebrasUtils(),
	}
}

func (hs *HebrasService) ScrapeHebras(urls []string) (teaHebras []models.HebrasTea) {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			c := colly.NewCollector()
			hs.scrapeTeaBlends(&teaHebras, c, mutex)
			c.OnRequest(func(r *colly.Request) {
				r.Headers.Set("User-Agent", constants.USER_AGENT)
			})
			c.Visit(url)
			c.Wait() //Este wait va a ser individual para cada gorooutine
		}(url)

	}

	wg.Wait()

	return teaHebras
}

// ----------------------------HELPERS--------------------------------

// scrapeTeaBlends Realiza el scrapeo en la pagina de TeaBlends
func (hs *HebrasService) scrapeTeaBlends(teaHebras *[]models.HebrasTea, c *colly.Collector, mutex *sync.Mutex) {
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
			rawPrice, err := hs.Utils.FormatTeaBlendPrice(tea.Price)
			if err != nil {
				log.Infof(nil, "Error al parsear precio: %v", err.Error())
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
