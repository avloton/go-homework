// Изучить документацию для API сервиса coincap.io (https://pro.coincap.io/api-docs)
// Создайте аккуаунт, создайте API ключ и на ЯП Go сделайте GET запрос: /assets/. Вывод включает много полей. Выведите на экран ограниченным списком:
// [{
//     name: Bitcoin,
//     priceUsd: 57435.2862859343831301,
// },
// {
//     name: Ethereum,
//     priceUsd: 3686.1846963999934439,

// }
// ...
// ]

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://rest.coincap.io/v3/assets?ids=bitcoin,ethereum", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	request.Header.Set("Authorization", "Bearer 9cfed58054d22a551d3c42389578be7d7f46e20de8221c581b07b1823c61cb1b")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var result map[string]any
	json.Unmarshal(body, &result)
	
	if data, ok := result["data"].([]any); ok {
		for _, v := range data {
			var name, price string
			if r, ok := v.(map[string]any); ok {
				if str, ok := r["name"].(string); ok {
					name = fmt.Sprintf("name: %s", str)
				}			
				if str, ok := r["priceUsd"].(string); ok {
					price = fmt.Sprintf("priceUsd: %s", str)
				}		
			}
			fmt.Println(name)
			fmt.Println(price)
		}
	}
}