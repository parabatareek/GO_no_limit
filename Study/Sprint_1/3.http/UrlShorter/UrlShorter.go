package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	//адрес сервиса
	endpoint := "http://localhost:8080"
	// контейнер данных для запроса
	data := url.Values{}

	// приглашение на ввод URL для сокращения
	fmt.Println("введите длинный URL")

	// потоковое чтение консоли
	reader := bufio.NewReader(os.Stdin)

	// чтение данных введенных через консоль
	long, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	long = strings.TrimSuffix(long, "\n")
	// заполняем контейнер данными
	data.Set("url", long)

	// конструирует HTTP-клиент
	client := &http.Client{}

	//конструируем запрос
	// запрос методом POST должен, кроме заголовков, содержать тело
	// тело должно быть источником потокового чтения io.Reader
	// в большинстве случаев отлично подходит bytes.Buffer
	request, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// в заголовках запроса сообщаем,
	// что данные кодированы стандартной URL-схемой
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	// отправляем запрос и получаем ответ
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// печатаем код ответа
	fmt.Println("статус-код", response.Status)
	defer response.Body.Close()

	// читаем поток из тела ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// печатаем поток из тела ответа
	fmt.Println(string(body))
}
