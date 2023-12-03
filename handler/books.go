package handler

import (
	"encoding/json"
	"net/http"
	"online-library/controller"
	"online-library/model"
	"online-library/utils"
	"strconv"
)

// GetList : get list data of books. only logged in user can see picked up schedule of books
func GetList(w http.ResponseWriter, r *http.Request) {
	var (
		parGenre = r.URL.Query().Get("genre")
		parLimit = r.URL.Query().Get("limit")
		parPage  = r.URL.Query().Get("page")

		limit, page, statusCode int
		reqGetBooks             *model.ReqGetBooks
		retData                 []*model.RetData
		err                     error
	)

	// set up parameter for library api
	limit, err = strconv.Atoi(parLimit)
	page, err = strconv.Atoi(parPage)
	if page == 0 {
		page = 1
	}
	reqGetBooks = &model.ReqGetBooks{
		Subjects: parGenre,
		Limit:    limit,
		Offset:   (page - 1) * limit,
	}

	// hit get books data library
	retData, statusCode, err = controller.GetBooksData(reqGetBooks)
	if err != nil {
		utils.ReturnResponse(w, statusCode, err.Error(), nil)
		return
	}

	utils.ReturnResponse(w, statusCode, "", retData)
}

// SetPickup : set pickup schedule time for a book
func SetPickup(w http.ResponseWriter, r *http.Request) {
	var (
		reqPickup *model.ReqPickup
		err       error
	)

	// parse json from request body
	err = json.NewDecoder(r.Body).Decode(&reqPickup)
	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, utils.ErrorFailedReadRequest(), nil)
		return
	}

	// set pickup time for a book
	err = controller.SetBookPickup(reqPickup)
	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ReturnResponse(w, http.StatusOK, "Successfully save pickup time", nil)
}
