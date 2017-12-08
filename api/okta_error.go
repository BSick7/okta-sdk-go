package api

import "fmt"

type OktaError struct {
	Id         string           `json:"errorId"`
	StatusCode int              `json:"-"`
	ErrorCode  string           `json:"errorCode"`
	Summary    string           `json:"errorSummary"`
	Link       string           `json:"errorLink"`
	Causes     []OktaErrorCause `json:"errorCauses"`
}

type OktaErrorCause struct {
	Summary string `json:"errorSummary"`
}

func NewOktaError(res *ClientExecutorResponse) *OktaError {
	err := &OktaError{}
	res.BodyJSON(err)
	err.StatusCode = res.StatusCode()
	return err
}

func (e *OktaError) Error() string {
	return fmt.Sprintf("[id=%s] %d %s %s", e.Id, e.StatusCode, e.ErrorCode, e.Summary)
}

// Error Codes
// See https://developer.okta.com/reference/error_codes/index
/*
400
E0000001	API validation failed.
E0000002	The request was not valid.
E0000003	The request body was not well-formed.
E0000018	Bad request. Accept and/or Content-Type headers are likely not set.
E0000019	Bad request. Accept and/or Content-Type headers likely do not match supported values.
E0000020	Bad request.
E0000021	Bad request. Accept and/or Content-Type headers likely do not match supported values.
E0000024	Bad request. This operation on app metadata is not yet supported.
E0000025	App version assignment failed.
E0000027	Bad group push request.
E0000028	The request is missing a required parameter.
E0000029	Invalid paging request.
E0000030	Bad request. Invalid date. Dates must be of the form yyyy-MM-dd''T''HH:mm:ss.SSSZZ, e.g. 2013-01-01T12:00:00.000-07:00.
E0000031	Bad request. Invalid filter parameter.
E0000033	Bad request. Can't specify a search query and filter in the same request.
E0000037	Type mismatch exception.
E0000040	Application label must not be the same as an existing application label.
E0000041	Credentials should not be set on this resource based on the scheme.
E0000045	Field mapping bad request.
E0000050	Invalid SCIM data from client.
E0000053	Invalid SCIM filter.
E0000054	Invalid pagination properties.
E0000063	Invalid combination of parameters specified.
E0000081	Cannot modify the test attribute because it is a reserved attribute for this application.

401
E0000004	Authentication failed.
E0000011	Invalid token provided.
E0000015	You do not have permission to access the feature you are requesting.
E0000064	Password is expired and must be changed.

403
E0000005	Invalid session.
E0000006	You do not have permission to perform the requested action.
E0000013	Invalid client app ID.
E0000014	Update of credentials failed.
E0000016	Activation failed because the user is already active.
E0000017	Password reset failed.
E0000023	Operation failed because user profile is mastered under another system.
E0000032	Unlock is not allowed for this user.
E0000034	Forgot password not allowed on specified user.
E0000035	Change password not allowed on specified user.
E0000036	Change recovery question not allowed on specified user.
E0000038	This operation is not allowed in the user's current status.
E0000039	Operation on application settings failed.
E0000042	Setting the error page redirect URL failed.
E0000043	Self-service application assignment is not enabled.
E0000044	Self-service application assignment is not supported.
E0000046	Deactivate application for user forbidden.
E0000056	Delete application forbidden.
E0000057	Access to this application is denied due to a policy.
E0000058	Access to this application requires MFA.
E0000061	Tab error.
E0000069	User Locked.

404
E0000007	Not found.
E0000008	The requested path was not found.
E0000012	Unsupported media type.
E0000022	The endpoint does not support the provided HTTP method.
E0000026	This endpoint has been deprecated.
E0000048	Entity not found exception.
E0000060	Unsupported operation.

409
E0000055	Duplicate group.
E0000112	Cannot update this user because they are still being activated. Please try again in a few minutes.

429
E0000047	API call exceeded rate limit due to too many requests.
E0000109	An SMS message was recently sent. Please wait 30 seconds before trying again.

500
E0000009	Internal Server Error.
E0000049	Invalid SCIM data from SCIM implementation.
E0000051	No response from SCIM implementation.
E0000059	The connector configuration could not be tested. Make sure that the URL and Authentication Parameters are correct, and that there is an implementation available at the URL provided.

501
E0000052	Endpoint not implemented.

503
E0000010	Service is in read-only mode.
*/
