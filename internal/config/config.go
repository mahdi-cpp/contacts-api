package config

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	RootDir = "/app/iris/com.iris.contacts/"
	users   = "users"
)

func GetPath(file string) string {
	return filepath.Join(RootDir, file)
}

func GetUserPath(userID string) string {
	pp := filepath.Join(RootDir, users, userID)
	fmt.Println(pp)
	return pp
}

func GetUserMetadataPath(id string) string {
	pp := filepath.Join(RootDir, users, id, "metadata")
	fmt.Println(pp)
	return pp
}

var (
	Mahdi    uuid.UUID
	Parsa    uuid.UUID
	Ali      uuid.UUID
	Maryam   uuid.UUID
	Parastoo uuid.UUID
	Fharhad  uuid.UUID
	Nader    uuid.UUID
)
var (
	Digikala uuid.UUID
	Varzesh3 uuid.UUID
)
var (
	ChatID1   uuid.UUID
	ChatID2   uuid.UUID
	ChatID3   uuid.UUID
	MessageID uuid.UUID
)

func initUsers() {
	var err error

	Mahdi, err = uuid.Parse("0199b306-d156-7c6b-b122-1b309599fb82")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

	Parsa, err = uuid.Parse("0199edc1-469b-7de7-bf36-c6cbf299b874")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

	Ali, err = uuid.Parse("0199ee86-f4ef-722a-8355-f860ce9513a3")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

	Maryam, err = uuid.Parse("0199ee39-47f9-7673-88aa-e311f3ad0575")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

	Parastoo, err = uuid.Parse("0199ee92-2bed-7a9b-921a-e897e3cbc42c")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

	Fharhad, err = uuid.Parse("0199fd8b-a0fe-7dbd-bc76-a0f4638b24d2")
	if err != nil {
		log.Fatalf("failed to parse Fharhad: %v", err)
	}

	Nader, err = uuid.Parse("0199fd92-bcdc-7328-b719-8fd2251954ad")
	if err != nil {
		log.Fatalf("failed to parse Fharhad: %v", err)
	}
}

func initChats() {
	var err error

	Digikala, err = uuid.Parse("018f3a8b-1b32-7292-b2d9-1237a7b8c8d2")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}
	Varzesh3, err = uuid.Parse("018f3a8b-1b32-7293-c1d4-8765f5d1e2f3")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

}

// The Init function is called before main() and is ideal for initialization
func Init() {
	var err error

	initUsers()

	ChatID1, err = uuid.Parse("018f3a8b-1b32-7295-a2c7-87654b4d4567")
	if err != nil {
		log.Fatalf("failed to parse ChatID1: %v", err)
	}

	ChatID2, err = uuid.Parse("01992ecc-bb15-7ba6-b340-cc0366eee30a")
	if err != nil {
		log.Fatalf("failed to parse ChatID1: %v", err)
	}

	ChatID3, err = uuid.Parse("01992530-c81c-7d64-ac4f-a4f29678cfc0")
	if err != nil {
		log.Fatalf("failed to parse ChatID1: %v", err)
	}

	MessageID, err = uuid.Parse("01991bc4-faad-7b70-aedc-f20ea4146898")
	if err != nil {
		log.Fatalf("failed to parse MessageID: %v", err)
	}
}
