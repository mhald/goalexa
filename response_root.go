package goalexa

type ResponseRoot struct {
	Version           string         `json:"version"`
	SessionAttributes map[string]any `json:"sessionAttributes,omitempty"`
	Response          Response       `json:"response"`
}

type Response struct {
	OutputSpeech     *OutputSpeech     `json:"outputSpeech,omitempty"`
	Card             *Card             `json:"card,omitempty"`
	Reprompt         *Reprompt         `json:"reprompt,omitempty"`
	ShouldEndSession *bool             `json:"shouldEndSession,omitempty"`
	Directives       []any             `json:"directives,omitempty"`
	CanFulfillIntent *CanFulfillIntent `json:"canFulfillIntent,omitempty"`
}

type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}

type OutputSpeechType string

const (
	OutputSpeechTypeUnspecified OutputSpeechType = ""
	OutputSpeechTypePlainText   OutputSpeechType = "PlainText"
	OutputSpeechTypeSSML        OutputSpeechType = "SSML"
)

type OutputSpeech struct {
	Type         OutputSpeechType `json:"type,omitempty"`
	Text         string           `json:"text,omitempty"`
	SSML         string           `json:"ssml,omitempty"`
	PlayBehavior AudioPlayerPlayBehavior     `json:"playBehavior,omitempty"`
}

type CardImage struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

type CardType string

const (
	CardTypeUnspecified              CardType = ""
	CardTypeSimple                   CardType = "Simple"
	CardTypeStandard                 CardType = "Standard"
	CardTypeLinkAccount              CardType = "LinkAccount"
	CardTypeAskForPermissionsConsent CardType = "AskForPermissionsConsent"
)

type Card struct {
	Type    CardType  `json:"type"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Text    string    `json:"text,omitempty"`
	Image   CardImage `json:"image,omitempty"`
}

func NewResponseRoot() *ResponseRoot {
	return &ResponseRoot{
		Version: "1.0",
		Response: Response{
			Directives: []any{},
		},
		SessionAttributes: map[string]any{},
	}
}

//
//
// CanFulfillIntentRequest response structs

type CanFulfillIntent struct {
	CanFulfill CfirAnswer          `json:"canFulfill"`
	Slots      map[string]CfirSlot `json:"slots,omitempty"`
}

type CfirSlot struct {
	CanUnderstand CfirAnswer `json:"canUnderstand"`
	CanFulfill    CfirAnswer `json:"canFulfill"`
}

type CfirAnswer string

const (
	CfirAnswerYes   CfirAnswer = "YES"
	CfirAnswerNo    CfirAnswer = "NO"
	CfirAnswerMaybe CfirAnswer = "MAYBE"
)
