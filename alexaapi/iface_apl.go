package alexaapi

//
//
// Interface: Alexa.Presentation.APL

const (
	RequestTypeAplLoadIndexListData RequestType = "Alexa.Presentation.APL.LoadIndexListData"
	RequestTypeAplLoadTokenListData RequestType = "Alexa.Presentation.APL.LoadTokenListData"
	RequestTypeAplRuntimeError      RequestType = "Alexa.Presentation.APL.RuntimeError"
	RequestTypeAplUserEvent         RequestType = "Alexa.Presentation.APL.UserEvent"
)

const (
	DirectiveTypeAplRenderDocument   DirectiveType = "Alexa.Presentation.APL.RenderDocument"
	DirectiveTypeAplExecuteCommands  DirectiveType = "Alexa.Presentation.APL.ExecuteCommands"
	DirectiveTypeSendIndexListData   DirectiveType = "Alexa.Presentation.APL.SendIndexListData"
	DirectiveTypeSendTokenListData   DirectiveType = "Alexa.Presentation.APL.SendTokenListData"
	DirectiveTypeUpdateIndexListData DirectiveType = "Alexa.Presentation.APL.UpdateIndexListData"
)

const (
	AplCommandAnimateItem       string = "AnimateItem"
	AplCommandAutoPage          string = "AutoPage"
	AplCommandClearFocus        string = "ClearFocus"
	AplCommandFinish            string = "Finish"
	AplCommandIdle              string = "Idle"
	AplCommandOpenUrl           string = "OpenURL"
	AplCommandParallel          string = "Parallel"
	AplCommandReinflate         string = "Reinflate"
	AplCommandScroll            string = "Scroll"
	AplCommandScrollToComponent string = "ScrollToComponent"
	AplCommandScrollToIndex     string = "ScrollToIndex"
	AplCommandSelect            string = "Select"
	AplCommandSendEvent         string = "SendEvent"
	AplCommandSequential        string = "Sequential"
	AplCommandSetFocus          string = "SetFocus"
	AplCommandSetPage           string = "SetPage"
	AplCommandSetState          string = "SetState"
	AplCommandSetValue          string = "SetValue"
	AplCommandSpeakItem         string = "SpeakItem"
	AplCommandSpeakList         string = "SpeakList"
)

//
//
// Directive: Alexa.Presentation.APL.RenderDocument

type DirectiveAplRenderDocument struct {
	Type        DirectiveType  `json:"type"`
	Token       string         `json:"token"`
	Document    AplDocument    `json:"document"`
	Datasources map[string]any `json:"datasources,omitempty"`
}

type AplDocument struct {
	Type AplDocumentType `json:"type"`

	// When type = Link
	// e.g. doc://alexa/apl/documents/my-document
	Src string `json:"src,omitempty"`
}

type AplDocumentType string

const (
	AplDocumentTypeUnspecified AplDocumentType = ""
	AplDocumentTypeLink        AplDocumentType = "Link"
	AplDocumentTypeApl         AplDocumentType = "APL"
)

//
//
// Request: Alexa.Presentation.APL.UserEvent

type RequestAplUserEvent struct {
	RequestCommon
	Token      string         `json:"token,omitempty"`
	Arguments  []any          `json:"arguments,omitempty"`
	Source     map[string]any `json:"source,omitempty"`
	Components map[string]any `json:"components,omitempty"`
}
