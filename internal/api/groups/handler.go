package group_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/contacts-api/internal/application"
	"github.com/mahdi-cpp/contacts-api/internal/collections/contact"
	"github.com/mahdi-cpp/contacts-api/internal/collections/group"
	"github.com/mahdi-cpp/contacts-api/internal/help"
	"github.com/mahdi-cpp/iris-tools/mygin"
)

type GroupHandler struct {
	appManager *application.AppManager
}

func New(manager *application.AppManager) *GroupHandler {
	return &GroupHandler{appManager: manager}
}

func (h *GroupHandler) Create(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.SendError(c, "user id invalid", http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var request *group.Group
	err = json.NewDecoder(c.Req.Body).Decode(&request)
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	create, err := accountManager.GroupManager.Create(request)
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("groups is created:", request.Title)

	c.JSON(http.StatusOK, create)
}

func (h *GroupHandler) Read(c *mygin.Context) {

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

	groupID := c.Param("id")
	id, err := uuid.Parse(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	item, err := accountManager.GroupManager.Read(id)
	if err != nil {
		c.JSON(http.StatusNotFound, mygin.H{"error": "Photo not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *GroupHandler) ReadAll(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		fmt.Println("user id invalid")
		return
	}

	page, err := c.GetQueryInt("page")
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
	}
	size, err := c.GetQueryInt("size")
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
	}

	fmt.Println("page:", page)
	fmt.Println("size:", size)

	with := &group.SearchOptions{
		Page: page,
		Size: size,
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Printf("Decode error: %v\n", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
		return
	}

	items, err := accountManager.GroupManager.ReadAll(with)
	if err != nil {
		c.JSON(http.StatusInternalServerError, mygin.H{"error": "failed account Read"})
		return
	}

	fmt.Println("ReadAll count", len(items))

	c.JSON(http.StatusOK, items)
}

func (h *GroupHandler) ReadGroups(c *mygin.Context) {

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		fmt.Println("user id invalid")
		return
	}

	var with *group.SearchOptions
	err := json.NewDecoder(c.Req.Body).Decode(&with)
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Printf("Decode error: %v\n", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
		return
	}

	collections := accountManager.GroupManager.ReadCollections(with)

	fmt.Println("Read Group Collections count", len(collections))

	c.JSON(http.StatusOK, mygin.H{"collections": collections})
}

func (h *GroupHandler) ReadGroupContacts(c *mygin.Context) {

	fmt.Println("ReadGroupContacts")

	userID, ok := help.GetUserID(c)
	if !ok {
		help.AbortWithUserIDInvalid(c)
		fmt.Println("user id invalid")
		return
	}

	var with *contact.SearchOptions
	err := json.NewDecoder(c.Req.Body).Decode(&with)
	if err != nil {
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Printf("Decode error: %v\n", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err.Error()})
		return
	}

	photos, err := accountManager.GroupManager.ReadCollectionContacts(with)

	fmt.Println("Read Group Photos count", len(photos))

	c.JSON(http.StatusOK, photos)
}

func (h *GroupHandler) Delete(c *mygin.Context) {

	fmt.Println("Group Delete")

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

	groupID := c.GetQuery("id")
	fmt.Println("Group Delete groups id: ", groupID)

	id, err := uuid.Parse(groupID)
	if err != nil {
		fmt.Println("Decode error:", err)
		c.JSON(http.StatusBadRequest, mygin.H{"error": err})
		return
	}

	err = accountManager.GroupManager.Delete(id)
	if err != nil {
		fmt.Println("Group delete error:", err)
		c.JSON(http.StatusNotFound, mygin.H{"error": "delete groups " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *GroupHandler) AddContact(c *mygin.Context) {

	fmt.Println("AddPhoto")

	userID, ok := help.GetUserID(c)
	if !ok {
		help.SendError(c, "user id invalid", http.StatusBadRequest)
		return
	}

	accountManager, err := h.appManager.GetAccountManager(userID)
	if err != nil {
		fmt.Println("account error:", err)
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var request *contact.CollectionPhoto
	err = json.NewDecoder(c.Req.Body).Decode(&request)
	if err != nil {
		fmt.Println("Decode error:", err)
		help.SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = accountManager.GroupManager.IsExist(request.ParentID)
	if err != nil {
		fmt.Println("Group isExist error:", err)
		help.SendError(c, err.Error(), http.StatusBadRequest)
	}

	for _, photoID := range request.PhotoIDs {
		err := accountManager.GroupManager.AddContact(request.ParentID, photoID)
		if err != nil {
			fmt.Println("Group addPhoto error:", err)
			help.SendError(c, err.Error(), http.StatusBadRequest)
			return
		}
	}

	fmt.Println("photos add to groups with id: ", request.ParentID.String())

	c.JSON(http.StatusOK, "create")
}
