package controller

import (
	"errors"
	"online-library/model"
	"online-library/utils"
	"testing"
)

// TestGetBooksData : to test function GetBooksData
func TestGetBooksData(t *testing.T) {
	type respData struct {
		lenRetData int
		statusCode int
		err        error
	}

	// set multiple test case
	testCase := map[string]struct {
		input  *model.ReqGetBooks
		result *respData
	}{
		"no subject": {
			input: &model.ReqGetBooks{
				Subjects: "",
			},
			result: &respData{
				lenRetData: 0,
				statusCode: 400,
				err:        errors.New(utils.ErrorRequired("genre of books")),
			},
		},
		"no data": {
			input: &model.ReqGetBooks{
				Subjects: "adafa statq",
			},
			result: &respData{
				lenRetData: 0,
				statusCode: 200,
				err:        errors.New(utils.ErrorEmptyData()),
			},
		},
		"success": {
			input: &model.ReqGetBooks{
				Subjects: "love",
				Limit:    10,
			},
			result: &respData{
				lenRetData: 10,
				statusCode: 200,
				err:        nil,
			},
		},
	}

	// run each test case
	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			res, statCode, error := GetBooksData(test.input)
			if len(res) != test.result.lenRetData || statCode != test.result.statusCode || (error != nil && error.Error() != test.result.err.Error()) {
				t.Errorf("Error test : %v", error)
			}
		})
	}
}

// TestSetBookPickup : to test function SetBookPickup
func TestSetBookPickup(t *testing.T) {
	// set multiple test case
	testCase := map[string]struct {
		input *model.ReqPickup
		err   error
	}{
		"wrong time": {
			input: &model.ReqPickup{
				Key:        "book1",
				PickupTime: "2023-20-12 25:60",
			},
			err: errors.New(utils.ErrorFailedReadRequest()),
		},
		"pickup before now": {
			input: &model.ReqPickup{
				Key:        "book2",
				PickupTime: "10-11-2023 12:00",
			},
			err: errors.New(utils.ErrorCantBeBeforeNow()),
		},
		"success": {
			input: &model.ReqPickup{
				Key:        "book3",
				PickupTime: "15-12-2023 10:00",
			},
			err: nil,
		},
		"duplicate book": {
			input: &model.ReqPickup{
				Key:        "book3",
				PickupTime: "25-12-2023 12:00",
			},
			err: errors.New(utils.ErrorAlreadyPickedup()),
		},
	}

	// run each test case
	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			error := SetBookPickup(test.input)
			if (error != nil && test.err != nil && error.Error() != test.err.Error()) || (test.err != nil && error == nil) || (test.err == nil && error != nil) {
				t.Errorf("Error test : %v", error)
			}
		})
	}

	// check if pickupInfo data has been filled
	expectedTime := "15-12-2023 10:00"
	if pickupInfo["book3"] != expectedTime {
		t.Errorf("Error pickup time. Expected %s", expectedTime)
	}
}
