// Package search implements the Azure ARM Search service API version
// 2015-02-28.
// client that can be used to manage Azure Search services and API keys.
package search

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
	"net/url"
)

// QueryKeysClient is the client for the QueryKeys methods of the Search
// service.
type QueryKeysClient struct {
	ManagementClient
}

// NewQueryKeysClient creates an instance of the QueryKeysClient client.
func NewQueryKeysClient(subscriptionID string) QueryKeysClient {
	return NewQueryKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQueryKeysClientWithBaseURI creates an instance of the QueryKeysClient
// client.
func NewQueryKeysClientWithBaseURI(baseURI string, subscriptionID string) QueryKeysClient {
	return QueryKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List returns the list of query API keys for the given Azure Search service.
//
// resourceGroupName is the name of the resource group within the current
// subscription. serviceName is the name of the Search service for which to
// list query keys.
func (client QueryKeysClient) List(resourceGroupName string, serviceName string) (result ListQueryKeysResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, serviceName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search/QueryKeysClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "search/QueryKeysClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "search/QueryKeysClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client QueryKeysClient) ListPreparer(resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"serviceName":       url.QueryEscape(serviceName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{serviceName}/listQueryKeys"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client QueryKeysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client QueryKeysClient) ListResponder(resp *http.Response) (result ListQueryKeysResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
