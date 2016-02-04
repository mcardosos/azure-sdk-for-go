// Package redis implements the Azure ARM Redis service API version 2015-08-01.
// .Net client wrapper for the REST API for Azure Redis Cache Management
// Service
package redis

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

const (
	// APIVersion is the version of the Redis
	APIVersion = "2015-08-01"

	// DefaultBaseURI is the default URI used for the service Redis
	DefaultBaseURI = "https://management.azure.com"
)

// ManagementClient is the base client for Redis.
type ManagementClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the ManagementClient client.
func New(subscriptionID string) ManagementClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return ManagementClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CreateOrUpdate create a redis cache, or replace (overwrite/recreate, with
// potential downtime) an existing cache
//
// resourceGroupName is the name of the resource group. name is the name of
// the redis cache. parameters is parameters supplied to the CreateOrUpdate
// redis operation.
func (client ManagementClient) CreateOrUpdate(resourceGroupName string, name string, parameters CreateOrUpdateParameters) (result ResourceWithAccessKey, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, name, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "redis/ManagementClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagementClient) CreateOrUpdatePreparer(resourceGroupName string, name string, parameters CreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagementClient) CreateOrUpdateResponder(resp *http.Response) (result ResourceWithAccessKey, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a redis cache. This operation takes a while to complete.
//
// resourceGroupName is the name of the resource group. name is the name of
// the redis cache.
func (client ManagementClient) Delete(resourceGroupName string, name string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "redis/ManagementClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagementClient) DeletePreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManagementClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a redis cache (resource description).
//
// resourceGroupName is the name of the resource group. name is the name of
// the redis cache.
func (client ManagementClient) Get(resourceGroupName string, name string) (result ResourceType, ae error) {
	req, err := client.GetPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "redis/ManagementClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagementClient) GetPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetResponder(resp *http.Response) (result ResourceType, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all redis caches in the specified subscription.
func (client ManagementClient) List() (result ListResult, ae error) {
	req, err := client.ListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "redis/ManagementClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ManagementClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Cache/Redis/"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListResponder(resp *http.Response) (result ListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ManagementClient) ListNextResults(lastResults ListResult) (result ListResult, ae error) {
	req, err := lastResults.ListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "redis/ManagementClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// ListByResourceGroup gets all redis caches in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client ManagementClient) ListByResourceGroup(resourceGroupName string) (result ListResult, ae error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "ListByResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "ListByResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "redis/ManagementClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ManagementClient) ListByResourceGroupPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListByResourceGroupResponder(resp *http.Response) (result ListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client ManagementClient) ListByResourceGroupNextResults(lastResults ListResult) (result ListResult, ae error) {
	req, err := lastResults.ListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "ListByResourceGroup", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "ListByResourceGroup", resp, "Failure sending next results request request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "redis/ManagementClient", "ListByResourceGroup", resp, "Failure responding to next results request request")
	}

	return
}

// ListKeys retrieve a redis cache's access keys. This operation requires
// write permission to the cache resource.
//
// resourceGroupName is the name of the resource group. name is the name of
// the redis cache.
func (client ManagementClient) ListKeys(resourceGroupName string, name string) (result ListKeysResult, ae error) {
	req, err := client.ListKeysPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "ListKeys", nil, "Failure preparing request")
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "ListKeys", resp, "Failure sending request")
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "redis/ManagementClient", "ListKeys", resp, "Failure responding to request")
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client ManagementClient) ListKeysPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/listKeys"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListKeysResponder(resp *http.Response) (result ListKeysResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegenerateKey regenerate redis cache's access keys. This operation requires
// write permission to the cache resource.
//
// resourceGroupName is the name of the resource group. name is the name of
// the redis cache. parameters is specifies which key to reset.
func (client ManagementClient) RegenerateKey(resourceGroupName string, name string, parameters RegenerateKeyParameters) (result ListKeysResult, ae error) {
	req, err := client.RegenerateKeyPreparer(resourceGroupName, name, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "RegenerateKey", nil, "Failure preparing request")
	}

	resp, err := client.RegenerateKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis/ManagementClient", "RegenerateKey", resp, "Failure sending request")
	}

	result, err = client.RegenerateKeyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "redis/ManagementClient", "RegenerateKey", resp, "Failure responding to request")
	}

	return
}

// RegenerateKeyPreparer prepares the RegenerateKey request.
func (client ManagementClient) RegenerateKeyPreparer(resourceGroupName string, name string, parameters RegenerateKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/regenerateKey"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// RegenerateKeySender sends the RegenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) RegenerateKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client ManagementClient) RegenerateKeyResponder(resp *http.Response) (result ListKeysResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
