package db

// import (
// 	"kiokumushi-api/app/config"
// )

// var db *dynamodb.DynamoDB

// func Init() {
// 	c := config.GetConfig()
// 	db = dynamodb.New(session.New(&aws.Config{
// 		Region:      aws.String(c.GetString("db.region")),
// 		Credentials: credentials.NewEnvCredentials(),
// 		Endpoint:    aws.String(c.GetString("db.endpoint")),
// 		DisableSSL:  aws.Bool(c.GetBool("db.disable_ssl")),
// 	}))
// }

// func GetDB() *dynamodb.DynamoDB {
// 	return db
// }
