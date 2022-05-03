package goalexa

type EchoRequest struct {
	Version string      `json:"version"`
	Session EchoSession `json:"session"`
	Request EchoReqBody `json:"request"`
	Context EchoContext `json:"context"`
}

type EchoSession struct {
	New         bool   `json:"new"`
	SessionID   string `json:"sessionId"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
	Attributes map[string]interface{} `json:"attributes"`
	User       struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
	} `json:"user"`
}

type EchoContext struct {
	System struct {
		ApiEndpoint    string `json:"apiEndpoint,omitempty"`
		ApiAccessToken string `json:"apiAccessToken,omitempty"`
		Device         struct {
			DeviceId string `json:"deviceId,omitempty"`
		} `json:"device,omitempty"`
		Application struct {
			ApplicationID string `json:"applicationId,omitempty"`
		} `json:"application,omitempty"`
		Unit struct {
			UnitId           string `json:"unitId,omitempty"`
			PersistentUnitId string `json:"persistentUnitId,omitempty"`
		} `json:"unit,omitempty"`
		User struct {
			AccessToken string `json:"accessToken,omitempty"`
			UserId      string `json:"userId,omitempty"`
			Permissions struct {
				ConsentToken string `json:"consentToken,omitempty"`
			} `json:"permissions,omitempty"`
		} `json:"user,omitempty"`
	} `json:"System,omitempty"`
	Viewport    *Viewport           `json:"Viewport,omitempty"`
	AudioPlayer *audioPlayerContext `json:"AudioPlayer,omitempty"`
}

type audioPlayerContext struct {
	PlayerActivity AudioPlayerActivity `json:"playerActivity"`
}

type AudioPlayerActivity string

const (
	AudioPlayerActivityIdle           = "IDLE"
	AudioPlayerActivityPlaying        = "PLAYING"
	AudioPlayerActivityPaused         = "PAUSED"
	AudioPlayerActivityFinished       = "FINISHED"
	AudioPlayerActivityBufferUnderrun = "BUFFER_UNDERRUN"
)

type Viewport struct {
	Dpi   int      `json:"dpi,omitempty"`
	Mode  string   `json:"mode,omitempty"`
	Touch []string `json:"touch,omitempty"`
}

type DialogState string

const (
	StartedDialogState    DialogState = "STARTED"
	InProgressDialogState             = "IN_PROGRESS"
	CompletedDialogState              = "COMPLETED"
)

type EchoReqBody struct {
	Type        string                  `json:"type"`
	RequestID   string                  `json:"requestId"`
	Timestamp   string                  `json:"timestamp"`
	DialogState DialogState             `json:"dialogState,omitempty"`
	Intent      EchoIntent              `json:"intent,omitempty"`
	Reason      string                  `json:"reason,omitempty"`
	Message     map[string]string       `json:"message"`
	Locale      string                  `json:"locale"`
	Body        EchoReqBodyInternalBody `json:"body,omitempty"`
}

type EchoReqBodyInternalBody struct {
	ReferenceId string `json:"referenceId,omitempty"`
}

type EchoIntent struct {
	Name  string              `json:"name"`
	Slots map[string]EchoSlot `json:"slots"`
}

type EchoSlot struct {
	Name string `json:"name"`

	SlotValue EchoSlotValue `json:"slotValue"`

	// Deprecated, use SlotValue instead
	Value string `json:"value"`

	// Deprecated, use SlotValue instead
	Resolutions *EchoSlotResolutions `json:"resolutions,omitempty"`
}

type EchoSlotValue struct {
	Type        EchoSlotValueType    `json:"type"`
	Resolutions *EchoSlotResolutions `json:"resolutions,omitempty"`
	Value       string               `json:"value"` // when type=Simple, value is the user's spoken utterance
	Values      []*EchoSlotValue     `json:"values,omitempty"`
}

type EchoSlotValueType string

const (
	EchoSlotValueType_Unknown EchoSlotValueType = ""
	EchoSlotValueType_Simple  EchoSlotValueType = "Simple"
	EchoSlotValueType_List    EchoSlotValueType = "List"
)

type EchoSlotResolutions struct {
	ResolutionsPerAuthority []EchoSlotAuthority `json:"resolutionsPerAuthority,omitempty"`
}

type EchoSlotAuthority struct {
	Authority string                             `json:"authority"`
	Values    []EchoSlotAuthorityValueCollection `json:"values,omitempty"`
}

type EchoSlotAuthorityValueCollection struct {
	Value *EchoSlotAuthorityValue `json:"value,omitempty"`
}

type EchoSlotAuthorityValue struct {
	Id   string `json:"id"`   // machine-readable id
	Name string `json:"name"` // canonical
}
