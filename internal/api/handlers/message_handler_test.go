package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"
)

// TestAPIMessages simulates the curl command
func TestAPIMessages(t *testing.T) {

	// آدرس URL و پورت سرور
	serverURL := "http://localhost:50151"
	apiPath := "/api/messages"

	// ساخت یک URL با استفاده از پکیج url
	// این کار برای اطمینان از فرمت صحیح پارامترها بسیار مهم است
	u, err := url.Parse(serverURL + apiPath)
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}

	// اضافه کردن پارامترهای کوئری
	params := u.Query()
	params.Set("chatId", "aliali")
	params.Set("limit", "23")
	u.RawQuery = params.Encode()

	// ایجاد یک درخواست GET
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// بررسی وضعیت پاسخ (مثلاً کد 200)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}

	// خواندن پاسخ از Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// چاپ پاسخ
	fmt.Printf("status: %s\n", resp.Status)
	fmt.Printf("Body: %s\n", body)
}
