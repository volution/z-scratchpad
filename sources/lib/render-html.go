

package zscratchpad


import "bytes"
import "html"


import html_template "html/template"




func DocumentRenderToHtmlDocument (_document *Document, _export bool, _theme string, _templates *Templates, _buffer *bytes.Buffer) (*Error) {
	_documentHtml, _error := DocumentRenderToHtml (_document, _export)
	if _error != nil {
		return _error
	}
	_themeCssAsset := ""
	if _theme == "default" {
		_theme = "github"
	}
	switch _theme {
		case "plain" :
			_themeCssAsset = "assets/css-export/plain.css"
		case "github" :
			_themeCssAsset = "assets/css-export/github-min.css"
		case "modest" :
			_themeCssAsset = "assets/css-export/modest-min.css"
		default :
			return errorw (0x922a8ee1, nil)
	}
	_, _themeCssData, _error := TemplatesAssetResolve (_templates, _themeCssAsset)
	if _error != nil {
		return _error
	}
	_context := struct {
			Document *Document
			DocumentHtml html_template.HTML
			ThemeCss html_template.CSS
		} {
			_document,
			html_template.HTML (_documentHtml),
			html_template.CSS (_themeCssData),
		}
	if _error := _templates.documentExportHtmlDocument.Execute (_buffer, _context); _error != nil {
		return errorw (0xf6bb6151, _error)
	}
	return nil
}




func DocumentRenderToHtml (_document *Document, _export bool) (string, *Error) {
	
	if _export {
		if _document.RenderHtmlExport != "" {
			return _document.RenderHtmlExport, nil
		}
	} else {
		if _document.RenderHtml != "" {
			return _document.RenderHtml, nil
		}
	}
	
	_format := _document.Format
	if _format == "" {
		_format = "text"
		// return "", errorf (0xaff80238, "format empty")
	}
	
	_render := ""
	_error := (*Error) (nil)
	
	switch _format {
		
		case "text" :
			_render, _error = documentRenderTextToHtml (_document.BodyLines)
		
		case "snippets" :
			_render, _error = documentRenderSnippetsToHtml (_document.BodyLines)
		
		case "commonmark" :
			_render, _error = documentRenderCommonmarkToHtml (_document.BodyLines)
		
		case "gemini" :
			_render, _error = documentRenderGeminiToHtml (_document.BodyLines)
		
		default :
			return "", errorf (0xaf60ea6d, "format invalid `%s`", _document.Format)
	}
	
	if _error != nil {
		return "", _error
	}
	
	_render, _outcome, _error := DocumentSanitizeHtml (_document, _render, !_export)
	if _error != nil {
		return "", _error
	}
	
	if _export {
		_document.RenderHtmlExport = _render
	} else {
		_document.RenderHtml = _render
		_document.HtmlLinks = _outcome.UrlsLabel
	}
	
	return _render, nil
}




func documentRenderCommonmarkToHtml (_source []string) (string, *Error) {
	return parseAndRenderCommonmarkToHtml (_source)
}

func documentRenderGeminiToHtml (_source []string) (string, *Error) {
	return parseAndRenderGeminiToHtml (_source)
}

func documentRenderSnippetsToHtml (_source []string) (string, *Error) {
	return parseAndRenderSnippetsToHtml (_source)
}




func documentRenderTextToHtml (_source []string) (string, *Error) {
	
	_buffer := BytesBufferNewSize (128 * 1024)
	defer BytesBufferRelease (_buffer)
	
	_buffer.WriteString ("<pre><code>")
	for _, _line := range _source {
		_line = html.EscapeString (_line)
		_buffer.WriteString (_line)
		_buffer.WriteString ("\n")
	}
	_buffer.WriteString ("</code></pre>\n")
	
	_output := string (_buffer.Bytes ())
	
	return _output, nil
}

