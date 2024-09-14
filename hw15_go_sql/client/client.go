package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kvs4/home_work_basic/hw15_go_sql/internal/repository"
)

const userName = "Ivan"

func Run(url string) {
	// ------------- Users -------------

	body, status := GetUsers(url, "")
	if status != -1 {
		fmt.Println("Result GetUsers body:", body, "Status:", status)
	}

	body, status = GetUsers(url, userName)
	if status != -1 {
		fmt.Println("Result GetUser body:", body, "Status:", status)
	}

	body, status = CreateUser(url)
	if status != -1 {
		fmt.Println("Result CreateUser body:", body, "Status:", status)
	}

	body, status = ChangeEmailUser(url)
	if status != -1 {
		fmt.Println("Result ChangeEmailUser body:", body, "Status:", status)
	}

	// ------------- Orders -------------

	body, status = CreateOrder(url)
	if status != -1 {
		fmt.Println("Result CreateOrder body:", body, "Status:", status)
	}

	body, status = GetUserStatistics(url)
	if status != -1 {
		fmt.Println("Result GetUserStatistics body:", body, "Status:", status)
	}

	body, status = GetUsersOrders(url)
	if status != -1 {
		fmt.Println("Result GetUsersOrders body:", body, "Status:", status)
	}

	// ------------- Products -------------

	body, status = CreateProduct(url)
	if status != -1 {
		fmt.Println("Result CreateProduct body:", body, "Status:", status)
	}

	body, status = GetProducts(url)
	if status != -1 {
		fmt.Println("Result GetProducts body:", body, "Status:", status)
	}

	body, status = GetProduct(url)
	if status != -1 {
		fmt.Println("Result GetProduct body:", body, "Status:", status)
	}

	body, status = ChangePriceProduct(url)
	if status != -1 {
		fmt.Println("Result ChangePriceProduct body:", body, "Status:", status)
	}

	// ------------- Deletion -------------

	body, status = DeleteOrder(url)
	if status != -1 {
		fmt.Println("Result DeleteOrder body:", body, "Status:", status)
	}

	body, status = DeleteUser(url)
	if status != -1 {
		fmt.Println("Result DeleteUser body:", body, "Status:", status)
	}

	body, status = DeleteProduct(url)
	if status != -1 {
		fmt.Println("Result DeleteProduct body:", body, "Status:", status)
	}
}

func GetUsers(url string, name string) (string, int) {
	var urlUsers strings.Builder
	var strParams string

	if name != "" {
		strParams = fmt.Sprint("?name=", name)
	} else {
		strParams = "/"
	}

	urlUsers.WriteString(url)
	urlUsers.WriteString(fmt.Sprint("users", strParams))
	url = urlUsers.String()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("GetUsers. NewRequest error: %w", err))
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("GetUsers. Request error: %w", err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("GetUsers. Body reading error: %w", err))
		return "", -1
	}

	return string(body), resp.StatusCode
}

/*func GetUser(url string) (string, int) {
	var urlUsers strings.Builder

	urlUsers.WriteString(url)
	urlUsers.WriteString("user?name=Ivan")
	url = urlUsers.String()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("GetUser. NewRequest error: %w", err))
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("GetUser. Request error: %w", err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("GetUser. Body reading error: %w", err))
		return "", -1
	}

	return string(body), resp.StatusCode
}*/

func CreateUser(url string) (string, int) {
	name := userName
	email := "ivan@google.com"
	pwd := "789"
	user := repository.CreateUserParams{
		Name:     &name,
		Email:    &email,
		Password: &pwd,
	}
	userByte, err := json.Marshal(user)
	if err != nil {
		fmt.Println(fmt.Errorf("CreateUser. User marshaling error: %w", err))
		os.Exit(1)
	}

	var urlUsers strings.Builder

	urlUsers.WriteString(url)
	urlUsers.WriteString("users/create")
	url = urlUsers.String()

	bodyreq := bytes.NewReader(userByte)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bodyreq)
	if err != nil {
		fmt.Println(fmt.Errorf("CreateUser. NewRequest error: %w", err))
		os.Exit(1)
	}

	respPost, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("CreateUser. Request error: %w", err))
		os.Exit(1)
	}
	defer respPost.Body.Close()

	bodyPost, err := io.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("CreateUser. Body reading error: %w", err))
		return "", -1
	}

	return string(bodyPost), respPost.StatusCode
}

func ChangeEmailUser(url string) (string, int) {
	name := userName
	email := "ivan-petrov@google.com"
	user := repository.ChangeEmailUserParams{
		Email: &email,
		Name:  &name,
	}
	userByte, err := json.Marshal(user)
	if err != nil {
		fmt.Println(fmt.Errorf("ChangeEmailUser. User marshaling error: %w", err))
		os.Exit(1)
	}

	var urlUsers strings.Builder

	urlUsers.WriteString(url)
	urlUsers.WriteString("users/changeemail")
	url = urlUsers.String()

	bodyreq := bytes.NewReader(userByte)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, url, bodyreq)
	if err != nil {
		fmt.Println(fmt.Errorf("ChangeEmailUser. NewRequest error: %w", err))
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("ChangeEmailUser. Request error: %w", err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("ChangeEmailUser. Body reading error: %w", err))
		return "", -1
	}

	return string(body), resp.StatusCode
}

func DeleteUser(url string) (string, int) {
	user := userName
	userByte, err := json.Marshal(user)
	if err != nil {
		fmt.Println(fmt.Errorf("DeleteUser. User marshaling error: %w", err))
		os.Exit(1)
	}

	var urlUsers strings.Builder

	urlUsers.WriteString(url)
	urlUsers.WriteString("users/delete")
	url = urlUsers.String()

	bodyreq := bytes.NewReader(userByte)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, url, bodyreq)
	if err != nil {
		fmt.Println(fmt.Errorf("DeleteUser. NewRequest error: %w", err))
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("DeleteUser. Request error: %w", err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("DeleteUser. Body reading error: %w", err))
		return "", -1
	}

	return string(body), resp.StatusCode
}

// ------------ Orders --------------

func CreateOrder(url string) (string, int) {
	// email := "ivan@google.com"

	// productID := "02b91e73-8ac4-4c7a-a381-f47698719f5f"

	// userID, err := pgtype.parseUUID("6f7ae3ed-3231-4eb2-9086-a7b8c7a278d4")

	orderDate, err := time.Parse(time.DateTime, "2024-09-14 12:01:00")
	if err != nil {
		fmt.Println(fmt.Errorf("CreateOrder. OrderDate parsing error: %w", err))
		os.Exit(1)
	}

	user := struct {
		UserID      string
		OrderDate   time.Time
		TotalAmount string
		ProductID   string
	}{
		UserID:      "6f7ae3ed-3231-4eb2-9086-a7b8c7a278d4",
		OrderDate:   orderDate,
		TotalAmount: "205.5",
		ProductID:   "02b91e73-8ac4-4c7a-a381-f47698719f5f",
	}
	userByte, err := json.Marshal(user)
	if err != nil {
		fmt.Println(fmt.Errorf("CreateOrder. User marshaling error: %w", err))
		os.Exit(1)
	}

	var urlfull strings.Builder

	urlfull.WriteString(url)
	urlfull.WriteString("orders/create")
	url = urlfull.String()

	bodyreq := bytes.NewReader(userByte)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bodyreq)
	if err != nil {
		fmt.Println(fmt.Errorf("CreateOrder. NewRequest error: %w", err))
		os.Exit(1)
	}

	respPost, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("CreateOrder. Request error: %w", err))
		os.Exit(1)
	}
	defer respPost.Body.Close()

	bodyPost, err := io.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("CreateOrder. Body reading error: %w", err))
		return "", -1
	}

	return string(bodyPost), respPost.StatusCode
}

func DeleteOrder(url string) (string, int) {
	return MakeDeleteRequest(url, "orders/delete", "e8794041-958c-4ff8-bad8-88a10c2c94d8", "DeleteOrder")

	/*orderID := "e8794041-958c-4ff8-bad8-88a10c2c94d8"
	orderIDByte, err := json.Marshal(orderID)
	if err != nil {
		fmt.Println(fmt.Errorf("DeleteOrder. User marshaling error: %w", err))
		os.Exit(1)
	}

	var urlfull strings.Builder

	urlfull.WriteString(url)
	urlfull.WriteString("orders/delete")
	url = urlfull.String()

	bodyreq := bytes.NewReader(orderIDByte)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, url, bodyreq)
	if err != nil {
		fmt.Println(fmt.Errorf("DeleteOrder. NewRequest error: %w", err))
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("DeleteOrder. Request error: %w", err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("DeleteOrder. Body reading error: %w", err))
		return "", -1
	}

	return string(body), resp.StatusCode
	*/
}

func GetUserStatistics(url string) (string, int) {
	return MakeGetRequest(url, "orders/userstatistics?name=Jack", "GetUserStatistics")
}

func GetUsersOrders(url string) (string, int) {
	return MakeGetRequest(url, "orders/usersorders", "GetUsersOrders")
}

func MakeGetRequest(url string, path string, funcName string) (string, int) {
	var urlfull strings.Builder

	urlfull.WriteString(url)
	urlfull.WriteString(path)
	url = urlfull.String()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. NewRequest error: %w", funcName, err))
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. Request error: %w", funcName, err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. Body reading error: %w", funcName, err))
		return "", -1
	}

	return string(body), resp.StatusCode
}

// ------------ Products --------------

func GetProducts(url string) (string, int) {
	return MakeGetRequest(url, "products", "GetProducts")
}

func GetProduct(url string) (string, int) {
	return MakeGetRequest(url, "product?name=cheesecake", "GetProduct")
}

func MakePostPutRequest(url string, method string, path string,
	name string, price string, funcName string,
) (string, int) {
	product := struct {
		Name  string
		Price string
	}{
		Name:  name,
		Price: price,
	}
	productByte, err := json.Marshal(product)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. Product marshaling error: %w", funcName, err))
		os.Exit(1)
	}

	var urlfull strings.Builder

	urlfull.WriteString(url)
	urlfull.WriteString(path)
	url = urlfull.String()

	bodyreq := bytes.NewReader(productByte)
	req, err := http.NewRequestWithContext(context.Background(), method, url, bodyreq)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. NewRequest error: %w", funcName, err))
		os.Exit(1)
	}

	respPost, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. Request error: %w", funcName, err))
		os.Exit(1)
	}
	defer respPost.Body.Close()

	bodyPost, err := io.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. Body reading error: %w", funcName, err))
		return "", -1
	}

	return string(bodyPost), respPost.StatusCode
}

func CreateProduct(url string) (string, int) {
	return MakePostPutRequest(url, http.MethodPost, "products/create", "cheesecake", "156.7", "CreateProduct")
}

func ChangePriceProduct(url string) (string, int) {
	return MakePostPutRequest(url, http.MethodPut, "products/changeprice", "cheesecake", "185.1", "ChangePriceProduct")
	/*product := struct {
		Name  string
		Price string
	}{
		Name:  "cheesecake",
		Price: "185.1",
	}
	productByte, err := json.Marshal(product)
	if err != nil {
		fmt.Println(fmt.Errorf("ChangePriceProduct. Product marshaling error: %w", err))
		os.Exit(1)
	}

	var urlfull strings.Builder

	urlfull.WriteString(url)
	urlfull.WriteString("products/changeprice")
	url = urlfull.String()

	bodyreq := bytes.NewReader(productByte)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, url, bodyreq)
	if err != nil {
		fmt.Println(fmt.Errorf("ChangePriceProduct. NewRequest error: %w", err))
		os.Exit(1)
	}

	respPost, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("ChangePriceProduct. Request error: %w", err))
		os.Exit(1)
	}
	defer respPost.Body.Close()

	bodyPost, err := io.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("ChangePriceProduct. Body reading error: %w", err))
		return "", -1
	}

	return string(bodyPost), respPost.StatusCode
	*/
}

func DeleteProduct(url string) (string, int) {
	return MakeDeleteRequest(url, "products/delete", "cheesecake", "DeleteProduct")
}

func MakeDeleteRequest(url string, path string, name string, funcName string) (string, int) {
	// name := "cheesecake"
	nameByte, err := json.Marshal(name)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. Product marshaling error: %w", funcName, err))
		os.Exit(1)
	}

	var urlfull strings.Builder

	urlfull.WriteString(url)
	urlfull.WriteString(path)
	url = urlfull.String()

	bodyreq := bytes.NewReader(nameByte)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, url, bodyreq)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. NewRequest error: %w", funcName, err))
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. Request error: %w", funcName, err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("%s. Body reading error: %w", funcName, err))
		return "", -1
	}

	return string(body), resp.StatusCode
}
