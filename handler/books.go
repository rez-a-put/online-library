package handler

import (
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
		respMsg                 string
		err                     error
	)

	// simple genre validation
	if parGenre == "" {
		statusCode = http.StatusBadRequest
		respMsg = utils.ErrorRequired("genre of books")

		utils.ReturnResponse(w, statusCode, respMsg, nil)
		return
	}

	// set up parameter for library api
	limit, err = strconv.Atoi(parLimit)
	page, err = strconv.Atoi(parPage)
	reqGetBooks = &model.ReqGetBooks{
		Subjects: parGenre,
		Limit:    limit,
		Offset:   (page - 1) * limit,
	}

	// hit get books data library
	retData, err = controller.GetBooksData(reqGetBooks)
	if err != nil {
		statusCode = http.StatusBadRequest
		respMsg = err.Error()
	}

	utils.ReturnResponse(w, http.StatusOK, respMsg, retData)
}

// SetPickup : set pickup schedule time for a book
func SetPickup(w http.ResponseWriter, r *http.Request) {
	var err error

	// set pickup time for a book
	err = controller.SetBookPickup(r)
	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ReturnResponse(w, http.StatusOK, "Successfully save pickup time", nil)
}
