package alexaapi

//
//
// Request: SessionEndedRequest

type RequestSessionEndedRequest struct {
	RequestCommon
	Reason string                    `json:"reason,omitempty"`
	Error  *SessionEndedRequestError `json:"error,omitempty"`
}

type SessionEndedRequestErrorType string

const (
	SessionEndedRequestErrorTypeUnspecified               SessionEndedRequestErrorType = ""
	SessionEndedRequestErrorTypeInvalidResponse           SessionEndedRequestErrorType = "INVALID_RESPONSE"
	SessionEndedRequestErrorTypeDeviceCommunicationError  SessionEndedRequestErrorType = "DEVICE_COMMUNICATION_ERROR"
	SessionEndedRequestErrorTypeInternalServiceError      SessionEndedRequestErrorType = "INTERNAL_SERVICE_ERROR"
	SessionEndedRequestErrorTypeAppServiceEndpointTimeout SessionEndedRequestErrorType = "ENDPOINT_TIMEOUT"
)

type SessionEndedRequestError struct {
	Type    SessionEndedRequestErrorType `json:"type"`
	Message string                       `json:"message"`
}
