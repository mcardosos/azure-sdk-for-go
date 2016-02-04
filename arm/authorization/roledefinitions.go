// Package authorization implements the Azure ARM Authorization service API
// version 2015-07-01.
package authorization

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

// RoleDefinitionsClient is the client for the RoleDefinitions methods of the
// Authorization service.
type RoleDefinitionsClient struct {
	ManagementClient
}

// NewRoleDefinitionsClient creates an instance of the RoleDefinitionsClient
// client.
func NewRoleDefinitionsClient(subscriptionID string) RoleDefinitionsClient {
	return NewRoleDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoleDefinitionsClientWithBaseURI creates an instance of the
// RoleDefinitionsClient client.
func NewRoleDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) RoleDefinitionsClient {
	return RoleDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a role definition.
//
// scope is scope roleDefinitionID is role definition id. roleDefinition is
// role definition.
func (client RoleDefinitionsClient) CreateOrUpdate(scope string, roleDefinitionID string, roleDefinition RoleDefinition) (result RoleDefinition, ae error) {
	req, err := client.CreateOrUpdatePreparer(scope, roleDefinitionID, roleDefinition)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RoleDefinitionsClient) CreateOrUpdatePreparer(scope string, roleDefinitionID string, roleDefinition RoleDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleDefinitionId": url.QueryEscape(roleDefinitionID),
		"scope":            scope,
		"subscriptionId":   url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}"),
		autorest.WithJSON(roleDefinition),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the role definition.
//
// scope is scope roleDefinitionID is role definition id.
func (client RoleDefinitionsClient) Delete(scope string, roleDefinitionID string) (result RoleDefinition, ae error) {
	req, err := client.DeletePreparer(scope, roleDefinitionID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RoleDefinitionsClient) DeletePreparer(scope string, roleDefinitionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleDefinitionId": url.QueryEscape(roleDefinitionID),
		"scope":            scope,
		"subscriptionId":   url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) DeleteResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get role definition by name (GUID).
//
// scope is scope roleDefinitionID is role definition Id
func (client RoleDefinitionsClient) Get(scope string, roleDefinitionID string) (result RoleDefinition, ae error) {
	req, err := client.GetPreparer(scope, roleDefinitionID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RoleDefinitionsClient) GetPreparer(scope string, roleDefinitionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleDefinitionId": url.QueryEscape(roleDefinitionID),
		"scope":            scope,
		"subscriptionId":   url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID get role definition by name (GUID).
//
// roleDefinitionID is fully qualified role definition Id
func (client RoleDefinitionsClient) GetByID(roleDefinitionID string) (result RoleDefinition, ae error) {
	req, err := client.GetByIDPreparer(roleDefinitionID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "GetByID", nil, "Failure preparing request")
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "GetByID", resp, "Failure sending request")
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client RoleDefinitionsClient) GetByIDPreparer(roleDefinitionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleDefinitionId": roleDefinitionID,
		"subscriptionId":   url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{roleDefinitionId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetByIDResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all role definitions that are applicable at scope and above. Use
// atScopeAndBelow filter to search below the given scope as well
//
// scope is scope filter is the filter to apply on the operation.
func (client RoleDefinitionsClient) List(scope string, filter string) (result RoleDefinitionListResult, ae error) {
	req, err := client.ListPreparer(scope, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RoleDefinitionsClient) ListPreparer(scope string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope":          scope,
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{scope}/providers/Microsoft.Authorization/roleDefinitions"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) ListResponder(resp *http.Response) (result RoleDefinitionListResult, err error) {
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
func (client RoleDefinitionsClient) ListNextResults(lastResults RoleDefinitionListResult) (result RoleDefinitionListResult, ae error) {
	req, err := lastResults.RoleDefinitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "authorization/RoleDefinitionsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}
