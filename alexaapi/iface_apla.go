package alexaapi

//
//
// Interface: Alexa.Presentation.APLA

const (
	DirectiveTypeAplaRenderDocument DirectiveType = "Alexa.Presentation.APLA.RenderDocument"
)

//
//
// Directive: Alexa.Presentation.APLA.RenderDocument

type DirectiveAlexaPresentationAplaRenderDocument struct {
	Type DirectiveType `json:"type"`
	// TODO
}

func CreateDirectiveAplaRenderDocument() *DirectiveAlexaPresentationAplaRenderDocument {
	return &DirectiveAlexaPresentationAplaRenderDocument{
		Type: DirectiveTypeAplaRenderDocument,
	}
}
