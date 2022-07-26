package alexaapi

//
//
// Request: IntentRequest

type RequestIntentRequest struct {
	RequestCommon
	DialogState DialogState `json:"dialogState,omitempty"`
	Intent      *Intent     `json:"intent,omitempty"`
}

type DialogState string

const (
	DialogStateUnspecified DialogState = ""
	DialogStateStarted     DialogState = "STARTED"
	DialogStateInProgress  DialogState = "IN_PROGRESS"
	DialogStateCompleted   DialogState = "COMPLETED"
)

type Intent struct {
	Name               string             `json:"name"`
	ConfirmationStatus ConfirmationStatus `json:"confirmationStatus,omitempty"`
	Slots              map[string]Slot    `json:"slots"`
}

type ConfirmationStatus string

const (
	ConfirmationStatusUnspecified ConfirmationStatus = ""
	ConfirmationStatusNone        ConfirmationStatus = "NONE"
	ConfirmationStatusConfirmed   ConfirmationStatus = "CONFIRMED"
	ConfirmationStatusDenied      ConfirmationStatus = "DENIED"
)

type Slot struct {
	Name               string             `json:"name"`
	ConfirmationStatus ConfirmationStatus `json:"confirmationStatus,omitempty"`
	SlotValue          SlotValue          `json:"slotValue"`
	// Deprecated, use SlotValue instead
	Value string `json:"value"`
	// Deprecated, use SlotValue instead
	Resolutions *SlotResolutions `json:"resolutions,omitempty"`
}

type SlotValue struct {
	Type        SlotValueType    `json:"type"`
	Resolutions *SlotResolutions `json:"resolutions,omitempty"`
	Value       string           `json:"value"` // when type=Simple, value is the user's spoken utterance
	Values      []*SlotValue     `json:"values,omitempty"`
}

type SlotValueType string

const (
	SlotValueTypeUnspecified SlotValueType = ""
	SlotValueTypeSimple      SlotValueType = "Simple"
	SlotValueTypeList        SlotValueType = "List"
)

type SlotResolutions struct {
	ResolutionsPerAuthority []SlotAuthority `json:"resolutionsPerAuthority,omitempty"`
}

type SlotAuthority struct {
	Authority string                         `json:"authority"`
	Values    []SlotAuthorityValueCollection `json:"values,omitempty"`
}

type SlotAuthorityValueCollection struct {
	Value *SlotAuthorityValue `json:"value,omitempty"`
}

type SlotAuthorityValue struct {
	Id   string `json:"id"`   // machine-readable id
	Name string `json:"name"` // canonical
}
