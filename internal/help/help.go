package help

import (
	"bytes"
	"fmt"
	"io"
	"net/url"

	"github.com/goccy/go-json"
	"github.com/mahdi-cpp/contacts-api/internal/config"
	"github.com/mahdi-cpp/iris-tools/mygin"

	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func SendError(c *mygin.Context, message string, code int) {
	c.JSON(http.StatusBadRequest, mygin.H{"message": message, "code": code})
}

// GetUserID از Gin context، user_id را به صورت string دریافت می‌کند.
func GetUserID(c *mygin.Context) (uuid.UUID, bool) {

	//// این تابع باید بعد از middleware احراز هویت استفاده شود
	//userID, exists := c.Read("X-User-ID")
	//if !exists {
	//	return "", false
	//}
	//
	//userIDStr, ok := userID.(string)
	//if !ok {
	//	return "", false
	//}
	//
	//return userIDStr, true

	//// این تابع باید بعد از middleware احراز هویت استفاده شود
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		return uuid.Nil, false
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return [16]byte{}, false
	}

	return id, true
}

func GetFileSize(filepath string) (int64, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func GetUserId(c *mygin.Context) (int, error) {

	userIDStr := c.GetHeader("userID")
	fmt.Println(userIDStr)

	return strconv.Atoi(userIDStr)
}

func GenerateUUID() (uuid.UUID, error) {
	u7, err2 := uuid.NewV7()
	if err2 != nil {
		return uuid.UUID{}, fmt.Errorf("error generating UUIDv7: %w", err2)
	}
	return u7, nil
}

func StrPtr(str string) *string {
	return &str
}

func BoolPtr(bool bool) *bool {
	return &bool
}

func MakeRequestBody(method, endpoint string, body interface{}) (*http.Response, error) {

	// Build URL with query parameters
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}

	// Marshal body if provided
	var bodyReader io.Reader

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshaling body: %w", err)
	}
	bodyReader = bytes.NewReader(jsonData)

	fmt.Println(u.String())
	fmt.Println("")

	// create request
	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("user_id", config.Mahdi.String())

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	return resp, nil
}

func MakeRequestParam(method, endpoint string, queryParams interface{}) ([]byte, error) {

	// ۱. تبدیل پارامترهای کوئری (struct یا map) به map[string]interface{} با استفاده از JSON
	var paramsMap map[string]interface{}

	if queryParams != nil {
		// Marshal کردن به JSON برای تبدیل struct به بایت
		data, err := json.Marshal(queryParams)
		if err != nil {
			return nil, fmt.Errorf("marshalling query params to json: %w", err)
		}

		// Unmarshal کردن JSON به نقشه (map) برای استخراج پارامترها بر اساس تگ‌های JSON
		if err := json.Unmarshal(data, &paramsMap); err != nil {
			return nil, fmt.Errorf("unmarshalling json to map: %w", err)
		}
	}

	// ۲. ساخت URL با پارامترهای کوئری
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}

	if paramsMap != nil {
		q := u.Query()
		for key, value := range paramsMap {
			// اضافه کردن فقط مقادیر غیر خالی
			valStr := fmt.Sprintf("%v", value)
			if valStr != "" {
				q.Add(key, valStr)
			}
		}
		u.RawQuery = q.Encode()
	}

	fullURL := u.String()
	// این خط را می‌توانید برای دیباگ کردن حذف کنید
	// fmt.Println(fullURL)

	// ۳. ساخت درخواست HTTP
	req, err := http.NewRequest(method, fullURL, nil) // body برای کوئری‌ها nil است
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// ۴. اجرای درخواست
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	// ۵. بررسی وضعیت کد پاسخ
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, resp.Status)
	}

	// ۶. خواندن بدنه پاسخ
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	return respBody, nil
}
