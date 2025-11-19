/*
Создайте слайс строк, содержащий названия городов. Реализуйте функции для добавления нового города, удаления города по имени и поиска города в списке. Продемонстрируйте работу этих функций на примере.
*/

package main

import "fmt"

func main() {

	towns := []string{"Москва", "Екатеринбург", "Новосибирск"}
	//Добавляем город
	towns = addTown(towns, "Самара")
	fmt.Println(towns)
	//Удаляем город
	towns = deleteTown(towns, "Москва")
	fmt.Println(towns)
	//Ищем город
	searchTown(towns, "Новосибирск")
}

func addTown(s []string, town string) []string {
	s = append(s, town)
	return s
}
func deleteTown(s []string, town string) []string {
	result := make([]string, 0, len(s))
	for _,v := range s {
		if v != town {
			result = append(result, v)
		}
	}
	return result
}
func searchTown(s []string, town string) {
	for i,v := range s {
		if v == town {
			fmt.Printf("Найден город: %v, индекс элемента: %v\n", town, i)
			break
		}
	}
}