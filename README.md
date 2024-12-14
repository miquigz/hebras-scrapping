# hebras-scrapping

Proyecto que realiza Scrapping en P'aginas de t'e en hebras (Argentina), para saber su precio

Tecnologias utilizadas:
Gorilla Mux - Go Colly (Scrapping)

MVC Pattern

### GET /api/1/scrape/hebras

Response example:
```json
[
  {
    "price": "Desde $10.000,00",
    "rawPrice": 10000,
    "from": "TeaBlends",
    "img": "https://static.wixstatic.com/media/eb5541_2340a7741cbb412a99f292c43766b8e2~mv2.png/v1/fill/w_49,h_37,al_c,q_85,usm_0.66_1.00_0.01,blur_2,enc_auto/eb5541_2340a7741cbb412a99f292c43766b8e2~mv2.png",
    "link": "https://www.teablendsforyou.com.ar/product-page/té-negro-repostero",
    "grams": "40g"
  }
]
```