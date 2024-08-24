package main

import "fmt"

type Item struct {
	Id    int
	name  string
	price int
}

var items = []Item{
	{1, "Roti", 2000},
	{2, "Selai", 3000},
	{3, "Kacang", 4000},
}

type ItemNotFoundException struct {
	Id int
}

func (i *ItemNotFoundException) Error() string {
	return fmt.Sprintf("Item ID : %d not found", i.Id)
}

func getItem(id int) (Item, error) {
	for _, v := range items {
		if v.Id == id {
			return v, nil
		}
	}
	return Item{}, &ItemNotFoundException{Id: id}
}

func main() {
	result, err := getItem(7)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
