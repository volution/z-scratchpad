

package zscratchpad


import "io/fs"
import "path"

import html_template "html/template"
import text_template "text/template"


import embedded "github.com/cipriancraciun/z-scratchpad/embedded"




type Templates struct {
	
	homeHtml *html_template.Template
	homeText *text_template.Template
	
	indexHtml *html_template.Template
	indexText *text_template.Template
	
	librariesIndexHtml *html_template.Template
	librariesIndexText *text_template.Template
	
	documentsIndexHtml *html_template.Template
	documentsIndexText *text_template.Template
	
	libraryViewHtml *html_template.Template
	libraryViewText *text_template.Template
	
	documentViewHtml *html_template.Template
	documentViewText *text_template.Template
	
	documentExportHtml *html_template.Template
	documentExportHtmlGithub *html_template.Template
	documentExportText *text_template.Template
	documentExportSource *text_template.Template
	
	urlOpenHtml *html_template.Template
	urlErrorHtml *html_template.Template
	
	versionHtml *html_template.Template
	
	assets fs.FS
}




func TemplatesNew () (*Templates, *Error) {
	
	_templates := & Templates {
			assets : embedded.Assets,
		}
	
	
	if _template, _error := html_template.New ("") .Parse (embedded.HomeHtml); _error == nil {
		_templates.homeHtml = _template
	} else {
		return nil, errorw (0x2f22cfcf, _error)
	}
	
	if _template, _error := text_template.New ("") .Parse (embedded.HomeText); _error == nil {
		_templates.homeText = _template
	} else {
		return nil, errorw (0xc165205e, _error)
	}
	
	
	if _template, _error := html_template.New ("") .Parse (embedded.IndexHtml); _error == nil {
		_templates.indexHtml = _template
	} else {
		return nil, errorw (0x0956c86c, _error)
	}
	
	if _template, _error := text_template.New ("") .Parse (embedded.IndexText); _error == nil {
		_templates.indexText = _template
	} else {
		return nil, errorw (0x53d38d91, _error)
	}
	
	
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
	
	if _template, _error := html_template.New ("") .Parse (embedded.DocumentExportHtmlGithub); _error == nil {
		_templates.documentExportHtmlGithub = _template
	} else {
		return nil, errorw (0xc10eb414, _error)
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
	
	
	if _template, _error := html_template.New ("") .Parse (embedded.UrlOpenHtml); _error == nil {
		_templates.urlOpenHtml = _template
	} else {
		return nil, errorw (0x5f69a2f1, _error)
	}
	
	if _template, _error := html_template.New ("") .Parse (embedded.UrlErrorHtml); _error == nil {
		_templates.urlErrorHtml = _template
	} else {
		return nil, errorw (0xe1f336a8, _error)
	}
	
	
	if _template, _error := html_template.New ("") .Parse (embedded.VersionHtml); _error == nil {
		_templates.versionHtml = _template
	} else {
		return nil, errorw (0x9ac0bdea, _error)
	}
	
	
	for _, _topTemplate := range []*html_template.Template {
			_templates.homeHtml,
			_templates.indexHtml,
			_templates.librariesIndexHtml,
			_templates.documentsIndexHtml,
			_templates.libraryViewHtml,
			_templates.documentViewHtml,
			_templates.documentExportHtml,
			_templates.documentExportHtmlGithub,
			_templates.urlOpenHtml,
			_templates.urlErrorHtml,
			_templates.versionHtml,
	} {
		if _, _error := _topTemplate.New ("global-partials") .Parse (embedded.GlobalPartialsHtml); _error != nil {
			return nil, errorw (0x72af86c7, _error)
		}
	}
	
	return _templates, nil
}




func TemplatesAssetResolve (_templates *Templates, _path string) (string, []byte, *Error) {
	
	_data, _error := fs.ReadFile (_templates.assets, _path)
	if _error != nil {
		return "", nil, errorw (0x007f2426, _error)
	}
	
	// NOTE:  https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
	_contentType := ""
	switch path.Ext (_path) {
		
		case ".txt" :
			_contentType = "text/plain; charset=utf-8"
		case ".html" :
			_contentType = "text/plain; charset=utf-8"
		
		case ".css" :
			_contentType = "text/css; charset=utf-8"
		case ".js" :
			_contentType = "text/javascript; charset=utf-8"
		
		case ".png" :
			_contentType = "image/png"
		case ".jpeg" :
			_contentType = "image/jpeg"
		case ".svg" :
			_contentType = "image/svg+xml"
		case ".ico" :
			_contentType = "image/x-icon"
		
	}
	
	return _contentType, _data, nil
}

