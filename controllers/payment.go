package controllers

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"payment-service/config"
	"payment-service/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreatePaymentSession 创建Apple Pay会话
func CreatePaymentSession(c *gin.Context) {
	var input struct {
		ValidationURL string `json:"validationURL"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建Apple Pay会话
	session, err := validateMerchant(input.ValidationURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment session"})
		return
	}

	c.JSON(http.StatusOK, session)
}

// ProcessApplePayment 处理Apple Pay支付
func ProcessApplePayment(c *gin.Context) {
	var input models.ApplePayRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	// 生成订单ID
	if input.OrderID == "" {
		input.OrderID = generateOrderID()
	}

	// 创建支付记录
	payment := models.Payment{
		UserID:        userID.(uint),
		OrderID:       input.OrderID,
		Amount:        input.Amount,
		Currency:      "CNY",
		Status:        "pending",
		PaymentMethod: "apple_pay",
		ApplePayToken: serializeToken(input.Token),
		TransactionID: input.Token.TransactionIdentifier,
		ProductType:   input.ProductType,
		ProductID:     input.ProductID,
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	// 处理Apple Pay token
	success, err := processApplePayToken(input.Token)
	if err != nil || !success {
		payment.Status = "failed"
		config.DB.Save(&payment)
		c.JSON(http.StatusPaymentRequired, gin.H{"error": "Payment failed"})
		return
	}

	// 更新支付状态
	payment.Status = "completed"
	config.DB.Save(&payment)

	// 更新用户会员状态（如果是会员购买）
	if input.ProductType == "membership" {
		updateUserMembership(userID.(uint))
	}

	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"orderId":       payment.OrderID,
		"transactionId": payment.TransactionID,
		"status":        payment.Status,
	})
}

// GetPaymentStatus 获取支付状态
func GetPaymentStatus(c *gin.Context) {
	orderID := c.Param("orderId")

	var payment models.Payment
	if err := config.DB.Where("order_id = ?", orderID).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// RefundPayment 退款
func RefundPayment(c *gin.Context) {
	orderID := c.Param("orderId")

	var payment models.Payment
	if err := config.DB.Where("order_id = ?", orderID).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	if payment.Status != "completed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can only refund completed payments"})
		return
	}

	// 处理退款逻辑
	payment.Status = "refunded"
	config.DB.Save(&payment)

	c.JSON(http.StatusOK, gin.H{"success": true, "status": "refunded"})
}

// 辅助函数

func validateMerchant(validationURL string) (map[string]interface{}, error) {
	merchantID := os.Getenv("APPLE_MERCHANT_ID")
	if merchantID == "" {
		merchantID = "merchant.com.zhix.club"
	}

	certPath := os.Getenv("APPLE_PAY_CERT_PATH")
	keyPath := os.Getenv("APPLE_PAY_KEY_PATH")

	if certPath == "" || keyPath == "" {
		// 开发环境返回模拟数据
		return map[string]interface{}{
			"epochTimestamp":    time.Now().Unix(),
			"expiresAt":         time.Now().Add(5 * time.Minute).Unix(),
			"merchantSessionIdentifier": generateSessionID(),
			"nonce":             generateNonce(),
			"merchantIdentifier": merchantID,
			"domainName":        "localhost",
			"displayName":       "极志社区",
			"signature":         "mock_signature",
		}, nil
	}

	// 生产环境：使用证书调用Apple服务器
	payload := map[string]interface{}{
		"merchantIdentifier": merchantID,
		"displayName":        "极志社区",
		"initiative":         "web",
		"initiativeContext":  "zhix.club",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", validationURL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var session map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
		return nil, err
	}
	return session, nil
}

func processApplePayToken(token models.ApplePayToken) (bool, error) {
	// 在生产环境中，需要：
	// 1. 解密token.PaymentData
	// 2. 验证签名
	// 3. 调用支付网关处理支付
	// 4. 返回处理结果

	// 开发环境：模拟成功
	return true, nil
}

func updateUserMembership(userID uint) {
	// 直接更新数据库中的用户会员状态
	type User struct {
		ID        uint `gorm:"primarykey"`
		IsPremium bool
	}
	
	if err := config.DB.Model(&User{}).Where("id = ?", userID).Update("is_premium", true).Error; err != nil {
		fmt.Printf("Failed to update membership for user %d: %v\n", userID, err)
		return
	}
	
	fmt.Printf("Successfully updated membership for user %d\n", userID)
}

func generateOrderID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return "order_" + hex.EncodeToString(b)
}

func generateSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func generateNonce() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func serializeToken(token models.ApplePayToken) string {
	data, _ := json.Marshal(token)
	return string(data)
}
