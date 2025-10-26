package contact_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/application"
	"github.com/mahdi-cpp/contacts-api/internal/collections/contact"
	"github.com/mahdi-cpp/contacts-api/internal/help"
	"github.com/mahdi-cpp/iris-tools/mygin"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type PhotoHandler struct {
	appManager *application.AppManager
}

func New(manager *application.AppManager) *PhotoHandler {
	return &PhotoHandler{appManager: manager}
}

func SendError(c *mygin.Context, message string, code int) {
	c.JSON(http.StatusBadRequest, mygin.H{"message": message, "code": code})
}

func (h *PhotoHandler) Create(c *mygin.Context) {

	fmt.Println("1")

	userID, ok := help.GetUserID(c)
	if !ok {
		SendError(c, "user id invalid", http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var request *contact.Contact
	err = json.NewDecoder(c.Req.Body).Decode(&request)
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	create, err := accountManager.ContactManager.Create(request)
	if err != nil {
		fmt.Println(err)
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, create)
}

func (h *PhotoHandler) Read(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	photoID := c.Param("id")
	id, err := uuid.Parse(photoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	item, err := accountManager.ContactManager.Read(id)
	if err != nil {
		c.JSON(http.StatusNotFound, mygin.H{"error": "Photo not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *PhotoHandler) ReadAll(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		fmt.Println("user id invalid")
		return
	}

	var with *contact.SearchOptions
	err := json.NewDecoder(c.Req.Body).Decode(&with)
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Printf("Decode error: %v\n", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
		return
	}

	items, err := accountManager.ContactManager.ReadAll(with)
	if err != nil {
		c.JSON(http.StatusInternalServerError, mygin.H{"error": "failed account Read"})
		return
	}

	fmt.Println("ReadAll count", len(items))

	c.JSON(http.StatusOK, items)
}

func (h *PhotoHandler) Update(c *mygin.Context) {

	fmt.Println("1")

	userID, ok := help.GetUserID(c)
	if !ok {
		SendError(c, "user id invalid", http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var with *contact.UpdateOptions
	err = json.NewDecoder(c.Req.Body).Decode(&with)
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	//fmt.Println("with.Directory", with.Directory)

	err = accountManager.ContactManager.Update(with)
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, "update")
}

func (h *PhotoHandler) Delete(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	id := c.Param("contactId")
	photoId, err := uuid.Parse(id)
	if err != nil {
		return
	}

	err = accountManager.ContactManager.Delete(photoId)
	if err != nil {
		c.JSON(http.StatusNotFound, mygin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, "")
}
