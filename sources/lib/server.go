

package zscratchpad


import "bytes"
import "net"
import "net/http"
import "strings"

import html_template "html/template"
import text_template "text/template"




type Server struct {
	index *Index
	editor *Editor
	templates *Templates
	listener net.Listener
	http *http.Server
	globals *Globals
}


type serverHandler struct {
	server *Server
}




func ServerNew (_globals *Globals, _index *Index, _editor *Editor, _listener net.Listener) (*Server, *Error) {
	
	_templates := (*Templates) (nil)
	if _templates_0, _error := TemplatesNew (); _error == nil {
		_templates = _templates_0
	} else {
		return nil, _error
	}
	
	_server := & Server {
			globals : _globals,
			index : _index,
			editor : _editor,
			templates : _templates,
			listener : _listener,
		}
	
	_handler := & serverHandler {
			server : _server,
		}
	
	_http := & http.Server {
			Handler : _handler,
		}
	
	_server.http = _http
	
	return _server, nil
}




func ServerRun (_server *Server) (*Error) {
	
	_error := _server.http.Serve (_server.listener)
	if _error != http.ErrServerClosed {
		return errorw (0x12143330, _error)
	}
	
	return nil
}




func (_handler *serverHandler) ServeHTTP (_response http.ResponseWriter, _request *http.Request) () {
	if _error := ServerHandle (_handler.server, _request, _response); _error != nil {
		_message := _error.ToError () .Error ()
		http.Error (_response, _message, http.StatusInternalServerError)
	}
}




func ServerHandle (_server *Server, _request *http.Request, _response http.ResponseWriter) (*Error) {
	
	_server.globals.Mutex.Lock ()
	defer _server.globals.Mutex.Unlock ()
	
	if _request.Method != "GET" {
		return errorw (0x7f32157c, nil)
	}
	
	_path := _request.URL.Path
	if ! strings.HasPrefix (_path, "/") {
		return errorw (0x828c5f04, nil)
	}
	
	if _path == "/__/heartbeat" {
		return respondWithBuffer (_response, "text/plain", bytes.NewBufferString ("OK\n"))
	}
	
	if (_path == "/d") || (_path == "/d/") || (_path == "/documents") || (_path == "/documents/") {
		return ServerHandleDocumentsIndex (_server, _response)
	}
	if (_path == "/l") || (_path == "/l/") || (_path == "/libraries") || (_path == "/libraries/") {
		return ServerHandleLibrariesIndex (_server, _response)
	}
	
	if strings.HasPrefix (_path, "/l/") {
		_identifier := _path[3:]
		return ServerHandleLibraryView (_server, _identifier, _response)
	}
	if strings.HasPrefix (_path, "/d/") {
		_identifier := _path[3:]
		return ServerHandleDocumentView (_server, _identifier, _response)
	}
	if strings.HasPrefix (_path, "/dx/html/") {
		_identifier := _path[9:]
		return ServerHandleDocumentExportHtml (_server, _identifier, _response)
	}
	if strings.HasPrefix (_path, "/dx/text/") {
		_identifier := _path[9:]
		return ServerHandleDocumentExportText (_server, _identifier, _response)
	}
	if strings.HasPrefix (_path, "/dx/source/") {
		_identifier := _path[11:]
		return ServerHandleDocumentExportSource (_server, _identifier, _response)
	}
	
	if strings.HasPrefix (_path, "/de/") {
		_identifier := _path[4:]
		return ServerHandleDocumentEdit (_server, _identifier, _response)
	}
	
	if (_path == "/dc") || (_path == "/dc/") {
		_identifier := ""
		return ServerHandleDocumentCreate (_server, _identifier, _response)
	}
	if strings.HasPrefix (_path, "/dc/") {
		_identifier := _path[4:]
		return ServerHandleDocumentCreate (_server, _identifier, _response)
	}
	
	if _path == "/__/version" {
		return ServerHandleVersion (_server, _response)
	}
	if _path == "/__/sources.md5" {
		return ServerHandleSourcesMd5 (_server, _response)
	}
	if _path == "/__/sources.cpio" {
		return ServerHandleSourcesCpio (_server, _response)
	}
	
	switch _path {
		case "/favicon.ico", "/favicon.png" :
			_path = "/assets/favicons/" + _path[1:]
		case "/apple-touch-icon.png" :
			_path = "/assets/favicons/favicon.png"
	}
	
	if strings.HasPrefix (_path, "/assets/") {
		_path := _path[1:]
		return ServerHandleAsset (_server, _path, _response)
	}
	
	return errorw (0x7b01a78b, nil)
}




func ServerHandleLibrariesIndex (_server *Server, _response http.ResponseWriter) (*Error) {
	_libraries, _error := IndexLibrariesSelectAll (_server.index)
	if _error != nil {
		return _error
	}
	_context := struct { Libraries []*Library } { _libraries }
	return respondWithHtmlTemplate (_response, _server.templates.librariesIndexHtml, _context)
}


func ServerHandleDocumentsIndex (_server *Server, _response http.ResponseWriter) (*Error) {
	_documents, _error := IndexDocumentsSelectAll (_server.index)
	if _error != nil {
		return _error
	}
	_context := struct { Documents []*Document } { _documents }
	return respondWithHtmlTemplate (_response, _server.templates.documentsIndexHtml, _context)
}




func ServerHandleLibraryView (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	_library, _error := serverLibraryResolve (_server, _identifierUnsafe)
	if _error != nil {
		return _error
	}
	_documents, _error := IndexDocumentsSelectInLibrary (_server.index, _library.Identifier)
	if _error != nil {
		return _error
	}
	_context := struct {
			Library *Library
			Documents []*Document
		} {
			_library,
			_documents,
		}
	return respondWithHtmlTemplate (_response, _server.templates.libraryViewHtml, _context)
}




func ServerHandleDocumentView (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	_document, _library, _error := serverDocumentAndLibraryResolve (_server, _identifierUnsafe)
	if _error != nil {
		return _error
	}
	_documentHtml, _error := DocumentRenderToHtml (_document)
	if _error != nil {
		return _error
	}
	_context := struct {
			Library *Library
			Document *Document
			DocumentHtml html_template.HTML
		} {
			_library,
			_document,
			html_template.HTML (_documentHtml),
		}
	return respondWithHtmlTemplate (_response, _server.templates.documentViewHtml, _context)
}


func ServerHandleDocumentExportHtml (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	_document, _error := serverDocumentResolve (_server, _identifierUnsafe)
	if _error != nil {
		return _error
	}
	_documentHtml, _error := DocumentRenderToHtml (_document)
	if _error != nil {
		return _error
	}
	_context := struct {
			Document *Document
			DocumentHtml html_template.HTML
		} {
			_document,
			html_template.HTML (_documentHtml),
		}
	return respondWithHtmlTemplate (_response, _server.templates.documentExportHtml, _context)
}


func ServerHandleDocumentExportText (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	_document, _error := serverDocumentResolve (_server, _identifierUnsafe)
	if _error != nil {
		return _error
	}
	_documentText, _error := DocumentRenderToText (_document)
	if _error != nil {
		return _error
	}
	_context := struct {
			Document *Document
			DocumentText string
		} {
			_document,
			_documentText,
		}
	return respondWithTextTemplate (_response, _server.templates.documentExportText, _context)
}


func ServerHandleDocumentExportSource (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	_document, _error := serverDocumentResolve (_server, _identifierUnsafe)
	if _error != nil {
		return _error
	}
	_documentSource, _error := DocumentRenderToSource (_document)
	if _error != nil {
		return _error
	}
	_context := struct {
			Document *Document
			DocumentSource string
		} {
			_document,
			_documentSource,
		}
	return respondWithTextTemplate (_response, _server.templates.documentExportSource, _context)
}




func ServerHandleDocumentCreate (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	if _identifierUnsafe == "" {
		// FIXME:  Add support for random document creation!
		return errorw (0x19f48aa6, nil)
	}
	_libraryIdentifier := ""
	_documentName := ""
	if _libraryIdentifier == "" {
		if _libraryIdentifier_0, _error := LibraryParseIdentifier (_identifierUnsafe); _error == nil {
			_libraryIdentifier = _libraryIdentifier_0
		}
	}
	if _libraryIdentifier == "" {
		if _, _libraryIdentifier_0, _documentName_0, _error := DocumentParseIdentifier (_identifierUnsafe); _error == nil {
			_libraryIdentifier = _libraryIdentifier_0
			_documentName = _documentName_0
		}
	}
	if _libraryIdentifier == "" {
		return errorw (0x4f21b7fb, nil)
	}
	if _documentName == "" {
		_documentName = generateRandomToken ()
	}
	_identifier, _error := DocumentFormatIdentifier (_libraryIdentifier, _documentName)
	if _error != nil {
		return _error
	}
	_documentExisting, _error := IndexDocumentResolve (_server.index, _identifier)
	if _error != nil {
		return _error
	}
	if _documentExisting != nil {
		return errorw (0x054e7a60, nil)
	}
	if _libraryIdentifier == "" {
		return errorw (0x2b40ce32, nil)
	}
	_library, _error := IndexLibraryResolve (_server.index, _libraryIdentifier)
	if _error != nil {
		return _error
	}
	if _library == nil {
		return errorw (0x5e581595, nil)
	}
	if _server.editor == nil {
		return errorw (0x14317f29, nil)
	}
	_error = EditorDocumentCreate (_server.editor, _library, _documentName, false)
	if _error != nil {
		return _error
	}
	http.Error (_response, "", http.StatusNoContent)
	return nil
}


func ServerHandleDocumentEdit (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	_document, _library, _error := serverDocumentAndLibraryResolve (_server, _identifierUnsafe)
	if _error != nil {
		return _error
	}
	if _server.editor == nil {
		return errorw (0xee28afb6, nil)
	}
	_error = EditorDocumentEdit (_server.editor, _library, _document, false)
	if _error != nil {
		return _error
	}
	http.Error (_response, "", http.StatusNoContent)
	return nil
}




func ServerHandleVersion (_server *Server, _response http.ResponseWriter) (*Error) {
	_context := struct {
			
			ProjectUrl string
			
			BuildTarget string
			BuildTargetArch string
			BuildTargetOs string
			BuildCompilerType string
			BuildCompilerVersion string
			
			BuildVersion string
			BuildNumber string
			BuildTimestamp string
			
			BuildGitHash string
			BuildSourcesHash string
			
			UnameNode string
			UnameSystem string
			UnameRelease string
			UnameVersion string
			UnameMachine string
			
		} {
			
			PROJECT_URL,
			
			BUILD_TARGET,
			BUILD_TARGET_ARCH,
			BUILD_TARGET_OS,
			BUILD_COMPILER_TYPE,
			BUILD_COMPILER_VERSION,
			
			BUILD_VERSION,
			BUILD_NUMBER,
			BUILD_TIMESTAMP,
			
			BUILD_GIT_HASH,
			BUILD_SOURCES_HASH,
			
			UNAME_NODE,
			UNAME_SYSTEM,
			UNAME_RELEASE,
			UNAME_VERSION,
			UNAME_MACHINE,
			
		}
	return respondWithHtmlTemplate (_response, _server.templates.versionHtml, _context)
}


func ServerHandleSourcesMd5 (_server *Server, _response http.ResponseWriter) (*Error) {
	return respondWithBuffer (_response, "text/plain; charset=utf-8", bytes.NewBufferString (BUILD_SOURCES_MD5))
}

func ServerHandleSourcesCpio (_server *Server, _response http.ResponseWriter) (*Error) {
	_response.Header () .Add ("Content-Encoding", "gzip")
	return respondWithBuffer (_response, "text/plain; charset=utf-8", bytes.NewBuffer (BUILD_SOURCES_CPIO_GZ))
}



func ServerHandleAsset (_server *Server, _path string, _response http.ResponseWriter) (*Error) {
	_contentType, _body, _error := TemplatesAssetResolve (_server.templates, _path)
	if _error != nil {
		return _error
	}
	return respondWithBuffer (_response, _contentType, bytes.NewBuffer (_body))
}




func serverLibraryResolve (_server *Server, _identifierUnsafe string) (*Library, *Error) {
	if _identifierUnsafe == "" {
		return nil, errorw (0xbef72625, nil)
	}
	_identifier, _error := LibraryParseIdentifier (_identifierUnsafe)
	if _error != nil {
		return nil, _error
	}
	_library, _error := IndexLibraryResolve (_server.index, _identifier)
	if _error != nil {
		return nil, _error
	}
	if _library == nil {
		return nil, errorw (0xb1852bf9, nil)
	}
	return _library, nil
}

func serverDocumentResolve (_server *Server, _identifierUnsafe string) (*Document, *Error) {
	if _identifierUnsafe == "" {
		return nil, errorw (0xc7f50900, nil)
	}
	_identifier, _, _, _error := DocumentParseIdentifier (_identifierUnsafe)
	if _error != nil {
		return nil, _error
	}
	_document, _error := IndexDocumentResolve (_server.index, _identifier)
	if _error != nil {
		return nil, _error
	}
	if _document == nil {
		return nil, errorw (0x054e7a60, nil)
	}
	return _document, nil
}

func serverDocumentAndLibraryResolve (_server *Server, _identifierUnsafe string) (*Document, *Library, *Error) {
	_document, _error := serverDocumentResolve (_server, _identifierUnsafe)
	if _error != nil {
		return nil, nil, _error
	}
	if _document.Library == "" {
		return _document, nil, nil
	}
	_library, _error := serverLibraryResolve (_server, _document.Library)
	if _error != nil {
		return nil, nil, _error
	}
	return _document, _library, nil
}




func respondWithHtmlTemplate (_response http.ResponseWriter, _template *html_template.Template, _context interface{}) (*Error) {
	_buffer := bytes.NewBuffer (nil)
	if _error := _template.Execute (_buffer, _context); _error != nil {
		return errorw (0xfa7016b8, _error)
	}
	return respondWithHtmlBuffer (_response, _buffer)
}

func respondWithTextTemplate (_response http.ResponseWriter, _template *text_template.Template, _context interface{}) (*Error) {
	_buffer := bytes.NewBuffer (nil)
	if _error := _template.Execute (_buffer, _context); _error != nil {
		return errorw (0xb056a9fb, _error)
	}
	return respondWithBuffer (_response, "text/plain; charset=utf-8", _buffer)
}


func respondWithHtmlString (_response http.ResponseWriter, _body string) (*Error) {
	_buffer := bytes.NewBufferString (_body)
	return respondWithHtmlBuffer (_response, _buffer)
}

func respondWithHtmlBytes (_response http.ResponseWriter, _body []byte) (*Error) {
	_buffer := bytes.NewBuffer (_body)
	return respondWithHtmlBuffer (_response, _buffer)
}

func respondWithHtmlBuffer (_response http.ResponseWriter, _body *bytes.Buffer) (*Error) {
	return respondWithBuffer (_response, "text/html; charset=utf-8", _body)
}


func respondWithBuffer (_response http.ResponseWriter, _contentType string, _body *bytes.Buffer) (*Error) {
	
	_headers := _response.Header ()
	if _contentType == "" {
		return errorw (0xf2b50fc0, nil)
	}
	
	_headers.Add ("Content-Type", _contentType)
	
	_response.WriteHeader (http.StatusOK)
	
	if _, _error := _body.WriteTo (_response); _error != nil {
		return errorw (0xfaf6816b, _error)
	}
	
	return nil
}

