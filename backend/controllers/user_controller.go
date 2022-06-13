package controllers

import (
	"backend/configs"
	"backend/models"
	"backend/responses"
	"context"
	"fmt"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")

var registered_users *mongo.Collection = configs.GetCollection(configs.DB, "registeredUsers")

// func signup(w http.ResponseWriter , r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	var user users
// 	json.NewDecoder(r.Body).Decode(&user)
// 	var result users
// 	err := userCollection.FindOne(context.TODO(),bson.M{"email":user.Email}).Decode(&result)
// 	if err != nil{
// 		BcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
// 	  user.Password = string(BcryptPassword)
// 	  cur ,err := userCollection.InsertOne(context.TODO(),user)
// 	  if err != nil{
// 	   log.Fatal(err)
// 	  }
// 	  json.NewEncoder(w).Encode(cur.InsertedID)

//    func login(w http.ResponseWriter , r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	var user users
// 	json.NewDecoder(r.Body).Decode(&user)
// 	var result users
// 	err := userCollection.FindOne(context.TODO(),bson.M{"email":user.Email}).Decode(&result)
// 	if err  != nil{
// 	 json.NewEncoder(w).Encode("User not found")
// 	}else{
// 	err = bcrypt.CompareHashAndPassword([]byte(result.Password),[]byte(user.Password))
// 	if err != nil{
// 	 json.NewEncoder(w).Encode("Wrong password")
// 	}else{
// 	json.NewEncoder(w).Encode(result)
// 	}



var validate = validator.New()
func Register(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Register_Users
	
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		fmt.Println("dont",err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		fmt.Println("dont know")
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	fmt.Println(user.Email)
	finding := registered_users.FindOne(ctx,bson.M{"email":user.Email}).Decode(&user)
	fmt.Println(finding,"err")
	
	if finding != nil{
 		BcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
 	  user.Password = string(BcryptPassword)
	  fmt.Println(user.Password)
	}else{
		fmt.Println("user already existed")
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{ Message: "error"})
	}
 

	newUser := models.Register_Users{
		Id:            primitive.NewObjectID(),
		Name:          user.Name,
		Email:         user.Email,
		Password: user.Password,
	}

	result, err := registered_users.InsertOne(ctx, newUser)
	fmt.Println(result, "jhh")
	fmt.Println(newUser,"newuser")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON("success")
		// responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}}
	// )
}
func Login(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Register_Users
	
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		fmt.Println("dont",err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		fmt.Println("dont know")
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	 	var result models.Login_Users
	 	err := registered_users.FindOne(ctx,bson.M{"email":user.Email}).Decode(&result)
	 	if err  != nil{
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "user not found"})
	 	}else{
	 	err = bcrypt.CompareHashAndPassword([]byte(result.Password),[]byte(user.Password))
	 	if err != nil{
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "wrong password"})
	 	}else{
			return c.Status(http.StatusCreated).JSON("success")
		}
	}
}
	


func CreateProduct(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var product models.Product
	fmt.Println(product,"qwerty")
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&product); err != nil {
		fmt.Println("dont",err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&product); validationErr != nil {
		fmt.Println("dont know")
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newUser := models.Product{
		Id:            primitive.NewObjectID(),
		Name:          product.Name,
		Model:         product.Model,
		Specification: product.Specification,
		Price:         product.Price,
	}

	result, err := productCollection.InsertOne(ctx, newUser)
	fmt.Println(result, "jhh")
	fmt.Println(newUser,"newuser")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON("success")
		// responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}}
	// )
}

func GetAProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	productId := c.Params("productId")
	var product models.Product
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(productId)

	err := productCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": product}})
}

func EditAProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	productId := c.Params("productId")
	fmt.Println(productId)
	var product models.Product
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(productId)

	//validate the request body
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&product); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": product.Name, "model": product.Model, "Specification": product.Specification, "Price": product.Price}

	result, err := productCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated product details
	var updatedUser models.Product
	if result.MatchedCount == 1 {
		err := productCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		fmt.Println(err)
	}
fmt.Println(result)
	return c.Status(http.StatusOK).JSON(updatedUser)
}

func DeleteAProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	productId := c.Params("productId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(productId)

	result, err := productCollection.DeleteOne(ctx, bson.M{"id": objId})
	// fmt.Println(result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Product with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Product successfully deleted!"}},
	)
}

func GetAllProducts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.Product
	defer cancel()

	results, err := productCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.Product
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		users = append(users, singleUser)
	}

	return c.Status(http.StatusOK).JSON(users)
	// responses.UserResponse{ Data: &fiber.Map{"data": users}},
	// )
}
