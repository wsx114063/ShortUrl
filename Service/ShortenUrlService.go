package ShortenUrlService

import (
	crpyto "Project/StartGoLang/Crpyto"
	mongodb "Project/StartGoLang/Service/DBService"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type FormData struct {
	LongUrl string
}

type ShortenUrl struct {
	LongUrl  string `bson:"longUrl"`
	ShortUrl string `bson:"shortUrl"`
}

// 短網址跳轉
func ShortUrl(c *gin.Context) {
	param := c.Param("shortUrl")
	var result ShortenUrl
	mongoHost := os.Getenv("MONGODB_HOST")
	mongoPort := os.Getenv("MONGODB_PORT")
	connectionString := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	fmt.Println(connectionString)
	client, err := mongodb.ConnectToMongoDB(connectionString)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer mongodb.DisconnectFromMongoDB(client)
	mongodb.InitDbSet("Project", "ShortenUrl")

	// 查詢Json
	filterJson := fmt.Sprintf("{ \"shortUrl\": \"%s\" }", param)
	fmt.Println(filterJson)
	data, err := mongodb.Query(client, filterJson)
	if err != nil {
		log.Fatal(err)
	}

	bsonData, err := bson.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	if err := bson.Unmarshal(bsonData, &result); err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusMovedPermanently, result.LongUrl)
}

// 縮短網址
func Shorten(c *gin.Context) {
	var formData FormData
	if err := c.Bind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// DB建構
	mongoHost := os.Getenv("MONGODB_HOST")
	mongoPort := os.Getenv("MONGODB_PORT")
	connectionString := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	fmt.Println(connectionString)
	client, err := mongodb.ConnectToMongoDB(connectionString)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer mongodb.DisconnectFromMongoDB(client)
	mongodb.InitDbSet("Project", "ShortenUrl")

	// 建立雜湊後的url
	newHash := crpyto.NewHash(7)
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")
	newUrl := fmt.Sprintf("http://%s:%s/shortUrl/%s", apiHost, apiPort, newHash)
	dataExist, err := mongodb.IsDataExist(client, "ShortUrl", newUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dataExist fail"})
	}

	// 確保雜湊唯一值
	for {
		if !dataExist {
			break
		} else {
			dataExist, err = mongodb.IsDataExist(client, "ShortUrl", newHash)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "dataExist fail"})
			}
		}
	}

	// 創建要插入的文檔
	document := ShortenUrl{
		LongUrl:  formData.LongUrl,
		ShortUrl: newHash,
	}

	// 插入文檔
	if _, err := mongodb.InsertOneDocument(client, document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shortUrl": newUrl})
}
