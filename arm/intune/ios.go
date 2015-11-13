package intune

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
	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// IosClient is the microsoft.Intune Resource provider Api features in the
// swagger-2.0 specification
type IosClient struct {
	ManagementClient
}

// NewIosClient creates an instance of the IosClient client.
func NewIosClient(subscriptionID string) IosClient {
	return NewIosClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIosClientWithBaseURI creates an instance of the IosClient client.
func NewIosClientWithBaseURI(baseURI string, subscriptionID string) IosClient {
	return IosClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// AddAppForMAMPolicy add app to an iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyID is policy unique Id
// appID is application unique Id parameters is parameters supplied to add an
// app to an ios policy.
func (client IosClient) AddAppForMAMPolicy(hostName string, policyID string, appID string, parameters MAMPolicyAppIDOrGroupIDPayload) (result autorest.Response, ae error) {
	req, err := client.AddAppForMAMPolicyPreparer(hostName, policyID, appID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "AddAppForMAMPolicy", "Failure preparing request")
	}

	resp, err := client.AddAppForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "AddAppForMAMPolicy", "Failure sending request")
	}

	result, err = client.AddAppForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "AddAppForMAMPolicy", "Failure responding to request")
	}

	return
}

// AddAppForMAMPolicyPreparer prepares the AddAppForMAMPolicy request.
func (client IosClient) AddAppForMAMPolicyPreparer(hostName string, policyID string, appID string, parameters MAMPolicyAppIDOrGroupIDPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appId":    url.QueryEscape(appID),
		"hostName": url.QueryEscape(hostName),
		"policyId": url.QueryEscape(policyID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyId}/apps/{appId}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// AddAppForMAMPolicySender sends the AddAppForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) AddAppForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// AddAppForMAMPolicyResponder handles the response to the AddAppForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) AddAppForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// AddGroupForMAMPolicy add group to an iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyID is policy unique Id
// groupID is group Id parameters is parameters supplied to the Create or
// update app to an android policy operation.
func (client IosClient) AddGroupForMAMPolicy(hostName string, policyID string, groupID string, parameters MAMPolicyAppIDOrGroupIDPayload) (result autorest.Response, ae error) {
	req, err := client.AddGroupForMAMPolicyPreparer(hostName, policyID, groupID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "AddGroupForMAMPolicy", "Failure preparing request")
	}

	resp, err := client.AddGroupForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "AddGroupForMAMPolicy", "Failure sending request")
	}

	result, err = client.AddGroupForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "AddGroupForMAMPolicy", "Failure responding to request")
	}

	return
}

// AddGroupForMAMPolicyPreparer prepares the AddGroupForMAMPolicy request.
func (client IosClient) AddGroupForMAMPolicyPreparer(hostName string, policyID string, groupID string, parameters MAMPolicyAppIDOrGroupIDPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":  url.QueryEscape(groupID),
		"hostName": url.QueryEscape(hostName),
		"policyId": url.QueryEscape(policyID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyId}/groups/{groupId}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// AddGroupForMAMPolicySender sends the AddGroupForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) AddGroupForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// AddGroupForMAMPolicyResponder handles the response to the AddGroupForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) AddGroupForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdateMAMPolicy creates or updates iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyID is policy unique Id
// parameters is parameters supplied to the Create or update an android
// policy operation.
func (client IosClient) CreateOrUpdateMAMPolicy(hostName string, policyID string, parameters IOSMAMPolicy) (result IOSMAMPolicy, ae error) {
	req, err := client.CreateOrUpdateMAMPolicyPreparer(hostName, policyID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "CreateOrUpdateMAMPolicy", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "CreateOrUpdateMAMPolicy", "Failure sending request")
	}

	result, err = client.CreateOrUpdateMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "CreateOrUpdateMAMPolicy", "Failure responding to request")
	}

	return
}

// CreateOrUpdateMAMPolicyPreparer prepares the CreateOrUpdateMAMPolicy request.
func (client IosClient) CreateOrUpdateMAMPolicyPreparer(hostName string, policyID string, parameters IOSMAMPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
		"policyId": url.QueryEscape(policyID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyId}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateMAMPolicySender sends the CreateOrUpdateMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) CreateOrUpdateMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// CreateOrUpdateMAMPolicyResponder handles the response to the CreateOrUpdateMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) CreateOrUpdateMAMPolicyResponder(resp *http.Response) (result IOSMAMPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteAppForMAMPolicy delete App for Ios Policy
//
// hostName is location hostName for the tenant policyID is policy unique Id
// appID is application unique Id
func (client IosClient) DeleteAppForMAMPolicy(hostName string, policyID string, appID string) (result autorest.Response, ae error) {
	req, err := client.DeleteAppForMAMPolicyPreparer(hostName, policyID, appID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteAppForMAMPolicy", "Failure preparing request")
	}

	resp, err := client.DeleteAppForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteAppForMAMPolicy", "Failure sending request")
	}

	result, err = client.DeleteAppForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "DeleteAppForMAMPolicy", "Failure responding to request")
	}

	return
}

// DeleteAppForMAMPolicyPreparer prepares the DeleteAppForMAMPolicy request.
func (client IosClient) DeleteAppForMAMPolicyPreparer(hostName string, policyID string, appID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appId":    url.QueryEscape(appID),
		"hostName": url.QueryEscape(hostName),
		"policyId": url.QueryEscape(policyID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyId}/apps/{appId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteAppForMAMPolicySender sends the DeleteAppForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) DeleteAppForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteAppForMAMPolicyResponder handles the response to the DeleteAppForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) DeleteAppForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteGroupForMAMPolicy delete Group for iOS Policy
//
// hostName is location hostName for the tenant policyID is policy unique Id
// groupID is application unique Id
func (client IosClient) DeleteGroupForMAMPolicy(hostName string, policyID string, groupID string) (result autorest.Response, ae error) {
	req, err := client.DeleteGroupForMAMPolicyPreparer(hostName, policyID, groupID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteGroupForMAMPolicy", "Failure preparing request")
	}

	resp, err := client.DeleteGroupForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteGroupForMAMPolicy", "Failure sending request")
	}

	result, err = client.DeleteGroupForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "DeleteGroupForMAMPolicy", "Failure responding to request")
	}

	return
}

// DeleteGroupForMAMPolicyPreparer prepares the DeleteGroupForMAMPolicy request.
func (client IosClient) DeleteGroupForMAMPolicyPreparer(hostName string, policyID string, groupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":  url.QueryEscape(groupID),
		"hostName": url.QueryEscape(hostName),
		"policyId": url.QueryEscape(policyID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyId}/groups/{groupId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteGroupForMAMPolicySender sends the DeleteGroupForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) DeleteGroupForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteGroupForMAMPolicyResponder handles the response to the DeleteGroupForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) DeleteGroupForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteMAMPolicy delete Ios Policy
//
// hostName is location hostName for the tenant policyID is policy unique Id
func (client IosClient) DeleteMAMPolicy(hostName string, policyID string) (result autorest.Response, ae error) {
	req, err := client.DeleteMAMPolicyPreparer(hostName, policyID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteMAMPolicy", "Failure preparing request")
	}

	resp, err := client.DeleteMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteMAMPolicy", "Failure sending request")
	}

	result, err = client.DeleteMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "DeleteMAMPolicy", "Failure responding to request")
	}

	return
}

// DeleteMAMPolicyPreparer prepares the DeleteMAMPolicy request.
func (client IosClient) DeleteMAMPolicyPreparer(hostName string, policyID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
		"policyId": url.QueryEscape(policyID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteMAMPolicySender sends the DeleteMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) DeleteMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteMAMPolicyResponder handles the response to the DeleteMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) DeleteMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetAppForMAMPolicy get apps for an iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyID is policy unique Id
// filter is the filter to apply on the operation. selectParameter is select
// specific fields in entity.
func (client IosClient) GetAppForMAMPolicy(hostName string, policyID string, filter string, top *int, selectParameter string) (result ApplicationCollection, ae error) {
	req, err := client.GetAppForMAMPolicyPreparer(hostName, policyID, filter, top, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetAppForMAMPolicy", "Failure preparing request")
	}

	resp, err := client.GetAppForMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetAppForMAMPolicy", "Failure sending request")
	}

	result, err = client.GetAppForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetAppForMAMPolicy", "Failure responding to request")
	}

	return
}

// GetAppForMAMPolicyPreparer prepares the GetAppForMAMPolicy request.
func (client IosClient) GetAppForMAMPolicyPreparer(hostName string, policyID string, filter string, top *int, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
		"policyId": url.QueryEscape(policyID),
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
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyId}/apps"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetAppForMAMPolicySender sends the GetAppForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) GetAppForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetAppForMAMPolicyResponder handles the response to the GetAppForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) GetAppForMAMPolicyResponder(resp *http.Response) (result ApplicationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetGroupsForMAMPolicy returns groups for a given iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyID is policy name for
// the tenant
func (client IosClient) GetGroupsForMAMPolicy(hostName string, policyID string) (result GroupsCollection, ae error) {
	req, err := client.GetGroupsForMAMPolicyPreparer(hostName, policyID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetGroupsForMAMPolicy", "Failure preparing request")
	}

	resp, err := client.GetGroupsForMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetGroupsForMAMPolicy", "Failure sending request")
	}

	result, err = client.GetGroupsForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetGroupsForMAMPolicy", "Failure responding to request")
	}

	return
}

// GetGroupsForMAMPolicyPreparer prepares the GetGroupsForMAMPolicy request.
func (client IosClient) GetGroupsForMAMPolicyPreparer(hostName string, policyID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
		"policyId": url.QueryEscape(policyID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyId}/groups"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetGroupsForMAMPolicySender sends the GetGroupsForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) GetGroupsForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetGroupsForMAMPolicyResponder handles the response to the GetGroupsForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) GetGroupsForMAMPolicyResponder(resp *http.Response) (result GroupsCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMAMPolicies returns Intune iOSPolicies.
//
// hostName is location hostName for the tenant filter is the filter to apply
// on the operation. selectParameter is select specific fields in entity.
func (client IosClient) GetMAMPolicies(hostName string, filter string, top *int, selectParameter string) (result IOSMAMPolicyCollection, ae error) {
	req, err := client.GetMAMPoliciesPreparer(hostName, filter, top, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicies", "Failure preparing request")
	}

	resp, err := client.GetMAMPoliciesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicies", "Failure sending request")
	}

	result, err = client.GetMAMPoliciesResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicies", "Failure responding to request")
	}

	return
}

// GetMAMPoliciesPreparer prepares the GetMAMPolicies request.
func (client IosClient) GetMAMPoliciesPreparer(hostName string, filter string, top *int, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
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
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetMAMPoliciesSender sends the GetMAMPolicies request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) GetMAMPoliciesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetMAMPoliciesResponder handles the response to the GetMAMPolicies request. The method always
// closes the http.Response Body.
func (client IosClient) GetMAMPoliciesResponder(resp *http.Response) (result IOSMAMPolicyCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMAMPolicyByID returns Intune iOS policies.
//
// hostName is location hostName for the tenant policyID is policy unique Id
// selectParameter is select specific fields in entity.
func (client IosClient) GetMAMPolicyByID(hostName string, policyID string, selectParameter string) (result IOSMAMPolicy, ae error) {
	req, err := client.GetMAMPolicyByIDPreparer(hostName, policyID, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicyByID", "Failure preparing request")
	}

	resp, err := client.GetMAMPolicyByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicyByID", "Failure sending request")
	}

	result, err = client.GetMAMPolicyByIDResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicyByID", "Failure responding to request")
	}

	return
}

// GetMAMPolicyByIDPreparer prepares the GetMAMPolicyByID request.
func (client IosClient) GetMAMPolicyByIDPreparer(hostName string, policyID string, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
		"policyId": url.QueryEscape(policyID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetMAMPolicyByIDSender sends the GetMAMPolicyByID request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) GetMAMPolicyByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetMAMPolicyByIDResponder handles the response to the GetMAMPolicyByID request. The method always
// closes the http.Response Body.
func (client IosClient) GetMAMPolicyByIDResponder(resp *http.Response) (result IOSMAMPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PatchMAMPolicy patch an iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyID is policy unique Id
// parameters is parameters supplied to the Create or update an android
// policy operation.
func (client IosClient) PatchMAMPolicy(hostName string, policyID string, parameters IOSMAMPolicy) (result IOSMAMPolicy, ae error) {
	req, err := client.PatchMAMPolicyPreparer(hostName, policyID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "PatchMAMPolicy", "Failure preparing request")
	}

	resp, err := client.PatchMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "PatchMAMPolicy", "Failure sending request")
	}

	result, err = client.PatchMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "PatchMAMPolicy", "Failure responding to request")
	}

	return
}

// PatchMAMPolicyPreparer prepares the PatchMAMPolicy request.
func (client IosClient) PatchMAMPolicyPreparer(hostName string, policyID string, parameters IOSMAMPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
		"policyId": url.QueryEscape(policyID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyId}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// PatchMAMPolicySender sends the PatchMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) PatchMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// PatchMAMPolicyResponder handles the response to the PatchMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) PatchMAMPolicyResponder(resp *http.Response) (result IOSMAMPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
