package goalexa

//
//
// Interface: APL

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

// TODO: Fill in the following helpers

func CreateDirectiveAplRenderDocument() *Directive {
	return &Directive{
		Type: DirectiveTypeAplRenderDocument,
	}
}

func CreateDirectiveAplExecuteCommands() *Directive {
	return &Directive{
		Type: DirectiveTypeAplExecuteCommands,
	}
}

func CreateDirectiveSendIndexListData() *Directive {
	return &Directive{
		Type: DirectiveTypeSendIndexListData,
	}
}

func CreateDirectiveSendTokenListData() *Directive {
	return &Directive{
		Type: DirectiveTypeSendTokenListData,
	}
}

func CreateDirectiveUpdateIndexListData() *Directive {
	return &Directive{
		Type: DirectiveTypeUpdateIndexListData,
	}
}
