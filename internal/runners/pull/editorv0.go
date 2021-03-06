package pull

type resultEditorV0 struct {
	Result *resultEditorV0Data `json:"result,omitempty"`
}

type resultEditorV0Data struct {
	Changed bool `json:"changed"`
}

func (f *outputFormat) editorV0Format() interface{} {
	return resultEditorV0{
		&resultEditorV0Data{
			f.Success,
		},
	}
}
