# veil-api-client-go

veil-api-client-go - это клиентская библиотека на Go для доступа к [ECP VeiL REST API](https://veil.mashtab.org/docs/latest/api/).

## Использование

```
require (
    github.com/jsc-masshtab/veil-api-client-go
)
```

Токен интеграции можно сгенерировать в [интерфейсе ECP VeiL](https://veil.mashtab.org/docs/latest/base/operator_guide/security/users/#_11).

```
// For example 
apiUrl := "http://192.168.11.105"
token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjo1NiwidXNlcm5hbWUiOiJidXIiLCJleHAiOjE5NTU0Mjc5OTEsInNzbyI6ZmFsc2UsIm9yaWdfaWF0IjoxNjQwOTMxOTkxfQ.BCPJi1hE_uvlv_sCjLYwGGq2qKJU8dbR9UUC5Cy79AA"
client := NewClient(apiUrl, token, false)
response, _, err := client.DataPool.List()
```

Некоторые операции могут быть выполнены синхронно и асинхронно
```
// Последний аргумент asynced булевый
// true - для ожидания завершения операции
// false - если не нужно ждать завершения
client := NewClient(apiUrl, token, false)
vdisk, _, err := client.Vdisk.Create(NameGenerator("vdisk"), false, firstDp.Id, 0.1, true)
```

## Тесты

Запуск отдельных тестов:
```bash
$ go test -v github.com/jsc-masshtab/veil-api-client-go -run Test_DomainList
```

Для удобства можно использовать "VEIL_API_TOKEN" и "VEIL_API_URL" переменные окружения и передавать пустые строки в NewClient.
```sh
export VEIL_API_URL="http://192.168.11.105"
export VEIL_API_TOKEN="eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjo1NiwidXNlcm5hbWUiOiJidXIiLCJleHAiOjE5NTU0Mjc5OTEsInNzbyI6ZmFsc2UsIm9yaWdfaWF0IjoxNjQwOTMxOTkxfQ.BCPJi1hE_uvlv_sCjLYwGGq2qKJU8dbR9UUC5Cy79AA"
```