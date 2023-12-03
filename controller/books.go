package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"online-library/model"
	"online-library/utils"
	"regexp"
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
func GetBooksData(reqGetBooks *model.ReqGetBooks) (retData []*model.RetData, err error) {
	var (
		respLibraryApi  *model.RespLibraryApi
		baseURL         = utils.GetEnvByKey("LIBRARY_API_URL")
		params, subject string
	)

	subject = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(reqGetBooks.Subjects, "") // remove all characters except letters, numbers, spaces
	subject = strings.ReplaceAll(subject, " ", "_")                                           // change spaces into underscores
	subject = strings.ToLower(subject)                                                        // set string into lower case

	// hit library api
	request, err = http.NewRequest("GET", baseURL+"/subjects/"+subject+".json"+params, nil)
	if err != nil {
		err = errors.New(utils.ErrorRetrieveData())
		return nil, err
	}
	response, err = client.Do(request)
	if err != nil {
		err = errors.New(utils.ErrorRetrieveData())
		return nil, err
	}

	// parse json from response body
	err = json.NewDecoder(response.Body).Decode(&respLibraryApi)
	if err != nil {
		err = errors.New(utils.ErrorRetrieveData())
		return nil, err
	}

	defer response.Body.Close()

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

	return retData, nil
}

// SetBookPickup : to set pickup time of a book
func SetBookPickup(r *http.Request) (err error) {
	var (
		reqPickup  *model.ReqPickup
		pickupTime time.Time
		isExist    bool
	)

	// parse json from request body
	err = json.NewDecoder(r.Body).Decode(&reqPickup)
	if err != nil {
		err = errors.New(utils.ErrorFailedReadRequest())
		return err
	}

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
