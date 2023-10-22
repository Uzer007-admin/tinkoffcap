# tinkoffcap
 Tinkoff Hackaton 2023
Мы не успели дописать проект, увы.
https://tinkoff.web-gen.ru/ это ссылка на веб сайт.

https://tinkoff.web-gen.ru/api/service-auth/  api пользовательского микросервиса
https://tinkoff.web-gen.ru/api/service-main/  api основного сервиса
https://tinkoff.web-gen.ru/api/service-search/  api поискового микросервиса

https://tinkoff.web-gen.ru/db/

https://www.figma.com/file/96q6n1fI17wxcSapSVti84/Untitled?type=design&node-id=0%3A1&mode=design&t=4I3KJsu0U5xmCWIL-1
Сервис написан на Go
    Go:
        В Go применялась чистая архитектура, но в основном сервисе к утру, пришлось упростить подход.
        Под http сервер был выбран fasthttp, из-за скорости.
    Reindexer:
        Быстрая in-memory DB, которая дала возможности к полнотекстовому поиску, быстрому отклику. 