package service

import (
	"Hackthon2023/internal/model"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)

func Client(jsonData []byte, serv model.Auth) (model.Auth, error) {

	// Создаем новый клиент
	client := &fasthttp.Client{}

	// Создаем объект запроса
	req := fasthttp.AcquireRequest()

	// Устанавливаем метод и URL запроса
	req.SetRequestURI("http://example.com")
	req.Header.SetMethod("POST")

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "service-search")

	// Устанавливаем тело запроса
	req.SetBody(jsonData)

	// Создаем объект для ответа
	res := fasthttp.AcquireResponse()

	// Отправляем запрос и получаем ответ
	err := client.Do(req, res)
	if err != nil {
		log.Fatal(err)
	}

	// Выводим статус-код ответа и тело ответа
	//fmt.Printf("Status code: %d\n", res.StatusCode())
	if res.StatusCode() != 200 {
		log.Fatal(res.CloseBodyStream().Error())
	} else {
		//fmt.Printf("Response body: %s\n", res.Body())
		json.Unmarshal(res.Body(), serv)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Освобождаем объекты запроса и ответа
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
	return serv, err
}
