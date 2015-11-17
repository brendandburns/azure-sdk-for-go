package subscriptions

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

// TenantsClient is the client for the Tenants methods of the Subscriptions
// service.
type TenantsClient struct {
	Client
}

// NewTenantsClient creates an instance of the TenantsClient client.
func NewTenantsClient(subscriptionID string) TenantsClient {
	return NewTenantsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTenantsClientWithBaseURI creates an instance of the TenantsClient client.
func NewTenantsClientWithBaseURI(baseURI string, subscriptionID string) TenantsClient {
	return TenantsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets a list of the tenantIds.
func (client TenantsClient) List() (result TenantListResult, ae error) {
	req, err := client.ListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "subscriptions/TenantsClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "subscriptions/TenantsClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "subscriptions/TenantsClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client TenantsClient) ListPreparer() (*http.Request, error) {
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
		autorest.WithPath("/tenants"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TenantsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TenantsClient) ListResponder(resp *http.Response) (result TenantListResult, err error) {
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
func (client TenantsClient) ListNextResults(lastResults TenantListResult) (result TenantListResult, ae error) {
	req, err := lastResults.TenantListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "subscriptions/TenantsClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "subscriptions/TenantsClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "subscriptions/TenantsClient", "List", "Failure responding to next results request request")
	}

	return
}