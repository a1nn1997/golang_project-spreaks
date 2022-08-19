package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/a1nn1997/golang-jwt/database"
	"github.com/a1nn1997/golang-jwt/models"
	utils "github.com/a1nn1997/golang-jwt/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection =database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string)(string){
	bytes,err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!=nil{
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword),[]byte(userPassword))
	check := true
	msg := ""
	if err!=nil{
		msg= fmt.Sprintf("incorrect password")
		check=false
	}
	return check,msg
}


func SignUp() gin.HandlerFunc{

	return func(c *gin.Context){
		ctx,cancel :=context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		if err :=c.BindJSON(&user); err !=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email count"})
		}
		if count > 0{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email already exist"})
		}

		password :=HashPassword(*user.Password)
		user.Password = &password
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone count"})
		}
		if count > 0{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this phone number already exist"})
		}
		user.Created_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while updating time"})
		}
		user.Updated_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while updating time"})
		}
		user.ID=primitive.NewObjectID()
		user.User_id=user.ID.Hex()
		token, refreshToken, _ := utils.GenenrateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil{
			msg :=fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}
func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx,cancel :=context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User
		if err := c.BindJSON(&user); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		err:= userCollection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"email doesn't exist or password is not correct "})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true{
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		if foundUser.Email == nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"user not found"})
		}
		token, refreshToken, err := utils.GenenrateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"token not created"})
			return
		}
		utils.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id":foundUser.User_id}).Decode(&foundUser)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK, foundUser)
	}
}


func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context){
		if err := utils.CheckUserType(c, "ADMIN"); err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage,err := strconv.Atoi(c.Query("recordPerPage"))
		if err!=nil || recordPerPage<1{
			recordPerPage = 10
		}
		page,err1 := strconv.Atoi(c.Query("page"))
		if err1!=nil || page<1{
			page = 1
		}
	startIndex := (page - 1) * recordPerPage
	startIndex, err = strconv.Atoi("startIndex")

	matchStage := bson.D{{"$match",bson.D{{}}}}
	groupStage := bson.D{{"$group", 
		bson.D{{"_id" ,bson.D{{"_id","null"}}},
		{"total_count" ,bson.D{{"$sum",1}}}, 
		{"data", bson.D{{"$push","$$ROOT"}}}}}}
	projectStage := bson.D{{ "$project", 
		bson.D{{"_id",0},{"total_count",1},
{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
	}}}
	result,err :=userCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage })
	defer cancel()
	if err !=nil{
			
		//c.JSON(http.StatusInternalServerError, gin.H{"error":"error occurs while testing items"})
		c.JSON(http.StatusInternalServerError, gin.H{"error":err})
			return
		}
		var allusers []bson.M
			if err = result.All(ctx, &allusers); err!=nil{
				log.Fatal(err)
			}
		
		c.JSON(http.StatusOK, allusers[0])
	}
}

func GetUserById() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		if err :=utils.MathchUserTypeToUid(c,userId); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
