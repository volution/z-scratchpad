

package zscratchpad


import html_template "html/template"
import text_template "text/template"


import embedded "github.com/cipriancraciun/z-scratchpad/embedded"




type Templates struct {
	
	librariesIndexHtml *html_template.Template
	librariesIndexText *text_template.Template
	
	documentsIndexHtml *html_template.Template
	documentsIndexText *text_template.Template
	
	libraryViewHtml *html_template.Template
	libraryViewText *text_template.Template
	
	documentViewHtml *html_template.Template
	documentViewText *text_template.Template
	
	documentExportHtml *html_template.Template
	documentExportText *text_template.Template
	documentExportSource *text_template.Template
	
	versionHtml *html_template.Template
}




func TemplatesNew () (*Templates, *Error) {
	
	_templates := & Templates {}
	
	
	if _template, _error := html_template.New ("") .Parse (embedded.LibrariesIndexHtml); _error == nil {
		_templates.librariesIndexHtml = _template
	} else {
		return nil, errorw (0xddf3e63d, _error)
	}
	
	if _template, _error := text_template.New ("") .Parse (embedded.LibrariesIndexText); _error == nil {
		_templates.librariesIndexText = _template
	} else {
		return nil, errorw (0x058fefb6, _error)
	}
	
	
	if _template, _error := html_template.New ("") .Parse (embedded.DocumentsIndexHtml); _error == nil {
		_templates.documentsIndexHtml = _template
	} else {
		return nil, errorw (0x1a94856e, _error)
	}
	
	if _template, _error := text_template.New ("") .Parse (embedded.DocumentsIndexText); _error == nil {
		_templates.documentsIndexText = _template
	} else {
		return nil, errorw (0x87a616c1, _error)
	}
	
	
	if _template, _error := html_template.New ("") .Parse (embedded.LibraryViewHtml); _error == nil {
		_templates.libraryViewHtml = _template
	} else {
		return nil, errorw (0x8b27f3f3, _error)
	}
	
	if _template, _error := text_template.New ("") .Parse (embedded.LibraryViewText); _error == nil {
		_templates.libraryViewText = _template
	} else {
		return nil, errorw (0x02f74b4c, _error)
	}
	
	
	if _template, _error := html_template.New ("") .Parse (embedded.DocumentViewHtml); _error == nil {
		_templates.documentViewHtml = _template
	} else {
		return nil, errorw (0x1c9ad4b2, _error)
	}
	
	if _template, _error := text_template.New ("") .Parse (embedded.DocumentViewText); _error == nil {
		_templates.documentViewText = _template
	} else {
		return nil, errorw (0x2cf9a5f7, _error)
	}
	
	
	if _template, _error := html_template.New ("") .Parse (embedded.DocumentExportHtml); _error == nil {
		_templates.documentExportHtml = _template
	} else {
		return nil, errorw (0x91a8e5c7, _error)
	}
	
	if _template, _error := text_template.New ("") .Parse (embedded.DocumentExportText); _error == nil {
		_templates.documentExportText = _template
	} else {
		return nil, errorw (0xd15daf3d, _error)
	}
	
	if _template, _error := text_template.New ("") .Parse (embedded.DocumentExportSource); _error == nil {
		_templates.documentExportSource = _template
	} else {
		return nil, errorw (0x01bfce67, _error)
	}
	
	
	if _template, _error := html_template.New ("") .Parse (embedded.VersionHtml); _error == nil {
		_templates.versionHtml = _template
	} else {
		return nil, errorw (0x9ac0bdea, _error)
	}
	
	
	for _, _topTemplate := range []*html_template.Template {
			_templates.librariesIndexHtml,
			_templates.documentsIndexHtml,
			_templates.libraryViewHtml,
			_templates.documentViewHtml,
			_templates.documentExportHtml,
	} {
		if _, _error := _topTemplate.New ("global-navigation") .Parse (embedded.GlobalNavigationHtml); _error != nil {
			return nil, errorw (0xea0e8f0e, _error)
		}
		if _, _error := _topTemplate.New ("global-partials") .Parse (embedded.GlobalPartialsHtml); _error != nil {
			return nil, errorw (0x72af86c7, _error)
		}
	}
	
	return _templates, nil
}

