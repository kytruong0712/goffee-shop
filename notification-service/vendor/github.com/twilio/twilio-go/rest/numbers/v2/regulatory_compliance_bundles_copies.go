/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateBundleCopy'
type CreateBundleCopyParams struct {
	// The string that you assigned to describe the copied bundle.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateBundleCopyParams) SetFriendlyName(FriendlyName string) *CreateBundleCopyParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Creates a new copy of a Bundle. It will internally create copies of all the bundle items (identities and documents) of the original bundle
func (c *ApiService) CreateBundleCopy(BundleSid string, params *CreateBundleCopyParams) (*NumbersV2BundleCopy, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Copies"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2BundleCopy{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListBundleCopy'
type ListBundleCopyParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListBundleCopyParams) SetPageSize(PageSize int) *ListBundleCopyParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBundleCopyParams) SetLimit(Limit int) *ListBundleCopyParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of BundleCopy records from the API. Request is executed immediately.
func (c *ApiService) PageBundleCopy(BundleSid string, params *ListBundleCopyParams, pageToken, pageNumber string) (*ListBundleCopyResponse, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Copies"

	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBundleCopyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists BundleCopy records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBundleCopy(BundleSid string, params *ListBundleCopyParams) ([]NumbersV2BundleCopy, error) {
	response, errors := c.StreamBundleCopy(BundleSid, params)

	records := make([]NumbersV2BundleCopy, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams BundleCopy records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBundleCopy(BundleSid string, params *ListBundleCopyParams) (chan NumbersV2BundleCopy, chan error) {
	if params == nil {
		params = &ListBundleCopyParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV2BundleCopy, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageBundleCopy(BundleSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamBundleCopy(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamBundleCopy(response *ListBundleCopyResponse, params *ListBundleCopyParams, recordChannel chan NumbersV2BundleCopy, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Results
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListBundleCopyResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListBundleCopyResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListBundleCopyResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBundleCopyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
