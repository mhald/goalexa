package alexaapi

//
//
// Interface: Dialog

const (
	DirectiveTypeDialogDelegate              DirectiveType = "Dialog.Delegate"
	DirectiveTypeDialogElicitSlot            DirectiveType = "Dialog.ElicitSlot"
	DirectiveTypeDialogConfirmSlot           DirectiveType = "Dialog.ConfirmSlot"
	DirectiveTypeDialogConfirmIntent         DirectiveType = "Dialog.ConfirmIntent"
	DirectiveTypeDialogUpdateDynamicEntities DirectiveType = "Dialog.UpdateDynamicEntities"
)

func CreateDirectiveDialogDelegate(updatedIntent *Intent) *Directive {
	return &Directive{
		Type:          DirectiveTypeDialogDelegate,
		UpdatedIntent: updatedIntent,
	}
}

func CreateDirectiveDialogElicitSlot(updatedIntent *Intent, slotToElicit string) *Directive {
	return &Directive{
		Type:          DirectiveTypeDialogElicitSlot,
		UpdatedIntent: updatedIntent,
		SlotToElicit:  slotToElicit,
	}
}

func CreateDirectiveDialogConfirmSlot(updatedIntent *Intent, slotToConfirm string) *Directive {
	return &Directive{
		Type:          DirectiveTypeDialogConfirmSlot,
		UpdatedIntent: updatedIntent,
		SlotToConfirm: slotToConfirm,
	}
}

func CreateDirectiveDialogConfirmIntent(updatedIntent *Intent) *Directive {
	return &Directive{
		Type:          DirectiveTypeDialogConfirmIntent,
		UpdatedIntent: updatedIntent,
	}
}

// TODO
func CreateDirectiveDialogUpdateDynamicEntities() *Directive {
	return &Directive{
		Type: DirectiveTypeDialogUpdateDynamicEntities,
	}
}
