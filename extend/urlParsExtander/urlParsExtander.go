package urlparsextander

import "net/url"

// у структуры url.URL в стандартной библиотеке net/url не имплементированы методы получения всех ключей и значений

type CustomUrlMethods struct {
	url.URL
}

// Получение всех значений из query параметров в виде двумерного массива
func (c *CustomUrlMethods) GetAllValues() [][]string {
	values := [][]string{}
	for _, value := range c.URL.Query() {
		values = append(values, value)
	}
	return values
}

// Получение всех ключей query параметров
func (c *CustomUrlMethods) GetAllKeys() []string {
	keys := []string{}
	for key := range c.URL.Query() {
		keys = append(keys, key)
	}
	return keys
}

/* Пример использования
func GetAllKeysAndValue(res http.ResponseWriter, req *http.Request) {
	customMethods := CustomUrlMethods{
		*req.URL,
	}

	keys := customMethods.GetAllKeys()
	values := customMethods.GetAllValues()

	fmt.Println(keys)
	fmt.Println(values)
}
*/
