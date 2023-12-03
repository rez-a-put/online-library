package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"online-library/model"
	"online-library/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	request    *http.Request
	response   *http.Response
	client     = &http.Client{}
	pickupInfo = make(map[string]string)
	timeLayout = "02-01-2006 15:04"
)

// GetBooksData : to hit library api and return data
func GetBooksData(reqGetBooks *model.ReqGetBooks) (retData []*model.RetData, statusCode int, err error) {
	var (
		respLibraryApi  *model.RespLibraryApi
		baseURL         = utils.GetEnvByKey("LIBRARY_API_URL")
		params, subject string
	)

	// simple genre/subject validation
	if reqGetBooks.Subjects == "" {
		err = errors.New(utils.ErrorRequired("genre of books"))
		return nil, http.StatusBadRequest, err
	}

	subject = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(reqGetBooks.Subjects, "") // remove all characters except letters, numbers, spaces
	subject = strings.ReplaceAll(subject, " ", "_")                                           // change spaces into underscores
	subject = strings.ToLower(subject)                                                        // set string into lower case

	params = "?"
	if reqGetBooks.Limit > 0 {
		params += "limit=" + strconv.Itoa(reqGetBooks.Limit) + "&"
	}
	if reqGetBooks.Offset > 0 {
		params += "offset=" + strconv.Itoa(reqGetBooks.Offset) + "&"
	}
	params = strings.TrimSuffix(params, "&") // remove trailed ampersand
	params = strings.TrimSuffix(params, "?") // remove trailed question mark

	// hit library api
	request, err = http.NewRequest("GET", baseURL+"/subjects/"+subject+".json"+params, nil)
	if err != nil {
		err = errors.New(utils.ErrorRetrieveData())
		return nil, http.StatusBadRequest, err
	}
	response, err = client.Do(request)
	if err != nil {
		err = errors.New(utils.ErrorRetrieveData())
		return nil, http.StatusBadRequest, err
	}

	// parse json from response body
	err = json.NewDecoder(response.Body).Decode(&respLibraryApi)
	if err != nil {
		err = errors.New(utils.ErrorRetrieveData())
		return nil, http.StatusBadRequest, err
	}

	defer response.Body.Close()

	if respLibraryApi.DataCount == 0 {
		err = errors.New(utils.ErrorEmptyData())
		return nil, http.StatusOK, err
	}

	// setup return data
	for _, v := range respLibraryApi.Works {
		var authors string

		for _, val := range v.Authors {
			authors += val.Name
		}

		data := &model.RetData{
			Key:            v.Key,
			Title:          v.Title,
			Authors:        authors,
			EditionNumber:  v.EditionCount,
			PickupSchedule: pickupInfo[v.Key],
		}

		retData = append(retData, data)
	}

	return retData, http.StatusOK, nil
}

// SetBookPickup : to set pickup time of a book
func SetBookPickup(reqPickup *model.ReqPickup) (err error) {
	var (
		pickupTime time.Time
		isExist    bool
	)

	// parse pickup time
	pickupTime, err = time.Parse(timeLayout, reqPickup.PickupTime)
	if err != nil {
		err = errors.New(utils.ErrorFailedReadRequest())
		return err
	}

	// validate pickup time
	if pickupTime.Before(time.Now()) {
		err = errors.New(utils.ErrorCantBeBeforeNow())
		return err
	}

	// check if book already has a pickup time
	_, isExist = pickupInfo[reqPickup.Key]
	if isExist {
		err = errors.New(utils.ErrorAlreadyPickedup())
		return err
	}

	pickupInfo[reqPickup.Key] = pickupTime.Format(timeLayout)

	return nil
}
