

package zscratchpad


import "bytes"
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
	
	_outcome := & DocumentSanitizeHtmlOutcome {
			Urls : _extractLinksContext.urlsParsed,
			UrlsLabel : _extractLinksContext.urlsLabel,
		}
	
	_mangledBuffer := bytes.NewBuffer (nil)
	for _child := _node.FirstChild; _child != nil; _child = _child.NextSibling {
		if _error := html.Render (_mangledBuffer, _child); _error != nil {
			return "", nil, errorw (0xba050bcf, _error)
		}
	}
	
	_mangled := _mangledBuffer.String ()
	
	return _mangled, _outcome, nil
}




func extractLinks (_node *html.Node, _context *extractLinksContext) (*Error) {
	
	if _node.Type == html.ElementNode {
		if _node.DataAtom == atom.A {
			
			_url := (*url.URL) (nil)
			_title := ""
			for _, _attribute := range _node.Attr {
				switch _attribute.Key {
					case "href" :
						if _url_0, _error := url.Parse (_attribute.Val); _error == nil {
							_url = _url_0
						} else {
							return errorw (0xb6f30c85, nil)
						}
					case "title" :
						_title = strings.TrimSpace (_attribute.Val)
				}
			}
			
			if _url != nil {
				if _title == "" {
					if (_node.FirstChild != nil) && (_node.FirstChild.Type == html.TextNode) {
						_title = strings.TrimSpace (_node.FirstChild.Data)
					}
				}
			}
			
			_urlString := ""
			if _url != nil {
				
				_url.User = nil
				_url.Fragment = ""
				_url.RawFragment = ""
				_urlString = _url.String ()
				
				if _urlString == "" {
					_url = nil
				}
			}
			
			if _url != nil {
				_urlLabels := []string (nil)
				if _, _exists := _context.urlsParsed[_urlString]; _exists {
					_urlLabels = _context.urlsLabel[_urlString]
				} else {
					_context.urlsParsed[_urlString] = _url
					_urlLabels = make ([]string, 0, 16)
				}
				if _title != "" {
					_urlLabels = append (_urlLabels, _title)
					sort.Strings (_urlLabels)
				}
				_context.urlsLabel[_urlString] = _urlLabels
			}
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

