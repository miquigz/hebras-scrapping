package services

import (
	"github.com/gocolly/colly"
	"hebras-scrapping/models"
)

func ScrapeHebras(teaBlendsUrl bool) []models.HebrasTea {
	c := colly.NewCollector()

	var teaHebras []models.HebrasTea
	if teaBlendsUrl {
		c.OnHTML("li[data-hook=product-list-grid-item]", func(e *colly.HTMLElement) {
			scrapeTeaBlends(&teaHebras, *e)
		})
	}

	c.OnRequest(func(r *colly.Request) {
		// Set the User-Agent header
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (HTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")
	})
	//Estos son la mayoria lata de 40g:
	c.Visit("https://www.teablendsforyou.com.ar/te-negro?sort=price_ascending")
	c.Wait()
	return teaHebras
}

// ----------------------------HELPERS--------------------------------

// *esto mismo para el link: teablendsforyou.com.ar
func scrapeTeaBlends(teaHebras *[]models.HebrasTea, e colly.HTMLElement) {
	tea := new(models.HebrasTea)
	if e.ChildText("h3[data-hook=product-item-name]") != "" {
		tea.Name = e.ChildText("h3[data-hook=product-item-name]")
	}
	if e.ChildText("span[data-hook=price-range-from]") != "" {
		tea.Price = e.ChildText("span[data-hook=price-range-from]")
	} else {
		tea.Price = e.ChildText("span[data-hook=product-item-price-to-pay]")
	}
	if tea != nil {
		*teaHebras = append(*teaHebras, *tea)
	}
}
