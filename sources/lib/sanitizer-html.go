

package zscratchpad


import "bytes"
import "encoding/base64"
import "net/url"
import "strings"
import "sort"

import "golang.org/x/net/html"
import "golang.org/x/net/html/atom"
import "github.com/microcosm-cc/bluemonday"




type DocumentSanitizeHtmlOutcome struct {
	Urls map[string]*url.URL
	UrlsLabel map[string][]string
}




func DocumentSanitizeHtml (_document *Document, _unsafe string) (string, *DocumentSanitizeHtmlOutcome, *Error) {
	
	_parser := bluemonday.UGCPolicy()
	
	_parser.RequireParseableURLs (true)
	_parser.RequireNoFollowOnLinks (true)
	_parser.RequireNoReferrerOnLinks (true)
	_parser.RequireCrossOriginAnonymous (true)
	_parser.AllowAttrs("class") .Globally ()
	
	_unsafeBuffer := bytes.NewBufferString (_unsafe)
	
	_sanitizedBuffer := _parser.SanitizeReader (_unsafeBuffer)
	
	_node, _error := html.Parse (_sanitizedBuffer)
	if _error != nil {
		return "", nil, errorw (0x5b1d2f42, _error)
	}
	_node = _node.FirstChild
	_node = _node.FirstChild.NextSibling
	if (_node.Type != html.ElementNode) || (_node.DataAtom != atom.Body) {
		return "", nil, errorw (0x875a6da3, nil)
	}
	
	_extractLinksContext := & extractLinksContext {
			urlsParsed : make (map[string]*url.URL, 1024),
			urlsLabel : make (map[string][]string, 1024),
		}
	
	if _error := extractLinks (_node, _extractLinksContext); _error != nil {
		return "", nil, _error
	}
	
	if _error := verifyAnchors (_node); _error != nil {
		return "", nil, _error
	}
	
	_outcome := & DocumentSanitizeHtmlOutcome {
			Urls : _extractLinksContext.urlsParsed,
			UrlsLabel : _extractLinksContext.urlsLabel,
		}
	
	_mangledBuffer := BytesBufferNewSize (128 * 1024)
	defer BytesBufferRelease (_mangledBuffer)
	for _child := _node.FirstChild; _child != nil; _child = _child.NextSibling {
		if _error := html.Render (_mangledBuffer, _child); _error != nil {
			return "", nil, errorw (0xba050bcf, _error)
		}
	}
	
	_mangled := string (_mangledBuffer.Bytes ())
	
	return _mangled, _outcome, nil
}




func DocumentSanitizeUrl (_url *url.URL) (*Error) {
	
	_url.Scheme = strings.ToLower (_url.Scheme)
	_url.Host = strings.ToLower (_url.Host)
	
	if (_url.Scheme == "") && (_url.Host != "") {
		_url.Scheme = "http"
	}
	if (_url.Host == "") {
		_url.User = nil
	}
	if (_url.Scheme != "") && (_url.Opaque != "") {
		return nil
	}
	if _url.Opaque != "" {
		return errorw (0x8b6ea4c4, nil)
	}
	if (_url.Scheme == "") && (_url.Host == "") && (_url.Path == "") && (_url.RawQuery == "") {
		return nil
	}
	if _url.Host != "" {
		return nil
	}
	if _url.Path == "" {
		return errorw (0x1a8fa951, nil)
	}
	_path := _url.Path
	if _path == "/" {
		return nil
	}
	if strings.HasPrefix (_path, ".") {
		logf ('e', 0x316d2c87, "`%s`", _path)
		_urlString := _url.String ()
		*_url = url.URL {
				Path : "/ue/" + base64.RawURLEncoding.EncodeToString ([]byte (_urlString)),
			}
		return nil
	}
	if strings.HasPrefix (_path, "/d/") {
		// _identifier := _path[4:]
		// FIXME: ...
		return nil
	}
	if strings.HasPrefix (_path, "/l/") {
		// _identifier := _path[4:]
		// FIXME: ...
		return nil
	}
	if strings.HasPrefix (_path, "/i/") {
		// _identifier := _path[4:]
		// FIXME: ...
		return nil
	}
	return errorw (0x13ced7ac, nil)
}




func extractLinks (_node *html.Node, _context *extractLinksContext) (*Error) {
	
	_mangleAttribute := func (_node *html.Node, _urlAttribute string, _labelAttribute string, _action bool) (*Error) {
		
		_urlUnsafe := ""
		_urlFound := false
		_label := ""
		_labelFromTitle := ""
		
		for _, _attribute := range _node.Attr {
			switch _attribute.Key {
				case _urlAttribute :
					_urlUnsafe = _attribute.Val
					_urlFound = true
				case _labelAttribute :
					if _labelAttribute != "" {
						_label = _attribute.Val
					}
				case "title" :
					_labelFromTitle = _attribute.Val
			}
		}
		
		if !_urlFound {
			return nil
		}
		
		_urlUnsafe = strings.TrimSpace (_urlUnsafe)
		_label = strings.TrimSpace (_label)
		_labelFromTitle = strings.TrimSpace (_labelFromTitle)
		
		
		if _urlUnsafe == "" {
			_urlUnsafe = "/ue/"
		}
		
		if _label == "" {
			_label = _labelFromTitle
		}
		if _label == "" {
			if (_node.DataAtom == atom.A) && (_node.FirstChild != nil) && (_node.FirstChild.Type == html.TextNode) {
				_label = strings.TrimSpace (_node.FirstChild.Data)
			}
		}
		
		_url := (*url.URL) (nil)
		
		if _url_0, _error := url.Parse (_urlUnsafe); _error == nil {
			_url = _url_0
		} else {
			_error := errorw (0x665bacf5, _error)
			logErrorf ('e', 0x7c526e12, _error, "`%s`", _urlUnsafe)
			_urlUnsafe = "/ue/"
			_url = & url.URL { Path : "/ue/" }
		}
		
		if _error := DocumentSanitizeUrl (_url); _error != nil {
			logErrorf ('e', 0x4e98912e, _error, "`%s`", _urlUnsafe)
			_urlUnsafe = "/ue/"
			_url = & url.URL { Path : "/ue/" }
		}
		
		_urlString := _url.String ()
		_urlType := ""
		
		if strings.HasPrefix (_urlString, "/ue/") {
			_urlType = "error"
		} else if strings.HasPrefix (_urlString, "/") {
			_urlType = "internal"
		} else if strings.HasPrefix (_urlString, "#") {
			_urlType = "anchor"
		} else {
			_urlType = "external"
		}
		
		_urlOpenString := _urlString
		if _urlType == "external" {
			_urlOpenString = (& url.URL { Path : "/uo/" + base64.RawURLEncoding.EncodeToString ([]byte (_urlOpenString)) }) .String ()
		}
		
		_urlUseString := _urlString
		if _action {
			_urlUseString = _urlOpenString
		}
		
		for _index, _attribute := range _node.Attr {
			switch _attribute.Key {
				case _urlAttribute :
					_node.Attr[_index].Val = _urlUseString
			}
		}
		
		_node.Attr = append (_node.Attr, html.Attribute { "", "data-zs-url-original-" + _urlAttribute, _urlUnsafe })
		_node.Attr = append (_node.Attr, html.Attribute { "", "data-zs-url-type", _urlType })
		
		// NOTE:  From here on we just index it.
		
		_url.User = nil
		_url.Fragment = ""
		_url.RawFragment = ""
		_urlString = _url.String ()
		
		if _urlString == "" {
			return nil
		}
		
		_urlLabels := []string (nil)
		if _, _exists := _context.urlsParsed[_urlString]; _exists {
			_urlLabels = _context.urlsLabel[_urlString]
		} else {
			_context.urlsParsed[_urlString] = _url
			_urlLabels = make ([]string, 0, 16)
		}
		if _label != "" {
			_urlLabels = append (_urlLabels, _label)
			sort.Strings (_urlLabels)
		}
		_context.urlsLabel[_urlString] = _urlLabels
		
		return nil
	}
	
	if _node.Type == html.ElementNode {
		_error := (*Error) (nil)
		switch _node.DataAtom {
			case atom.A :
				_error = _mangleAttribute (_node, "href", "", true)
			case atom.Area :
				_error = _mangleAttribute (_node, "href", "alt", true)
			case atom.Blockquote, atom.Q, atom.Ins, atom.Del :
				_error = _mangleAttribute (_node, "cite", "", true)
			case atom.Img :
				_error = _mangleAttribute (_node, "src", "alt", false)
		}
		if _error != nil {
			return _error
		}
	}
	
	for _child := _node.FirstChild; _child != nil; _child = _child.NextSibling {
		if _error := extractLinks (_child, _context); _error != nil {
			return _error
		}
	}
	
	return nil
}


type extractLinksContext struct {
	urlsParsed map[string]*url.URL
	urlsLabel map[string][]string
}




func verifyAnchors (_node *html.Node) (*Error) {
	
	_anchors := make (map[string]bool, 1024)
	
	if _error := collectAnchors (_node, _anchors); _error != nil {
		return _error
	}
	if _error := validateAnchors (_node, _anchors); _error != nil {
		return _error
	}
	
	return nil
}


func collectAnchors (_node *html.Node, _anchors map[string]bool) (*Error) {
	
	_collectAttribute := func (_node *html.Node, _anchorAttribute string) (*Error) {
		
		_anchor := ""
		
		// FIXME:  Handle duplicate attributes!
		for _, _attribute := range _node.Attr {
			if _attribute.Key == _anchorAttribute {
				_anchor = _attribute.Val
			}
		}
		
		_anchor = strings.TrimSpace (_anchor)
		if _anchor == "" {
			return nil
		}
		
		if _anchor == "" {
			return nil
		}
		
		if _, _exists := _anchors[_anchor]; _exists {
			logf ('e', 0x93a6fd7a, "`%s`", _anchor)
			return nil
		}
		
		_anchors[_anchor] = true
		
		return nil
	}
	
	if _node.Type == html.ElementNode {
		if _error := _collectAttribute (_node, "id"); _error != nil {
			return _error
		}
	}
	
	for _child := _node.FirstChild; _child != nil; _child = _child.NextSibling {
		if _error := collectAnchors (_child, _anchors); _error != nil {
			return _error
		}
	}
	
	return nil
}


func validateAnchors (_node *html.Node, _anchors map[string]bool) (*Error) {
	
	_validateAttribute := func (_node *html.Node, _anchorAttribute string) (*Error) {
		
		_anchor := ""
		
		// FIXME:  Handle duplicate attributes!
		for _, _attribute := range _node.Attr {
			if _attribute.Key == _anchorAttribute {
				_anchor = _attribute.Val
			}
		}
		
		_anchor = strings.TrimSpace (_anchor)
		if _anchor == "" {
			return nil
		}
		
		if ! strings.HasPrefix (_anchor, "#") {
			return nil
		}
		
		_anchor = _anchor[1:]
		if _anchor == "" {
			return nil
		}
		
		if _, _exists := _anchors[_anchor]; _exists {
			return nil
		}
		
		logf ('e', 0xf5b51a6e, "`%s`", _anchor)
		
		for _index, _attribute := range _node.Attr {
			if _attribute.Key == "data-zs-url-type" {
				_node.Attr[_index].Val = "error"
				return nil
			}
		}
		
		_node.Attr = append (_node.Attr, html.Attribute { "", "data-zs-url-type", "error" })
		return nil
	}
	
	if _node.Type == html.ElementNode {
		_error := (*Error) (nil)
		switch _node.DataAtom {
			case atom.A :
				_error = _validateAttribute (_node, "href")
		}
		if _error != nil {
			return _error
		}
	}
	
	for _child := _node.FirstChild; _child != nil; _child = _child.NextSibling {
		if _error := validateAnchors (_child, _anchors); _error != nil {
			return _error
		}
	}
	
	return nil
}

