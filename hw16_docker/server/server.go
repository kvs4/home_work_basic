package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kvs4/home_work_basic/hw16_docker/internal/repository"
	"github.com/kvs4/home_work_basic/hw16_docker/pkg/pgdb"
)

func Run() {
	server := gin.Default()
	v1 := server.Group("/v1")

	v1.GET("/users", GetUsers)
	v1.GET("/user", GetUser)
	v1.POST("/users/create", CreateUser)
	v1.PUT("/users/changeemail", ChangeEmailUser)
	v1.DELETE("/users/delete", DeleteUser)

	v1.GET("/orders/userstatistics", GetUserStatistics)
	v1.GET("/orders/usersorders", GetUsersOrders)
	v1.POST("/orders/create", CreateOrder)
	v1.DELETE("/orders/delete", DeleteOrder)

	v1.GET("/products", GetProducts)
	v1.GET("/product", GetProduct)
	v1.POST("/products/create", CreateProduct)
	v1.PUT("/products/changeprice", ChangePriceProduct)
	v1.DELETE("/products/delete", DeleteProduct)

	server.Run(":8080")
}

func GetUsers(c *gin.Context) {
	ctx := context.Background()
	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.GetUsers(ctx)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("GetUsers. Success: false, msg: %w", err)))
		return
	}
	c.JSONP(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	ctx := context.Background()
	name := c.Request.FormValue("name")
	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.GetUser(ctx, &name)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("GetUser. Success: false, msg: %w", err)))
		return
	}
	c.JSONP(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {
	ctx := context.Background()
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateUser. Error read body: %w", err)))
		return
	}

	var userParams repository.CreateUserParams
	err = json.Unmarshal(body, &userParams)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateUser. Body unmarshaling error: %w", err)))
	}
	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.CreateUser(ctx, userParams)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateUser. User creation error: %w", err)))
	}
	c.JSONP(http.StatusOK, result)
}

func ChangeEmailUser(c *gin.Context) {
	ctx := context.Background()
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("ChangeEmailUser. Error read body: %w", err)))
		return
	}

	var userParams repository.ChangeEmailUserParams
	err = json.Unmarshal(body, &userParams)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("ChangeEmailUser. Body unmarshaling error: %w", err)))
	}
	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.ChangeEmailUser(ctx, userParams)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("ChangeEmailUser. User changing error: %w", err)))
	}
	c.JSONP(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	ctx := context.Background()
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteUser. Error read body: %w", err)))
		return
	}

	var userName string
	err = json.Unmarshal(body, &userName)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteUser. Body unmarshaling error: %w", err)))
	}
	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.DeleteUser(ctx, &userName)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteUser. User deletion error: %w", err)))
	}
	c.JSONP(http.StatusOK, result)
}

// ------------ Products ------------

func GetProducts(c *gin.Context) {
	ctx := context.Background()
	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.GetProducts(ctx)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("GetProducts. Success: false, msg: %w", err)))
		return
	}
	c.JSONP(http.StatusOK, result)
}

func GetProduct(c *gin.Context) {
	ctx := context.Background()
	name := c.Request.FormValue("name")
	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.GetProduct(ctx, &name)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("GetProduct. Success: false, msg: %w", err)))
		return
	}
	c.JSONP(http.StatusOK, result)
}

func CreateProduct(c *gin.Context) {
	ctx := context.Background()
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateProduct. Error read body: %w", err)))
		return
	}

	var params struct {
		Name  string
		Price string
	}
	err = json.Unmarshal(body, &params)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateProduct. Body unmarshaling error: %w", err)))
	}
	var price pgtype.Numeric
	err = price.Scan(params.Price)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateProduct. Price parsing error: %w", err)))
	}

	CreateProductParams := repository.CreateProductParams{
		Name:  &params.Name,
		Price: price,
	}

	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.CreateProduct(ctx, CreateProductParams)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateProduct. Product creation error: %w", err)))
	}
	c.JSONP(http.StatusOK, result)
}

func ChangePriceProduct(c *gin.Context) {
	bodyPrice, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("ChangePriceProduct. Error read body: %w", err)))
		return
	}

	var paramsPrice struct {
		Name  string
		Price string
	}
	err = json.Unmarshal(bodyPrice, &paramsPrice)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("ChangePriceProduct. Body unmarshaling error: %w", err)))
	}
	var price pgtype.Numeric
	err = price.Scan(paramsPrice.Price)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("ChangePriceProduct. Price parsing error: %w", err)))
	}

	ctx := context.Background()
	ChangePriceProductParams := repository.ChangePriceProductParams{
		Name:  &paramsPrice.Name,
		Price: price,
	}
	repo := repository.New(pgdb.DB.Conn())
	resultChgProd, err := repo.ChangePriceProduct(ctx, ChangePriceProductParams)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("ChangePriceProduct. User changing error: %w", err)))
	}
	c.JSONP(http.StatusOK, resultChgProd)
}

func DeleteProduct(c *gin.Context) {
	ctx := context.Background()
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteProduct. Error read body: %w", err)))
		return
	}

	var name string
	err = json.Unmarshal(body, &name)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteProduct. Body unmarshaling error: %w", err)))
	}
	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.DeleteProduct(ctx, &name)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteProduct. User deletion error: %w", err)))
	}
	c.JSONP(http.StatusOK, result)
}

// ------------ Orders --------------

func CreateOrder(c *gin.Context) {
	ctx := context.Background()
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateOrder. Error read body: %w", err)))
		return
	}

	var params struct {
		UserID      string
		OrderDate   time.Time
		TotalAmount string
		ProductID   string
	}
	err = json.Unmarshal(body, &params)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateOrder. Body unmarshaling error: %w", err)))
		return
	}

	var (
		userID      pgtype.UUID
		orderDate   pgtype.Timestamptz
		totalAmount pgtype.Numeric
		productID   pgtype.UUID
	)
	err = userID.Scan(params.UserID)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateOrder. UserID parsing error: %w", err)))
	}
	err = orderDate.Scan(params.OrderDate)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateOrder. OrderDate parsing error: %w", err)))
	}
	err = totalAmount.Scan(params.TotalAmount)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateOrder. TotalAmount parsing error: %w", err)))
	}
	err = productID.Scan(params.UserID)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateOrder. ProductID parsing error: %w", err)))
	}

	if err != nil {
		return
	}

	OrderParams := repository.CreateOrderParams{
		UserID:      userID,
		OrderDate:   orderDate,
		TotalAmount: totalAmount,
	}

	repo := repository.New(pgdb.DB.Conn())

	TxOpt := pgx.TxOptions{
		IsoLevel:   pgx.Serializable,
		AccessMode: pgx.ReadWrite,
	}

	tx, err := pgdb.DB.Conn().BeginTx(ctx, TxOpt)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateOrder. Transaction beginning error: %w", err)))
		return
	}

	result, err := repo.CreateOrder(ctx, OrderParams)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateOrder. Order creation error: %w", err)))
		tx.Rollback(ctx)
		return
	}

	// ------------- OrderProducts -------------

	ParamsOrderProducts := repository.CreateOrderProductsParams{
		OrderID:   result.ID,
		ProductID: productID,
	}

	resultOP, err := repo.CreateOrderProducts(ctx, ParamsOrderProducts)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("CreateOrderProducts. OrderProduct creation error: %w", err)))
		tx.Rollback(ctx)
		return
	}

	tx.Commit(ctx)
	c.JSONP(http.StatusOK, result)
	c.JSONP(http.StatusOK, resultOP)
}

func DeleteOrder(c *gin.Context) {
	ctx := context.Background()
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteOrder. Error read body: %w", err)))
		return
	}

	var OrderID string
	err = json.Unmarshal(body, &OrderID)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteOrder. Body unmarshaling error: %w", err)))
		return
	}

	var orderID pgtype.UUID
	err = orderID.Scan(OrderID)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteOrder. OrderID parsing error: %w", err)))
		return
	}

	repo := repository.New(pgdb.DB.Conn())

	TxOpt := pgx.TxOptions{
		IsoLevel:   pgx.RepeatableRead,
		AccessMode: pgx.ReadWrite,
	}

	tx, err := pgdb.DB.Conn().BeginTx(ctx, TxOpt)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteOrder. Transaction beginning error: %w", err)))
		return
	}

	result, err := repo.DeleteOrder(ctx, orderID)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteOrder. Order deletion error: %w", err)))
		tx.Rollback(ctx)
		return
	}

	// ------------- OrderProducts -------------

	resultOP, err := repo.DeleteOrderProducts(ctx, orderID)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("DeleteOrderProducts. OrderProduct deletion error: %w", err)))
		tx.Rollback(ctx)
		return
	}

	tx.Commit(ctx)
	c.JSONP(http.StatusOK, result)
	c.JSONP(http.StatusOK, resultOP)
}

func GetUsersOrders(c *gin.Context) {
	ctx := context.Background()
	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.GetUsersOrders(ctx)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("GetUsersOrders. Success: false, msg: %w", err)))
		return
	}
	c.JSONP(http.StatusOK, result)
}

func GetUserStatistics(c *gin.Context) {
	ctx := context.Background()
	name := c.Request.FormValue("name")
	repo := repository.New(pgdb.DB.Conn())
	result, err := repo.GetUserStatistics(ctx, &name)
	if err != nil {
		c.JSONP(http.StatusBadRequest, fmt.Sprint(fmt.Errorf("GetUserStatistics. Success: false, msg: %w", err)))
		return
	}
	c.JSONP(http.StatusOK, result)
}
