# \DefaultAPI

All URIs are relative to *https://localhost:5001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NfconfigAccessMobilityGet**](DefaultAPI.md#NfconfigAccessMobilityGet) | **Get** /nfconfig/access-mobility | Get Access and Mobility Configuration
[**NfconfigPlmnGet**](DefaultAPI.md#NfconfigPlmnGet) | **Get** /nfconfig/plmn | Get PLMN ID Configuration
[**NfconfigPlmnSnssaiGet**](DefaultAPI.md#NfconfigPlmnSnssaiGet) | **Get** /nfconfig/plmn-snssai | Get PLMN-SNSSAI Configuration
[**NfconfigPolicyControlGet**](DefaultAPI.md#NfconfigPolicyControlGet) | **Get** /nfconfig/policy-control | Get Policy Control Configuration
[**NfconfigSessionManagementGet**](DefaultAPI.md#NfconfigSessionManagementGet) | **Get** /nfconfig/session-management | Get Session Management Configuration



## NfconfigAccessMobilityGet

> []AccessAndMobility NfconfigAccessMobilityGet(ctx).Execute()

Get Access and Mobility Configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.NfconfigAccessMobilityGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.NfconfigAccessMobilityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NfconfigAccessMobilityGet`: []AccessAndMobility
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.NfconfigAccessMobilityGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNfconfigAccessMobilityGetRequest struct via the builder pattern


### Return type

[**[]AccessAndMobility**](AccessAndMobility.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NfconfigPlmnGet

> []PlmnId NfconfigPlmnGet(ctx).Execute()

Get PLMN ID Configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.NfconfigPlmnGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.NfconfigPlmnGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NfconfigPlmnGet`: []PlmnId
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.NfconfigPlmnGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNfconfigPlmnGetRequest struct via the builder pattern


### Return type

[**[]PlmnId**](PlmnId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NfconfigPlmnSnssaiGet

> []PlmnSnssai NfconfigPlmnSnssaiGet(ctx).Execute()

Get PLMN-SNSSAI Configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.NfconfigPlmnSnssaiGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.NfconfigPlmnSnssaiGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NfconfigPlmnSnssaiGet`: []PlmnSnssai
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.NfconfigPlmnSnssaiGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNfconfigPlmnSnssaiGetRequest struct via the builder pattern


### Return type

[**[]PlmnSnssai**](PlmnSnssai.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NfconfigPolicyControlGet

> []PolicyControl NfconfigPolicyControlGet(ctx).Execute()

Get Policy Control Configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.NfconfigPolicyControlGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.NfconfigPolicyControlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NfconfigPolicyControlGet`: []PolicyControl
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.NfconfigPolicyControlGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNfconfigPolicyControlGetRequest struct via the builder pattern


### Return type

[**[]PolicyControl**](PolicyControl.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NfconfigSessionManagementGet

> []SessionManagement NfconfigSessionManagementGet(ctx).Execute()

Get Session Management Configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.NfconfigSessionManagementGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.NfconfigSessionManagementGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NfconfigSessionManagementGet`: []SessionManagement
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.NfconfigSessionManagementGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNfconfigSessionManagementGetRequest struct via the builder pattern


### Return type

[**[]SessionManagement**](SessionManagement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

