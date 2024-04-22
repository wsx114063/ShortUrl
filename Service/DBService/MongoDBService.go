package MongoDBService

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBSetStruct struct {
	DBName         string
	CollectionName string
}

// 包級別變數用於保存已初始化的 Config
var dbs DBSetStruct

func InitDbSet(dbName string, collectionName string) {
	dbs = DBSetStruct{
		DBName:         dbName,
		CollectionName: collectionName,
	}
}

func ConnectToMongoDB(dbURI string) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(dbURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	return client, nil
}

func InsertOneDocument(client *mongo.Client, document interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dbs.DBName).Collection(dbs.CollectionName)

	// 插入單個文檔
	result, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func InsertManyDocuments(client *mongo.Client, documents []interface{}) (*mongo.InsertManyResult, error) {
	// 選擇要插入的集合
	collection := client.Database(dbs.DBName).Collection(dbs.CollectionName)

	// 插入多個文檔
	result, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateOneDocument(client *mongo.Client, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	// 選擇要更新的集合
	collection := client.Database(dbs.DBName).Collection(dbs.CollectionName)

	// 更新單個文檔
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateManyDocuments(client *mongo.Client, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	// 選擇要更新的集合
	collection := client.Database(dbs.DBName).Collection(dbs.CollectionName)

	// 更新多個文檔
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func IsDataExist(client *mongo.Client, fieldName string, value string) (bool, error) {
	collection := client.Database(dbs.DBName).Collection(dbs.CollectionName)
	filter := bson.M{fieldName: value}
	result := collection.FindOne(context.TODO(), filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		} else {
			log.Fatal(err)
			return false, err
		}
	}
	return true, nil
}

func Query(client *mongo.Client, filterJSON string) (bson.M, error) {
	collection := client.Database(dbs.DBName).Collection(dbs.CollectionName)
	var filter bson.M
	if err := bson.UnmarshalExtJSON([]byte(filterJSON), true, &filter); err != nil {
		panic(err)
	}
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func DisconnectFromMongoDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}
