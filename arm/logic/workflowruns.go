package logic

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

// WorkflowRunsManagementClient is the client for the WorkflowRuns methods of
// the Logic service.
type WorkflowRunsManagementClient struct {
	ManagementClient
}

// NewWorkflowRunsManagementClient creates an instance of the
// WorkflowRunsManagementClient client.
func NewWorkflowRunsManagementClient(subscriptionID string) WorkflowRunsManagementClient {
	return NewWorkflowRunsManagementClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowRunsManagementClientWithBaseURI creates an instance of the
// WorkflowRunsManagementClient client.
func NewWorkflowRunsManagementClientWithBaseURI(baseURI string, subscriptionID string) WorkflowRunsManagementClient {
	return WorkflowRunsManagementClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancels a workflow run.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. runName is the workflow run name.
func (client WorkflowRunsManagementClient) Cancel(resourceGroupName string, workflowName string, runName string) (result autorest.Response, ae error) {
	req, err := client.CancelPreparer(resourceGroupName, workflowName, runName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "Cancel", "Failure preparing request")
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "Cancel", "Failure sending request")
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "Cancel", "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client WorkflowRunsManagementClient) CancelPreparer(resourceGroupName string, workflowName string, runName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"runName":           url.QueryEscape(runName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/cancel"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunsManagementClient) CancelSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client WorkflowRunsManagementClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a workflow run.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. runName is the workflow run name.
func (client WorkflowRunsManagementClient) Get(resourceGroupName string, workflowName string, runName string) (result WorkflowRun, ae error) {
	req, err := client.GetPreparer(resourceGroupName, workflowName, runName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowRunsManagementClient) GetPreparer(resourceGroupName string, workflowName string, runName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"runName":           url.QueryEscape(runName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunsManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowRunsManagementClient) GetResponder(resp *http.Response) (result WorkflowRun, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of workflow runs.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. top is the number of items to be included in the result. filter is
// the filter to apply on the operation.
func (client WorkflowRunsManagementClient) List(resourceGroupName string, workflowName string, top *int, filter string) (result WorkflowRunListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, workflowName, top, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkflowRunsManagementClient) ListPreparer(resourceGroupName string, workflowName string, top *int, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = top
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunsManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowRunsManagementClient) ListResponder(resp *http.Response) (result WorkflowRunListResult, err error) {
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
func (client WorkflowRunsManagementClient) ListNextResults(lastResults WorkflowRunListResult) (result WorkflowRunListResult, ae error) {
	req, err := lastResults.WorkflowRunListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowRunsManagementClient", "List", "Failure responding to next results request request")
	}

	return
}