package main

import(
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct{  //event
	Name string `json:"what is your name?"`
	Age int `json:"How old are you?"`
}

type MyResponse struct{    //response
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent)(MyResponse, error){   //first accept event than event response 
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main(){
	lambda.Start(HandleLambdaEvent)   //push fn to lambda
}