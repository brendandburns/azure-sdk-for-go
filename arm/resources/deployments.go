package resources

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// DeploymentsManagementClient is the client for the Deployments methods of
// the Resources service.
type DeploymentsManagementClient struct {
	ManagementClient
}

// NewDeploymentsManagementClient creates an instance of the
// DeploymentsManagementClient client.
func NewDeploymentsManagementClient(subscriptionID string) DeploymentsManagementClient {
	return NewDeploymentsManagementClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDeploymentsManagementClientWithBaseURI creates an instance of the
// DeploymentsManagementClient client.
func NewDeploymentsManagementClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsManagementClient {
	return DeploymentsManagementClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancel a currently running template deployment.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. deploymentName is the name of the deployment.
func (client DeploymentsManagementClient) Cancel(resourceGroupName string, deploymentName string) (result autorest.Response, ae error) {
	req, err := client.CancelPreparer(resourceGroupName, deploymentName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Cancel", "Failure preparing request")
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Cancel", "Failure sending request")
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Cancel", "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client DeploymentsManagementClient) CancelPreparer(resourceGroupName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    url.QueryEscape(deploymentName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/cancel"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsManagementClient) CancelSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client DeploymentsManagementClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CheckExistence checks whether deployment exists.
//
// resourceGroupName is the name of the resource group to check. The name is
// case insensitive. deploymentName is the name of the deployment.
func (client DeploymentsManagementClient) CheckExistence(resourceGroupName string, deploymentName string) (result autorest.Response, ae error) {
	req, err := client.CheckExistencePreparer(resourceGroupName, deploymentName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "CheckExistence", "Failure preparing request")
	}

	resp, err := client.CheckExistenceSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "CheckExistence", "Failure sending request")
	}

	result, err = client.CheckExistenceResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "CheckExistence", "Failure responding to request")
	}

	return
}

// CheckExistencePreparer prepares the CheckExistence request.
func (client DeploymentsManagementClient) CheckExistencePreparer(resourceGroupName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    url.QueryEscape(deploymentName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CheckExistenceSender sends the CheckExistence request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsManagementClient) CheckExistenceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent, http.StatusNotFound)
}

// CheckExistenceResponder handles the response to the CheckExistence request. The method always
// closes the http.Response Body.
func (client DeploymentsManagementClient) CheckExistenceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create a named template deployment using a template.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. deploymentName is the name of the deployment. parameters is
// additional parameters supplied to the operation.
func (client DeploymentsManagementClient) CreateOrUpdate(resourceGroupName string, deploymentName string, parameters Deployment) (result DeploymentExtended, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, deploymentName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DeploymentsManagementClient) CreateOrUpdatePreparer(resourceGroupName string, deploymentName string, parameters Deployment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    url.QueryEscape(deploymentName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsManagementClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusCreated)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DeploymentsManagementClient) CreateOrUpdateResponder(resp *http.Response) (result DeploymentExtended, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete begin deleting deployment.To determine whether the operation has
// finished processing the request, call GetLongRunningOperationStatus.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. deploymentName is the name of the deployment to be deleted.
func (client DeploymentsManagementClient) Delete(resourceGroupName string, deploymentName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, deploymentName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DeploymentsManagementClient) DeletePreparer(resourceGroupName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    url.QueryEscape(deploymentName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsManagementClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusAccepted, http.StatusNoContent)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DeploymentsManagementClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a deployment.
//
// resourceGroupName is the name of the resource group to get. The name is
// case insensitive. deploymentName is the name of the deployment.
func (client DeploymentsManagementClient) Get(resourceGroupName string, deploymentName string) (result DeploymentExtended, ae error) {
	req, err := client.GetPreparer(resourceGroupName, deploymentName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeploymentsManagementClient) GetPreparer(resourceGroupName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    url.QueryEscape(deploymentName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeploymentsManagementClient) GetResponder(resp *http.Response) (result DeploymentExtended, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get a list of deployments.
//
// resourceGroupName is the name of the resource group to filter by. The name
// is case insensitive. filter is the filter to apply on the operation. top
// is query parameters. If null is passed returns all deployments.
func (client DeploymentsManagementClient) List(resourceGroupName string, filter string, top *int) (result DeploymentListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, filter, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DeploymentsManagementClient) ListPreparer(resourceGroupName string, filter string, top *int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}
	if top != nil {
		queryParameters["$top"] = top
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeploymentsManagementClient) ListResponder(resp *http.Response) (result DeploymentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client DeploymentsManagementClient) ListNextResults(lastResults DeploymentListResult) (result DeploymentListResult, ae error) {
	req, err := lastResults.DeploymentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "List", "Failure responding to next results request request")
	}

	return
}

// Validate validate a deployment template.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. deploymentName is the name of the deployment. parameters is
// deployment to validate.
func (client DeploymentsManagementClient) Validate(resourceGroupName string, deploymentName string, parameters Deployment) (result DeploymentValidateResult, ae error) {
	req, err := client.ValidatePreparer(resourceGroupName, deploymentName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Validate", "Failure preparing request")
	}

	resp, err := client.ValidateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Validate", "Failure sending request")
	}

	result, err = client.ValidateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/DeploymentsManagementClient", "Validate", "Failure responding to request")
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client DeploymentsManagementClient) ValidatePreparer(resourceGroupName string, deploymentName string, parameters Deployment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    url.QueryEscape(deploymentName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.resources/deployments/{deploymentName}/validate"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsManagementClient) ValidateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusBadRequest)
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client DeploymentsManagementClient) ValidateResponder(resp *http.Response) (result DeploymentValidateResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}