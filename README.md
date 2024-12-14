# hebras-scrapping

Proyecto que realiza Scrapping en P'aginas de t'e en hebras (Argentina), para saber su precio

Tecnolog`ias utilizadas:
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
  },
  {
    "price": "Desde $12.000,00",
    "rawPrice": 12000,
    "from": "TeaBlends",
    "img": "https://static.wixstatic.com/media/eb5541_ca0777be62e94352b5d0d99d6cededce~mv2.jpg/v1/fill/w_147,h_110,al_c,q_80,usm_0.66_1.00_0.01,blur_2,enc_auto/eb5541_ca0777be62e94352b5d0d99d6cededce~mv2.jpg",
    "link": "https://www.teablendsforyou.com.ar/product-page/té-negro-camellia-flowers",
    "grams": "40g"
  },
  {
    "price": "Desde $12.000,00",
    "rawPrice": 12000,
    "from": "TeaBlends",
    "img": "https://static.wixstatic.com/media/eb5541_ffbcf83953c146aa88dfb050a696db4f~mv2.jpg/v1/fill/w_147,h_110,al_c,q_80,usm_0.66_1.00_0.01,blur_2,enc_auto/eb5541_ffbcf83953c146aa88dfb050a696db4f~mv2.jpg",
    "link": "https://www.teablendsforyou.com.ar/product-page/té-negro-cosecha-manual-brote-y-2-hojas-orgánico-argentino-premium",
    "grams": "40g"
  },
  {
    "price": "Desde $14.000,00",
    "rawPrice": 14000,
    "from": "TeaBlends",
    "img": "https://static.wixstatic.com/media/eb5541_26aff2b50daa49519914e0a34a3b4745~mv2.jpg/v1/fill/w_147,h_110,al_c,q_80,usm_0.66_1.00_0.01,blur_2,enc_auto/eb5541_26aff2b50daa49519914e0a34a3b4745~mv2.jpg",
    "link": "https://www.teablendsforyou.com.ar/product-page/té-negro-elegance",
    "grams": "40g"
  }
]
```