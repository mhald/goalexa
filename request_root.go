package goalexa

type RequestRoot struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Request Request `json:"request"`
	Context Context `json:"context"`
}

type Session struct {
	New         bool           `json:"new"`
	SessionId   string         `json:"sessionId"`
	Application Application    `json:"application"`
	Attributes  map[string]any `json:"attributes"`
	User        User           `json:"user"`
}

type Application struct {
	ApplicationId string `json:"applicationId"`
}

type UserPermissions struct {
	ConsentToken string `json:"consentToken,omitempty"`
}

type User struct {
	AccessToken string          `json:"accessToken,omitempty"`
	UserId      string          `json:"userId,omitempty"`
	Permissions UserPermissions `json:"permissions,omitempty"`
}

type Device struct {
	DeviceId             string `json:"deviceId,omitempty"`
	SupportedInterfaces  any    `json:"supportedInterfaces,omitempty"`
	PersistentEndpointId string `json:"persistentEndpointId,omitempty"`
}

type Unit struct {
	UnitId           string `json:"unitId,omitempty"`
	PersistentUnitId string `json:"persistentUnitId,omitempty"`
}

type Person struct {
	PersonId    string `json:"personId,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}

type System struct {
	ApiEndpoint    string      `json:"apiEndpoint,omitempty"`
	ApiAccessToken string      `json:"apiAccessToken,omitempty"`
	Device         Device      `json:"device,omitempty"`
	Application    Application `json:"application,omitempty"`
	Unit           Unit        `json:"unit,omitempty"`
	User           User        `json:"user,omitempty"`
	Person         Person      `json:"person,omitempty"`
}

type Context struct {
	System      System              `json:"System,omitempty"`
	Viewport    *Viewport           `json:"Viewport,omitempty"`
	AudioPlayer *AudioPlayerContext `json:"AudioPlayer,omitempty"`
}

type Viewport struct {
	Dpi   int      `json:"dpi,omitempty"`
	Mode  string   `json:"mode,omitempty"`
	Touch []string `json:"touch,omitempty"`
}

type RequestType string

const (
	RequestTypeUnspecified RequestType = ""
)
