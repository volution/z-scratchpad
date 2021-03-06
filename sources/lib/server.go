

package zscratchpad


import "bytes"
import "encoding/base64"
import "encoding/xml"
import "io"
import "net"
import "net/http"
import "net/url"
import "strings"

import html_template "html/template"
import text_template "text/template"


import . "github.com/volution/z-scratchpad/embedded"




type Server struct {
	
	index *Index
	editor *Editor
	browser *Browser
	templates *Templates
	listener net.Listener
	http *http.Server
	globals *Globals
	
	EditEnabled bool
	CreateEnabled bool
	BrowseEnabled bool
	ClipboardEnabled bool
	
	OpenExternalConfirm bool
	OpenExternalConfirmSkipForSchemas []string
	
	UrlBase string
	AuthenticationCookieName string
	AuthenticationCookieTimeout uint
	AuthenticationCookieSecret string
	
	reloadToken string
	
}


type serverHandler struct {
	server *Server
}




func ServerNew (_globals *Globals, _index *Index, _editor *Editor, _browser *Browser, _listener net.Listener) (*Server, *Error) {
	
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
			browser : _browser,
			templates : _templates,
			listener : _listener,
			reloadToken : generateRandomToken (),
		}
	
	_server.EditEnabled = true
	_server.CreateEnabled = true
	_server.BrowseEnabled = true
	_server.ClipboardEnabled = true
	
	_server.AuthenticationCookieName = "zscratchpad_authentication"
	_server.AuthenticationCookieTimeout = 28 * 24 * 3600
	_server.AuthenticationCookieSecret = generateRandomToken ()
	
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
	
	_server.globals.MutexLock ()
	defer _server.globals.MutexUnlock ()
	
	if _request.Method != "GET" {
		return errorw (0x7f32157c, nil)
	}
	
	_path := _request.URL.Path
	if ! strings.HasPrefix (_path, "/") {
		return errorw (0x828c5f04, nil)
	}
	
	if _path == "/__/heartbeat" {
		return respondWithTextString (_response, "OK\n")
	}
	
	_setAuthenticationCookie := func (_server *Server, _response http.ResponseWriter) (*Error) {
			_mac, _error := generateHmac (_server.AuthenticationCookieSecret, "/__/authenticate/{cookie}")
			if _error != nil {
				return _error
			}
			_cookie := & http.Cookie {
					Name : _server.AuthenticationCookieName,
					Value : _mac,
					Path : "/",
					MaxAge : int (_server.AuthenticationCookieTimeout),
					HttpOnly : true,
					SameSite : http.SameSiteStrictMode,
				}
			http.SetCookie (_response, _cookie)
			return nil
		}
	
	if _path == "/__/authenticate" {
		_mac, _error := generateHmac (_server.AuthenticationCookieSecret, "/__/authenticate/{token}")
		if _error != nil {
			return _error
		}
		logf ('i', 0x04cc15c3, "[server]  authentication token: `%s`;", _mac)
		logf ('i', 0x2a34dc7a, "[server]  authentication URL: `%s/__/authenticate/%s`;", strings.TrimRight (_server.UrlBase, "/"), _mac)
		return respondWithTextString (_response, "NOK")
	}
	if strings.HasPrefix (_path, "/__/authenticate/") {
		_mac := _path[17:]
		if _error := verifyHmac (_server.AuthenticationCookieSecret, "/__/authenticate/{token}", _mac, 60 * 1000); _error != nil {
			return _error
		}
		if _error := _setAuthenticationCookie (_server, _response); _error != nil {
			return _error
		}
		return respondWithTextString (_response, "OK")
	}
	if _path == "/__/deauthenticate" {
		_cookie := & http.Cookie {
				Name : _server.AuthenticationCookieName,
				Value : "",
				Path : "/",
				MaxAge : int (_server.AuthenticationCookieTimeout),
				HttpOnly : true,
				SameSite : http.SameSiteStrictMode,
			}
		http.SetCookie (_response, _cookie)
		return respondWithTextString (_response, "OK")
	}
	
	if _tokens, _ := _request.URL.Query () ["authenticate"]; _tokens != nil {
		if len (_tokens) != 1 {
			return errorw (0xc54dcef5, nil)
		}
		if _error := verifyHmac (_server.AuthenticationCookieSecret, "/__/authenticate/{query}", _tokens[0], 6 * 1000); _error != nil {
			return _error
		}
		if _error := _setAuthenticationCookie (_server, _response); _error != nil {
			return _error
		}
		return respondWithRedirect (_response, _path)
	} else if _cookie, _error := _request.Cookie (_server.AuthenticationCookieName); _error == nil {
		if _error := verifyHmac (_server.AuthenticationCookieSecret, "/__/authenticate/{cookie}", _cookie.Value, _server.AuthenticationCookieTimeout * 1000); _error != nil {
			return _error
		}
		if _error := _setAuthenticationCookie (_server, _response); _error != nil {
			return _error
		}
	} else {
		return errorw (0xcf851b50, _error)
	}
	
	if _path == "/__/reload" {
		return respondWithTextString (_response, _server.reloadToken)
	}
	
	if _path == "/" {
		return ServerHandleHome (_server, _response)
	}
	if (_path == "/i") || (_path == "/i/") {
		return ServerHandleIndex (_server, _response)
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
	if strings.HasPrefix (_path, "/df/") {
		_identifier := _path[4:]
		return ServerHandleDocumentFingerprint (_server, _identifier, _response)
	}
	if strings.HasPrefix (_path, "/dx/html-body/") {
		_identifier := _path[14:]
		return ServerHandleDocumentExportHtml (_server, _identifier, _response)
	}
	if strings.HasPrefix (_path, "/dx/html-document/") {
		_identifier := _path[18:]
		return ServerHandleDocumentExportHtmlDocument (_server, _identifier, "default", _response)
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
	
	if strings.HasPrefix (_path, "/ul/") {
		_url := _path[4:]
		return ServerHandleUrlLaunch (_server, _url, _response)
	}
	if strings.HasPrefix (_path, "/uo/") {
		_url := _path[4:]
		return ServerHandleUrlOpen (_server, _url, _response)
	}
	if strings.HasPrefix (_path, "/ue/") {
		_url := _path[4:]
		return ServerHandleUrlError (_server, _url, _response)
	}
	
	if strings.HasPrefix (_path, "/cs/") {
		_data := _path[4:]
		_data_0, _error := base64.RawURLEncoding.DecodeString (_data)
		if _error != nil {
			return errorw (0x0cedd6db, _error)
		}
		_data = string (_data_0)
		_data, _error = url.PathUnescape (_data)
		if _error != nil {
			return errorw (0x59d3edfb, _error)
		}
		return ServerHandleClipboardStore (_server, _data, _response)
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
	if _path == "/__/manual.txt" {
		return ServerHandleManualText (_server, _response)
	}
	if _path == "/__/manual.html" {
		return ServerHandleManualHtml (_server, _response)
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




func ServerHandleHome (_server *Server, _response http.ResponseWriter) (*Error) {
	_context := struct {
			Server *Server
		} {
			_server,
		}
	return respondWithHtmlTemplate (_response, _server.templates.homeHtml, _context, true)
}




func ServerHandleIndex (_server *Server, _response http.ResponseWriter) (*Error) {
	_libraries, _error := IndexLibrariesSelectAll (_server.index)
	if _error != nil {
		return _error
	}
	_documents, _error := IndexDocumentsSelectAll (_server.index)
	if _error != nil {
		return _error
	}
	_context := struct {
			Server *Server
			Libraries []*Library
			Documents []*Document
		} {
			_server,
			_libraries,
			_documents,
		}
	return respondWithHtmlTemplate (_response, _server.templates.indexHtml, _context, true)
}




func ServerHandleLibrariesIndex (_server *Server, _response http.ResponseWriter) (*Error) {
	_libraries, _error := IndexLibrariesSelectAll (_server.index)
	if _error != nil {
		return _error
	}
	_context := struct {
			Server *Server
			Libraries []*Library
		} {
			_server,
			_libraries,
		}
	return respondWithHtmlTemplate (_response, _server.templates.librariesIndexHtml, _context, true)
}


func ServerHandleDocumentsIndex (_server *Server, _response http.ResponseWriter) (*Error) {
	_documents, _error := IndexDocumentsSelectAll (_server.index)
	if _error != nil {
		return _error
	}
	_context := struct {
			Server *Server
			Documents []*Document
		} {
			_server,
			_documents,
		}
	return respondWithHtmlTemplate (_response, _server.templates.documentsIndexHtml, _context, true)
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
			Server *Server
			Library *Library
			Documents []*Document
		} {
			_server,
			_library,
			_documents,
		}
	return respondWithHtmlTemplate (_response, _server.templates.libraryViewHtml, _context, true)
}




func ServerHandleDocumentView (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	_document, _library, _error := serverDocumentAndLibraryResolve (_server, _identifierUnsafe)
	if _error != nil {
		return _error
	}
	_documentHtml, _error := DocumentRenderToHtml (_document, false)
	if _error != nil {
		return _error
	}
	_context := struct {
			Server *Server
			Library *Library
			Document *Document
			DocumentHtml html_template.HTML
		} {
			_server,
			_library,
			_document,
			html_template.HTML (_documentHtml),
		}
	return respondWithHtmlTemplate (_response, _server.templates.documentViewHtml, _context, true)
}


func ServerHandleDocumentFingerprint (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	_document, _error := serverDocumentResolve (_server, _identifierUnsafe)
	if _error != nil {
		return _error
	}
	return respondWithTextString (_response, _document.SourceFingerprint)
}


func ServerHandleDocumentExportHtml (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	_document, _error := serverDocumentResolve (_server, _identifierUnsafe)
	if _error != nil {
		return _error
	}
	_documentHtml, _error := DocumentRenderToHtml (_document, true)
	if _error != nil {
		return _error
	}
	_context := struct {
			Server *Server
			Document *Document
			DocumentHtml html_template.HTML
		} {
			_server,
			_document,
			html_template.HTML (_documentHtml),
		}
	return respondWithHtmlTemplate (_response, _server.templates.documentExportHtml, _context, true)
}


func ServerHandleDocumentExportHtmlDocument (_server *Server, _identifierUnsafe string, _theme string, _response http.ResponseWriter) (*Error) {
	_document, _error := serverDocumentResolve (_server, _identifierUnsafe)
	if _error != nil {
		return _error
	}
	_buffer := BytesBufferNewSize (128 * 1024)
	defer BytesBufferRelease (_buffer)
	if _error := DocumentRenderToHtmlDocument (_document, true, _theme, _server.templates, _buffer); _error != nil {
		return _error
	}
	return respondWithHtmlBuffer (_response, _buffer)
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
			Server *Server
			Document *Document
			DocumentText string
		} {
			_server,
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
			Server *Server
			Document *Document
			DocumentSource string
		} {
			_server,
			_document,
			_documentSource,
		}
	return respondWithTextTemplate (_response, _server.templates.documentExportSource, _context)
}




func ServerHandleDocumentCreate (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	if !_server.CreateEnabled {
		return errorw (0x744d1a48, nil)
	}
	if _server.editor == nil {
		return errorw (0x14317f29, nil)
	}
	if _error := WorkflowDocumentCreate (_identifierUnsafe, _server.index, _server.editor, false); _error != nil {
		return _error
	}
	http.Error (_response, "", http.StatusNoContent)
	return nil
}


func ServerHandleDocumentEdit (_server *Server, _identifierUnsafe string, _response http.ResponseWriter) (*Error) {
	if !_server.EditEnabled {
		return errorw (0x664c252f, nil)
	}
	if _server.editor == nil {
		return errorw (0xee28afb6, nil)
	}
	if _error := WorkflowDocumentEdit (_identifierUnsafe, _server.index, _server.editor, false); _error != nil {
		return _error
	}
	http.Error (_response, "", http.StatusNoContent)
	return nil
}




func ServerHandleUrlLaunch (_server *Server, _urlEncoded string, _response http.ResponseWriter) (*Error) {
	// FIXME:  We should add some type of signature so that we aren't injected malicious URL's!
	// FIXME:  We should make sure this is via a `POST` request!
	_urlLaunch_0, _error := base64.RawURLEncoding.DecodeString (_urlEncoded)
	if _error != nil {
		return errorw (0x06ca25ef, _error)
	}
	_urlLaunch := string (_urlLaunch_0)
	if !_server.BrowseEnabled {
		return errorw (0xcbe5ac01, nil)
	}
	if _server.browser == nil {
		return errorw (0x13f43f95, nil)
	}
	if _error := BrowserUrlExternalOpen (_server.browser, _urlLaunch, false); _error != nil {
		return _error
	}
	http.Error (_response, "", http.StatusNoContent)
	return nil
}


func ServerHandleUrlOpen (_server *Server, _urlEncoded string, _response http.ResponseWriter) (*Error) {
	// FIXME:  We should add some type of signature so that we aren't injected malicious URL's!
	_urlOpen_0, _error := base64.RawURLEncoding.DecodeString (_urlEncoded)
	_urlOpen := string (_urlOpen_0)
	if _error != nil {
		return errorw (0x34d08c61, _error)
	}
	if !_server.OpenExternalConfirm {
		return ServerHandleUrlLaunch (_server, _urlEncoded, _response)
	} else {
		for _, _schema := range _server.OpenExternalConfirmSkipForSchemas {
			if strings.HasPrefix (_urlOpen, _schema + ":") {
				return ServerHandleUrlLaunch (_server, _urlEncoded, _response)
			}
		}
	}
	_context := struct {
			Server *Server
			UrlEncoded string
			UrlOpen html_template.URL
		} {
			_server,
			_urlEncoded,
			html_template.URL (_urlOpen),
		}
	return respondWithHtmlTemplate (_response, _server.templates.urlOpenHtml, _context, true)
}


func ServerHandleUrlError (_server *Server, _urlEncoded string, _response http.ResponseWriter) (*Error) {
	_urlUnsafe_0, _error := base64.RawURLEncoding.DecodeString (_urlEncoded)
	if _error != nil {
		return errorw (0x33ccce60, _error)
	}
	_urlUnsafe := string (_urlUnsafe_0)
	_context := struct {
			Server *Server
			UrlEncoded string
			UrlUnsafe string
		} {
			_server,
			_urlEncoded,
			_urlUnsafe,
		}
	return respondWithHtmlTemplate (_response, _server.templates.urlErrorHtml, _context, true)
}




func ServerHandleClipboardStore (_server *Server, _data string, _response http.ResponseWriter) (*Error) {
	if !_server.ClipboardEnabled {
		return errorw (0x7569419a, nil)
	}
	if _server.editor == nil {
		return errorw (0x8fae00cd, nil)
	}
	if _error := EditorClipboardStore (_server.editor, _data); _error != nil {
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
	return respondWithHtmlTemplate (_response, _server.templates.versionHtml, _context, true)
}




func ServerHandleSourcesMd5 (_server *Server, _response http.ResponseWriter) (*Error) {
	return respondWithTextString (_response, BuildSourcesMd5)
}

func ServerHandleSourcesCpio (_server *Server, _response http.ResponseWriter) (*Error) {
	_response.Header () .Add ("Content-Encoding", "gzip")
	return respondWithBuffer (_response, "application/x-cpio", bytes.NewBuffer (BuildSourcesCpioGz))
}


func ServerHandleManualText (_server *Server, _response http.ResponseWriter) (*Error) {
	return respondWithTextString (_response, ZscratchpadManualTxt)
}

func ServerHandleManualHtml (_server *Server, _response http.ResponseWriter) (*Error) {
	return respondWithHtmlString (_response, ZscratchpadManualHtml)
}




func ServerHandleAsset (_server *Server, _path string, _response http.ResponseWriter) (*Error) {
	_contentType, _body, _error := TemplatesAssetResolve (_server.templates, _path)
	if _error != nil {
		return _error
	}
	return respondWithBuffer (_response, _contentType, bytes.NewBuffer (_body))
}




func serverLibraryResolve (_server *Server, _identifierUnsafe string) (*Library, *Error) {
	return WorkflowLibraryResolve (_identifierUnsafe, _server.index)
}

func serverDocumentResolve (_server *Server, _identifierUnsafe string) (*Document, *Error) {
	return WorkflowDocumentResolve (_identifierUnsafe, _server.index)
}

func serverDocumentAndLibraryResolve (_server *Server, _identifierUnsafe string) (*Document, *Library, *Error) {
	return WorkflowDocumentAndLibraryResolve (_identifierUnsafe, _server.index)
}




func respondWithHtmlTemplate (_response http.ResponseWriter, _template *html_template.Template, _context interface{}, _perhapsStrict bool) (*Error) {
	
	_buffer := bytes.NewBuffer (nil)
	if _error := _template.Execute (_buffer, _context); _error != nil {
		return errorw (0xfa7016b8, _error)
	}
	
	if _perhapsStrict && BUILD_DEVELOPMENT {
		_buffer := bytes.NewReader (_buffer.Bytes ())
		_decoder := xml.NewDecoder (_buffer)
		_decoder.Entity = xml.HTMLEntity
		if _token, _error := _decoder.Token (); _error == nil {
			if _directive, _ok := _token.(xml.Directive); _ok {
				if ! bytes.Equal (_directive, []byte ("doctype html")) {
					return errorw (0x4ca7b74e, nil)
				}
			} else {
				return errorw (0x4b180ca6, nil)
			}
		} else {
			return errorw (0x4fcdd2d3, _error)
		}
		_loopStart : for {
			_token, _error := _decoder.Token ()
			if _error != nil {
				return errorw (0x09480384, _error)
			}
			switch _token := _token.(type) {
				case xml.StartElement :
					if _token.Name.Local != "html" {
						return errorw (0x8adce50b, nil)
					}
					break _loopStart
				case xml.CharData :
					if len (bytes.TrimSpace (_token)) != 0 {
						return errorw (0x5be3689a, nil)
					}
				case xml.Comment :
					// NOP
				default :
					return errorw (0x603bc00b, nil)
			}
		}
		if _error := _decoder.Skip (); _error != nil {
			return errorw (0xaf8cf302, _error)
		}
		_loopEnd : for {
			_token, _error := _decoder.Token ()
			if _error == io.EOF {
				break _loopEnd
			}
			if _error != nil {
				return errorw (0x096981bf, _error)
			}
			switch _token := _token.(type) {
				case xml.CharData :
					if len (bytes.TrimSpace (_token)) != 0 {
						return errorw (0x45153a1a, nil)
					}
				case xml.Comment :
					// NOP
				default :
					return errorw (0x0375afba, nil)
			}
		}
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


func respondWithTextString (_response http.ResponseWriter, _body string) (*Error) {
	_buffer := bytes.NewBufferString (_body)
	return respondWithTextBuffer (_response, _buffer)
}

func respondWithTextBytes (_response http.ResponseWriter, _body []byte) (*Error) {
	_buffer := bytes.NewBuffer (_body)
	return respondWithTextBuffer (_response, _buffer)
}

func respondWithTextBuffer (_response http.ResponseWriter, _body *bytes.Buffer) (*Error) {
	return respondWithBuffer (_response, "text/plain; charset=utf-8", _body)
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


func respondWithRedirect (_response http.ResponseWriter, _url string) (*Error) {
	
	_headers := _response.Header ()
	
	_headers.Add ("Location", _url)
	
	_response.WriteHeader (http.StatusTemporaryRedirect)
	
	return nil
}

