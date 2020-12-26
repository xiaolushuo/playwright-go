package playwright

type fileChooserImpl struct {
	page          Page
	elementHandle ElementHandle
	isMultiple    bool
}

func (f *fileChooserImpl) Page() Page {
	return f.page
}

func (f *fileChooserImpl) Element() ElementHandle {
	return f.elementHandle
}

func (f *fileChooserImpl) IsMultiple() bool {
	return f.isMultiple
}

func (f *fileChooserImpl) SetFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error {
	return f.elementHandle.SetInputFiles(files, options...)
}

type InputFile struct {
	Name     string
	MimeType string
	Buffer   []byte
}

func newFileChooser(page Page, elementHandle ElementHandle, isMultiple bool) *fileChooserImpl {
	return &fileChooserImpl{
		page:          page,
		elementHandle: elementHandle,
		isMultiple:    isMultiple,
	}
}
