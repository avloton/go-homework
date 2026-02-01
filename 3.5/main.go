// Цель: Написать тесты для простого HTTP-сервера
// Задание:
// Напишите тесты для этого сервера, проверяющие:
//    - Корректность кода ответа (200 OK).
//    - Корректность заголовка Content-Type (application/json).
//    - Корректность тела ответа (должен быть JSON с полем message).

package main

import (
    "encoding/json"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "hello, world!"})
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}