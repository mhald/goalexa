package goalexa

type ResponseRoot struct {
	Version           string         `json:"version"`
	SessionAttributes map[string]any `json:"sessionAttributes,omitempty"`
	Response          Response       `json:"response"`
}

type DirectiveType string

const (
	DirectiveTypeUnspecified DirectiveType = ""
)

// Polymorphic
type Directive struct {
	// Common
	Type DirectiveType `json:"type"`

	// VideoApp.Launch
	VideoItem *VideoItem `json:"videoItem,omitempty"`

	// AudioPlayer.Play
	PlayBehavior AudioPlayerPlayBehavior `json:"playBehavior,omitempty"`
	AudioItem    *AudioPlayerAudioItem   `json:"audioItem,omitempty"`

	// AudioPlayer.ClearQueue
	ClearBehavior AudioPlayerClearQueueBehavior `json:"clearBehavior,omitempty"`

	// Dialog.Delegate
	// Dialog.ElicitSlot
	// Dialog.ConfirmSlot
	// Dialog.ConfirmIntent
	UpdatedIntent *Intent `json:"updatedIntent,omitempty"`

	// Dialog.ElicitSlot
	SlotToElicit string `json:"slotToElicit,omitempty"`

	// Dialog.ConfirmSlot
	SlotToConfirm string `json:"slotToConfirm,omitempty"`

	// Dialog.UpdateDynamicEntities
	// TODO

	// Alexa.Presentation.APL.RenderDocument
	// TODO

	// Alexa.Presentation.APL.ExecuteCommands
	// TODO

	// Alexa.Presentation.APL.SendIndexListData
	// TODO

	// Alexa.Presentation.APL.SendTokenListData
	// TODO

	// Alexa.Presentation.APL.UpdateIndexListData
	// TODO
}

type Response struct {
	OutputSpeech     *OutputSpeech     `json:"outputSpeech,omitempty"`
	Card             *Card             `json:"card,omitempty"`
	Reprompt         *Reprompt         `json:"reprompt,omitempty"`
	ShouldEndSession *bool             `json:"shouldEndSession,omitempty"`
	Directives       []*Directive      `json:"directives,omitempty"`
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
	Type         OutputSpeechType        `json:"type,omitempty"`
	Text         string                  `json:"text,omitempty"`
	SSML         string                  `json:"ssml,omitempty"`
	PlayBehavior AudioPlayerPlayBehavior `json:"playBehavior,omitempty"`
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
			Directives: []*Directive{},
		},
		SessionAttributes: map[string]any{},
	}
}

func (rr *ResponseRoot) AddDirective(directive *Directive) {
	if rr.Response.Directives == nil {
		rr.Response.Directives = []*Directive{}
	}
	rr.Response.Directives = append(rr.Response.Directives, directive)
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
