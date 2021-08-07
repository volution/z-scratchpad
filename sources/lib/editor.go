

package zscratchpad


import "os"
import "os/exec"
import "path"




type Editor struct {
	globals *Globals
	index *Index
}

type editSession struct {
	globals *Globals
	editor *Editor
	library *Library
	documentOld *Document
	documentNew *Document
	path string
	file *os.File
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




func EditorDocumentEdit (_editor *Editor, _library *Library, _document *Document, _synchronous bool) (*Error) {
	
	_path := _document.Path
	if _path == "" {
		return errorw (0xbdb59e67, nil)
	}
	
	logf ('d', 0x0edfabbf, "[editor-session]  opening file for `%s`...", _path)
	
	_file, _error := os.OpenFile (_path, os.O_RDWR, 0)
	if _error != nil {
		return errorw (0xa51cdc41, _error)
	}
	
	_session := & editSession {
			globals : _editor.globals,
			editor : _editor,
			library : _library,
			documentOld : _document,
			path : _path,
			file : _file,
		}
	
	if !_synchronous {
		go editSessionRun (_session)
		return nil
	}
	
	return editSessionRun (_session)
}


func EditorDocumentCreate (_editor *Editor, _library *Library, _documentName string, _synchronous bool) (*Error) {
	
	_path := path.Join (_library.Path, _documentName) + ".txt"
	
	logf ('d', 0x6292b948, "[editor-session]  creating file for `%s`...", _path)
	
	_file, _error := os.OpenFile (_path, os.O_RDWR | os.O_CREATE | os.O_EXCL, 0o640)
	if _error != nil {
		return errorw (0x5d8b586a, _error)
	}
	
	_session := & editSession {
			globals : _editor.globals,
			editor : _editor,
			library : _library,
			path : _path,
			file : _file,
		}
	
	if !_synchronous {
		go editSessionRun (_session)
		return nil
	}
	
	return editSessionRun (_session)
}


func editSessionRun (_session *editSession) (*Error) {
	
	logf ('d', 0x0edfabbf, "[editor-session]  launching editor for `%s`...", _session.path)
	
	_command := & exec.Cmd {
			Path : "/usr/bin/howl",
			Args : []string {"howl", "--", _session.path},
			Env : nil,
			Stdin : nil,
			Stdout : nil,
			Stderr : nil,
		}
	if _error := _command.Start (); _error != nil {
		_session.error = errorw (0x4b48b0bc, _error)
		return editSessionClose (_session)
	}
	_session.command = _command
	
	logf ('d', 0xff9ec344, "[editor-session]  waiting editor for `%s`...", _session.path)
	
	if _error := _session.command.Wait (); _error != nil {
		_session.error = errorw (0xe877f161, _error)
		return editSessionClose (_session)
	}
	
	return editSessionFinalize (_session)
}


func editSessionFinalize (_session *editSession) (*Error) {
	
	if _session.documentOld != nil {
		
		logf ('d', 0x48f7d5f5, "[editor-session]  reloading document for `%s`...", _session.path)
		
		if _document_0, _error := DocumentReload (_session.documentOld); _error == nil {
			_session.documentNew = _document_0
		} else {
			_session.error = _error
			return editSessionClose (_session)
		}
		
	} else {
		
		logf ('d', 0xff9b1d5d, "[editor-session]  loading document for `%s`...", _session.path)
		
		if _document_0, _error := DocumentLoadFromPath (_session.path); _error == nil {
			_session.documentNew = _document_0
		} else {
			_session.error = _error
			return editSessionClose (_session)
		}
	}
	
	if (_session.documentNew.Library == "") && (_session.library != nil) {
		_session.documentNew.Library = _session.library.Identifier
	}
	
	_useFileNameAsIdentifier := false
	_useFileExtensionAsFormat := false
	if _session.library != nil {
		_useFileNameAsIdentifier = _session.library.UseFileNameAsIdentifier
		_useFileExtensionAsFormat = _session.library.UseFileExtensionAsFormat
	}
	if _error := DocumentResolveIdentifier (_session.documentNew, _useFileNameAsIdentifier); _error != nil {
		_session.error = _error
		return editSessionClose (_session)
	}
	if _error := DocumentResolveFormat (_session.documentNew, _useFileExtensionAsFormat); _error != nil {
		_session.error = _error
		return editSessionClose (_session)
	}
	
	if _session.editor.index == nil {
		return editSessionClose (_session)
	}
	
	_session.globals.Mutex.Lock ()
	defer _session.globals.Mutex.Unlock ()
	
	if _session.documentOld != nil {
		
		logf ('d', 0x44c67acc, "[editor-session]  reindexing document for `%s`...", _session.path)
		
		if _error := IndexDocumentUpdate (_session.editor.index, _session.documentNew, _session.documentOld); _error != nil {
			_session.error = _error
			return editSessionClose (_session)
		}
		
	} else {
		
		logf ('d', 0x5ee2c034, "[editor-session]  indexing document for `%s`...", _session.path)
		
		if _error := IndexDocumentInclude (_session.editor.index, _session.documentNew); _error != nil {
			_session.error = _error
			return editSessionClose (_session)
		}
	}
	
	return editSessionClose (_session)
}


func editSessionClose (_session *editSession) (*Error) {
	
	if _session.error != nil {
		logErrorf ('e', 0x35a898d8, _session.error, "[editor-session]  failed for `%s`!", _session.path)
		return _session.error
	}
	
	logf ('d', 0x42c39bbe, "[editor-session]  succeeded for `%s`;", _session.path)
	
	return nil
}

