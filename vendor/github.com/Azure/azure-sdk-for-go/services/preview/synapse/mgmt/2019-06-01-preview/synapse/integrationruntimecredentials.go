package synapse

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IntegrationRuntimeCredentialsClient is the azure Synapse Analytics Management Client
type IntegrationRuntimeCredentialsClient struct {
	BaseClient
}

// NewIntegrationRuntimeCredentialsClient creates an instance of the IntegrationRuntimeCredentialsClient client.
func NewIntegrationRuntimeCredentialsClient(subscriptionID string) IntegrationRuntimeCredentialsClient {
	return NewIntegrationRuntimeCredentialsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationRuntimeCredentialsClientWithBaseURI creates an instance of the IntegrationRuntimeCredentialsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewIntegrationRuntimeCredentialsClientWithBaseURI(baseURI string, subscriptionID string) IntegrationRuntimeCredentialsClient {
	return IntegrationRuntimeCredentialsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Sync force the integration runtime to synchronize credentials across integration runtime nodes, and this will
// override the credentials across all worker nodes with those available on the dispatcher node. If you already have
// the latest credential backup file, you should manually import it (preferred) on any self-hosted integration runtime
// node than using this API directly.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// integrationRuntimeName - integration runtime name
func (client IntegrationRuntimeCredentialsClient) Sync(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationRuntimeCredentialsClient.Sync")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.IntegrationRuntimeCredentialsClient", "Sync", err.Error())
	}

	req, err := client.SyncPreparer(ctx, resourceGroupName, workspaceName, integrationRuntimeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeCredentialsClient", "Sync", nil, "Failure preparing request")
		return
	}

	resp, err := client.SyncSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeCredentialsClient", "Sync", resp, "Failure sending request")
		return
	}

	result, err = client.SyncResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeCredentialsClient", "Sync", resp, "Failure responding to request")
	}

	return
}

// SyncPreparer prepares the Sync request.
func (client IntegrationRuntimeCredentialsClient) SyncPreparer(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationRuntimeName": autorest.Encode("path", integrationRuntimeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/syncCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SyncSender sends the Sync request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationRuntimeCredentialsClient) SyncSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SyncResponder handles the response to the Sync request. The method always
// closes the http.Response Body.
func (client IntegrationRuntimeCredentialsClient) SyncResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
