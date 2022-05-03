package goalexa

import (
	"encoding/json"
)

type EchoResponse struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Response          EchoRespBody           `json:"response"`
}

type EchoRespBody struct {
	OutputSpeech     *EchoRespPayload         `json:"outputSpeech,omitempty"`
	Card             *EchoRespPayload         `json:"card,omitempty"`
	Reprompt         *EchoReprompt            `json:"reprompt,omitempty"`         // Pointer so it's dropped if empty in JSON response.
	ShouldEndSession *bool                    `json:"shouldEndSession,omitempty"` // Same
	Directives       []Directive              `json:"directives,omitempty"`
	CanFulfillIntent *CanFulfillIntentPayload `json:"canFulfillIntent,omitempty"`
}

type CanFulfillIntentPayload struct {
	CanFulfill CanFulfillIntentAnswer          `json:"canFulfill"`
	Slots      map[string]CanFulfillIntentSlot `json:"slots,omitempty"`
}

type CanFulfillIntentSlot struct {
	CanUnderstand CanFulfillIntentAnswer `json:"canUnderstand"`
	CanFulfill    CanFulfillIntentAnswer `json:"canFulfill"`
}

type CanFulfillIntentAnswer string

const (
	CanFulfillIntentAnswerYes   CanFulfillIntentAnswer = "YES"
	CanFulfillIntentAnswerNo    CanFulfillIntentAnswer = "NO"
	CanFulfillIntentAnswerMaybe CanFulfillIntentAnswer = "MAYBE"
)

type Directive map[string]interface{} // Shape differs wildly

type EchoReprompt struct {
	OutputSpeech EchoRespPayload `json:"outputSpeech,omitempty"`
}

type EchoRespImage struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

type EchoRespPayload struct {
	Type    string        `json:"type,omitempty"`
	Title   string        `json:"title,omitempty"`
	Text    string        `json:"text,omitempty"`
	SSML    string        `json:"ssml,omitempty"`
	Content string        `json:"content,omitempty"`
	Image   EchoRespImage `json:"image,omitempty"`
}

func NewEchoResponse() *EchoResponse {
	trueBool := true
	er := &EchoResponse{
		Version: "1.0",
		Response: EchoRespBody{
			ShouldEndSession: &trueBool,
			Directives:       []Directive{},
		},
		SessionAttributes: make(map[string]interface{}),
	}

	return er
}

func (this *EchoResponse) OutputSpeech(text string) *EchoResponse {
	this.Response.OutputSpeech = &EchoRespPayload{
		Type: "PlainText",
		Text: text,
	}

	return this
}

func (this *EchoResponse) Card(title string, content string) *EchoResponse {
	return this.SimpleCard(title, content)
}

func (this *EchoResponse) AccountLink(title string, content string) *EchoResponse {
	return this.LinkAccountCard()
}

func (this *EchoResponse) OutputSpeechSSML(text string) *EchoResponse {
	this.Response.OutputSpeech = &EchoRespPayload{
		Type: "SSML",
		SSML: text,
	}

	return this
}

func (this *EchoResponse) SimpleCard(title string, content string) *EchoResponse {
	this.Response.Card = &EchoRespPayload{
		Type:    "Simple",
		Title:   title,
		Content: content,
	}

	return this
}

func (this *EchoResponse) StandardCard(title string, content string, smallImg string, largeImg string) *EchoResponse {
	this.Response.Card = &EchoRespPayload{
		Type:    "Standard",
		Title:   title,
		Content: content,
	}

	if smallImg != "" {
		this.Response.Card.Image.SmallImageURL = smallImg
	}

	if largeImg != "" {
		this.Response.Card.Image.LargeImageURL = largeImg
	}

	return this
}

func (this *EchoResponse) LinkAccountCard() *EchoResponse {
	this.Response.Card = &EchoRespPayload{
		Type: "LinkAccount",
	}

	return this
}

func (this *EchoResponse) Reprompt(text string) *EchoResponse {
	this.Response.Reprompt = &EchoReprompt{
		OutputSpeech: EchoRespPayload{
			Type: "PlainText",
			Text: text,
		},
	}

	return this
}

func (this *EchoResponse) RepromptSSML(text string) *EchoResponse {
	this.Response.Reprompt = &EchoReprompt{
		OutputSpeech: EchoRespPayload{
			Type: "SSML",
			Text: text,
		},
	}

	return this
}

func (this *EchoResponse) JSON() ([]byte, error) {
	jsonStr, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}

	return jsonStr, nil
}

// AudioPlayer interface

type AudioPlayerPlayBehavior string

const (
	ReplaceAll      AudioPlayerPlayBehavior = "REPLACE_ALL"
	Enqueue         AudioPlayerPlayBehavior = "ENQUEUE"
	ReplaceEnqueued AudioPlayerPlayBehavior = "REPLACE_ENQUEUED"
)

func (this *EchoResponse) AudioPlayerPlay(
	behavior AudioPlayerPlayBehavior, streamUrl, token string, prevToken *string, offsetMs int,
) *EchoResponse {
	streamObj := map[string]interface{}{
		"url":                  streamUrl,
		"token":                token,
		"offsetInMilliseconds": offsetMs,
	}
	if prevToken != nil {
		streamObj["expectedPreviousToken"] = *prevToken
	}
	directive := map[string]interface{}{
		"type":         "AudioPlayer.Play",
		"playBehavior": behavior,
		"audioItem": map[string]interface{}{
			"stream": streamObj,
		},
	}
	this.Response.Directives = append(this.Response.Directives, directive)
	return this
}

func (this *EchoResponse) AudioPlayerStop() *EchoResponse {
	directive := map[string]interface{}{
		"type": "AudioPlayer.Stop",
	}
	this.Response.Directives = append(this.Response.Directives, directive)
	return this
}

type AudioPlayerClearQueueBehavior string

const (
	ClearEnqueued AudioPlayerClearQueueBehavior = "CLEAR_ENQUEUED"
	ClearAll      AudioPlayerClearQueueBehavior = "CLEAR_ALL"
)

func (this *EchoResponse) AudioPlayerClearQueue(clearBehavior AudioPlayerClearQueueBehavior) *EchoResponse {
	directive := map[string]interface{}{
		"type":          "AudioPlayer.ClearQueue",
		"clearBehavior": clearBehavior,
	}
	this.Response.Directives = append(this.Response.Directives, directive)
	return this
}

// VideoApp interface

func (this *EchoResponse) VideoAppLaunch(
	streamUrl, title, subtitle string,
) *EchoResponse {
	videoItemObj := map[string]interface{}{
		"source": streamUrl,
	}
	if title != "" || subtitle != "" {
		videoItemObj["metadata"] = map[string]interface{}{
			"title":    title,
			"subtitle": subtitle,
		}
	}
	directive := map[string]interface{}{
		"type":      "VideoApp.Launch",
		"videoItem": videoItemObj,
	}
	this.Response.Directives = append(this.Response.Directives, directive)
	return this
}
