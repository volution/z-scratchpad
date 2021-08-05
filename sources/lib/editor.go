

package zscratchpad




type Editor struct {
	globals *Globals
	index *Index
}




func EditorNew (_globals *Globals, _index *Index) (*Editor, *Error) {
	_editor := & Editor {
			globals : _globals,
		}
	return _editor, nil
}




func EditorDocumentOpen (_editor *Editor, _document *Document) (*Error) {
	return nil
}

