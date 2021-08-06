

package zscratchpad


import "os/exec"




type Editor struct {
	globals *Globals
	index *Index
}

type editSession struct {
	globals *Globals
	editor *Editor
	documentOld *Document
	documentNew *Document
	command *exec.Cmd
	error *Error
}




func EditorNew (_globals *Globals, _index *Index) (*Editor, *Error) {
	_editor := & Editor {
			globals : _globals,
			index : _index,
		}
	return _editor, nil
}




func EditorDocumentOpen (_editor *Editor, _document *Document) (*Error) {
	
	_path := _document.PathOriginal
	if _path == "" {
		return errorw (0xbdb59e67, nil)
	}
	
	logf ('d', 0xf0c4b437, "[editor-session]  created for `%s`...", _document.PathOriginal)
	
	_session := & editSession {
			globals : _editor.globals,
			editor : _editor,
			documentOld : _document,
		}
	
	go editorSessionRun (_session)
	
	return nil
}




func editorSessionRun (_session *editSession) (*Error) {
	
	_command := & exec.Cmd {
			Path : "/usr/bin/howl",
			Args : []string {"howl", "--", _session.documentOld.PathOriginal},
			Env : nil,
			Stdin : nil,
			Stdout : nil,
			Stderr : nil,
		}
	
	logf ('d', 0x0edfabbf, "[editor-session]  launching editor for `%s`...", _session.documentOld.PathOriginal)
	
	if _error := _command.Start (); _error != nil {
		_session.error = errorw (0x4b48b0bc, _error)
		return editorSessionClose (_session)
	}
	
	_session.command = _command
	
	logf ('d', 0xff9ec344, "[editor-session]  waiting editor for `%s`...", _session.documentOld.PathOriginal)
	
	if _error := _session.command.Wait (); _error != nil {
		_session.error = errorw (0xe877f161, _error)
		return editorSessionClose (_session)
	}
	
	logf ('d', 0x48f7d5f5, "[editor-session]  reloading document for `%s`...", _session.documentOld.PathOriginal)
	
	if _document_0, _error := DocumentReload (_session.documentOld); _error == nil {
		_session.documentNew = _document_0
	} else {
		_session.error = _error
		return editorSessionClose (_session)
	}
	
	if _session.editor.index == nil {
		return editorSessionClose (_session)
	}
	
	_session.globals.Mutex.Lock ()
	defer _session.globals.Mutex.Unlock ()
	
	logf ('d', 0x44c67acc, "[editor-session]  reindexing document for `%s`...", _session.documentOld.PathOriginal)
	
	if _error := IndexDocumentUpdate (_session.editor.index, _session.documentNew, _session.documentOld); _error != nil {
		_session.error = _error
		return editorSessionClose (_session)
	}
	
	return editorSessionClose (_session)
}




func editorSessionClose (_session *editSession) (*Error) {
	
	if _session.error != nil {
		logErrorf ('e', 0x35a898d8, _session.error, "[editor-session]  failed for `%s`!", _session.documentOld.PathOriginal)
		return _session.error
	}
	
	logf ('d', 0x42c39bbe, "[editor-session]  succeeded for `%s`;", _session.documentOld.PathOriginal)
	return nil
}

