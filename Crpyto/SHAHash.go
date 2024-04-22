package Crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func NewHash(length int32) (result string) {
	// 設置隨機種子
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成一個隨機數
	randomNumber := randomGenerator.Intn(1000)

	// 將隨機數轉換為字節切片
	data := []byte(fmt.Sprintf("%d", randomNumber))

	// 創建一個 SHA-256 哈希實例
	hash := sha256.New()

	// 添加數據到哈希實例中
	hash.Write(data)

	// 計算哈希值
	hashedData := hash.Sum(nil)

	// 將哈希值轉換為十六進制字串
	hashedString := hex.EncodeToString(hashedData)

	// 只取前7碼
	hashedString = hashedString[:length]
	return hashedString
}
