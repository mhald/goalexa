package alexaapi

import (
	"context"
	"encoding/json"

	"github.com/tidwall/gjson"
)

type Request interface {
	GetType() RequestType
	GetRequestId() string
	GetTimestamp() string
	GetLocale() string
	GetOtherFields() map[string]any
}

type RequestType string

// These are request types that don't belong to a specific interface.
//
// Other requests that DO belong to a specific interface are defined in their
// respective "iface_" files.
const (
	RequestTypeUnspecified             RequestType = ""
	RequestTypeCanFulfillIntentRequest RequestType = "CanFulfillIntentRequest"
	RequestTypeLaunchRequest           RequestType = "LaunchRequest"
	RequestTypeIntentRequest           RequestType = "IntentRequest"
	RequestTypeSessionEndedRequest     RequestType = "SessionEndedRequest"
	RequestTypeSessionResumedRequest   RequestType = "SessionResumedRequest"
)

type RequestCommon struct {
	Type      RequestType `json:"type"`
	RequestId string      `json:"requestId"`
	Timestamp string      `json:"timestamp"`
	Locale    string      `json:"locale"`

	// All fields other than those above
	otherFields map[string]any
}

func (rc *RequestCommon) GetType() RequestType {
	return rc.Type
}

func (rc *RequestCommon) GetRequestId() string {
	return rc.RequestId
}

func (rc *RequestCommon) GetTimestamp() string {
	return rc.Timestamp
}

func (rc *RequestCommon) GetLocale() string {
	return rc.Locale
}

func (rc *RequestCommon) GetOtherFields() map[string]any {
	return rc.otherFields
}

// Attempts to set a strongly-typed value into the root.Request field
// by "looking ahead" at the contents of the "type" field.
//
// If no match is found, a simple RequestCommon is used.
func SetRequestViaLookahead(ctx context.Context, reqRoot *RequestRoot, rootJson []byte) error {
	reqJson := []byte(gjson.GetBytes(rootJson, "request").String())

	unmarshalIntoOther := func(o *map[string]any) error {
		err := json.Unmarshal([]byte(reqJson), o)
		if err != nil {
			return err
		}
		delete(*o, "type")
		delete(*o, "requestId")
		delete(*o, "timestamp")
		delete(*o, "locale")
		return nil
	}

	requestType := RequestType(gjson.GetBytes(reqJson, "type").String())
	switch requestType {
	case RequestTypeIntentRequest:
		var r RequestIntentRequest
		err := json.Unmarshal(reqJson, &r)
		if err != nil {
			return err
		}
		err = unmarshalIntoOther(&r.otherFields)
		if err != nil {
			return err
		}
		reqRoot.Request = &r
		return nil

	case RequestTypeSessionEndedRequest:
		var r RequestSessionEndedRequest
		err := json.Unmarshal(reqJson, &r)
		if err != nil {
			return err
		}
		err = unmarshalIntoOther(&r.otherFields)
		if err != nil {
			return err
		}
		reqRoot.Request = &r
		return nil

	case RequestTypeSessionResumedRequest:
		var r RequestSessionResumedRequest
		err := json.Unmarshal(reqJson, &r)
		if err != nil {
			return err
		}
		err = unmarshalIntoOther(&r.otherFields)
		if err != nil {
			return err
		}
		reqRoot.Request = &r
		return nil

	case RequestTypeAplUserEvent:
		var r RequestAplUserEvent
		err := json.Unmarshal(reqJson, &r)
		if err != nil {
			return err
		}
		err = unmarshalIntoOther(&r.otherFields)
		if err != nil {
			return err
		}
		reqRoot.Request = &r
		return nil

	case RequestTypeAlexaAuthorizationGrant:
		var r RequestAlexaAuthorizationGrant
		err := json.Unmarshal(reqJson, &r)
		if err != nil {
			return err
		}
		err = unmarshalIntoOther(&r.otherFields)
		if err != nil {
			return err
		}
		reqRoot.Request = &r
		return nil
	}

	var r RequestCommon
	err := json.Unmarshal(reqJson, &r)
	if err != nil {
		return err
	}
	err = unmarshalIntoOther(&r.otherFields)
	if err != nil {
		return err
	}
	reqRoot.Request = &r

	return nil
}
