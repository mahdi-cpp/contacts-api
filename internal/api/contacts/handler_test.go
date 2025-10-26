package contact_handler

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/mahdi-cpp/contacts-api/internal/collections/contact"
	"github.com/mahdi-cpp/contacts-api/internal/help"
)

func TestAssetHandler_Create(t *testing.T) {

	currentURL := baseURL + "/contacts/api/contacts"

	bo := &contact.Contact{
		FirstName:   "Fharhad",
		LastName:    "Abdolmaleki",
		OriginalURL: "",
		//ThumbnailURL: "/app/iris/services/accounts/assets/thumbnails/cat",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	resp, err := help.MakeRequestBody("POST", currentURL, bo)
	if err != nil {
		t.Fatalf("create request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response: %v", err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		var r Error
		if err := json.Unmarshal(respBody, &r); err != nil {
			t.Fatalf("unmarshaling response: %v", err)
		}
		t.Fatalf("error %s", r.Message)
	}

	var a contact.Contact
	if err := json.Unmarshal(respBody, &a); err != nil {
		t.Fatalf("unmarshaling response: %v", err)
	}

	fmt.Println("new photo id: ", a.ID)
}

const baseURL = "http://localhost:50153"

func TestAssetHandler_Read(t *testing.T) {

	// 💡 مسیر را بدون اسلش انتهایی که در روتر ثبت نشده است، تنظیم کنید
	currentURL := baseURL + "/api/contacts"

	with := &contact.SearchOptions{
		Sort:      "id",
		SortOrder: "desc",
	}

	_, err := help.MakeRequestParam("GET", currentURL, with)
	if err != nil {
		t.Fatalf("read request failed: %v", err)
	}
}
