package data

// import (
// 	"context"
// 	"grpc-layout/configs"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// // 初始化mongo连接
// func ConnectToMongoDB() (*mongo.Database, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(configs.Cfg.MongoDB.Timeout)*time.Second)
// 	defer cancel()
// 	uri := configs.MongoDBURI()
// 	opt := options.Client().ApplyURI(uri)
// 	client, err := mongo.Connect(ctx, opt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = client.Ping(context.Background(), readpref.Primary())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client.Database(configs.Cfg.MongoDB.DBName), nil
// }

// func GetCasbinEnforcer(confPath string, databaseName string) (*casbin.Enforcer, error) {
// 	uri := configs.MongoDBURI()
// 	mongoClientOption := options.Client().ApplyURI(uri)
// 	client, err := mongodbadapter.NewAdapterWithClientOption(mongoClientOption, databaseName)
// 	if err != nil {
// 		log.Info().Str("uri", uri).Msgf("init mongo cli err:%v", err)
// 		return nil, err
// 	}
// 	e, err := casbin.NewEnforcer(confPath, client)
// 	if err != nil {
// 		log.Info().Str("path", confPath).Msgf("请检查casbin 模型文件路径是否正确,err:%v", err)
// 		return nil, err
// 	}
// 	// b, err := e.AddPolicy("alice", "data2", "read")
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// fmt.Printf("b: %v\n", b)
// 	// b, err = e.Enforce("alice1", "data2", "read")
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// fmt.Printf("b: %v\n", b)
// 	// e.SavePolicy()
// 	return e, nil
// }
