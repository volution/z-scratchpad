

package zscratchpad


import "bytes"
import "encoding/gob"
import "fmt"
import "os"
import "sort"
import "time"
import "syscall"




type Index struct {
	globals *Globals
	documents map[string]*Document
	libraries map[string]*Library
	libraryDocuments map[string]map[string]bool
	dirtyCallback func () ()
}


type IndexGob struct {
	Documents []*Document
	Libraries []*Library
	LibraryDocuments []IndexLibraryDocumentsGob
}

type IndexLibraryDocumentsGob struct {
	Library string
	Documents []string
}




func IndexNew (_globals *Globals) (*Index, *Error) {
	_index := & Index {
			globals : _globals,
			documents : make (map[string]*Document, 16 * 1024),
			libraries : make (map[string]*Library, 128),
			libraryDocuments : make (map[string]map[string]bool, 128),
		}
	return _index, nil
}




func IndexLoadFromPath (_index *Index, _path string) (bool, *Error) {
	
	_file, _error := os.OpenFile (_path, os.O_RDONLY, 0)
	if _error != nil {
		return false, errorw (0x8336e730, _error)
	}
	defer _file.Close ()
	
	_buffer := (*bytes.Buffer) (nil)
	
	if true {
		
		_stat, _error := _file.Stat ()
		if _error != nil {
			return false, errorw (0xab7f7fce, _error)
		}
		_size := _stat.Size ()
		if _size == 0 {
			return false, errorw (0x2ee0990b, _error)
		}
		
		_memory, _error := syscall.Mmap (int (_file.Fd ()), 0, int (_size), syscall.PROT_READ, syscall.MAP_SHARED)
		if _error != nil {
			return false, errorw (0x486ddaae, _error)
		}
		defer syscall.Munmap (_memory)
		
		_buffer = bytes.NewBuffer (_memory)
		
	} else {
		
		_buffer := BytesBufferNewSize (64 * 1024 * 1024)
		defer BytesBufferRelease (_buffer)
		
		_, _error = _buffer.ReadFrom (_file)
		if _error != nil {
			return false, errorw (0xc25b03c8, _error)
		}
	}
	
	return IndexLoadFromBuffer (_index, _buffer)
}


func IndexLoadFromBuffer (_index *Index, _buffer *bytes.Buffer) (bool, *Error) {
	
	if _buffer.Len () <= len (BUILD_SOURCES_HASH) {
		return false, errorw (0x179a8718, nil)
	}
	
	if string (_buffer.Next (len (BUILD_SOURCES_HASH))) != BUILD_SOURCES_HASH {
		return false, nil
	}
	
	_gob := & IndexGob {}
	
	if true {
		_dataSize, _error := _gob.Unmarshal (_buffer.Bytes ())
		if _error != nil {
			return false, errorw (0x2056077c, _error)
		}
		if _dataSize != uint64 (_buffer.Len ()) {
			return false, errorw (0x611cbd94, nil)
		}
	} else {
		_decoder := gob.NewDecoder (_buffer)
		_error := _decoder.Decode (_gob)
		if _error != nil {
			return false, errorw (0x17ed45c1, _error)
		}
	}
	
	if _error := IndexLoadData (_index, _gob); _error != nil {
		return false, _error
	}
	
	return true, nil
}


func IndexStoreToPath (_index *Index, _path string) (*Error) {
	
	_buffer := BytesBufferNewSize (64 * 1024 * 1024)
	defer BytesBufferRelease (_buffer)
	
	if _error := IndexStoreToBuffer (_index, _buffer); _error != nil {
		return _error
	}
	
	_pathTemp := fmt.Sprintf ("%s.%d-%d.tmp", _path, time.Now () .UnixMilli (), "tmp")
	
	_file, _error := os.OpenFile (_pathTemp, os.O_WRONLY | os.O_CREATE | os.O_EXCL, 0o640)
	if _error != nil {
		return errorw (0xa2237854, _error)
	}
	defer _file.Close ()
	
	_, _error = _buffer.WriteTo (_file)
	if _error != nil {
		return errorw (0x4170a0f9, _error)
	}
	
	_error = os.Rename (_pathTemp, _path)
	if _error != nil {
		return errorw (0x45fd20b2, _error)
	}
	
	return nil
}


func IndexStoreToBuffer (_index *Index, _buffer *bytes.Buffer) (*Error) {
	
	_buffer.WriteString (BUILD_SOURCES_HASH)
	
	_gob := & IndexGob {}
	
	if _error := IndexStoreData (_index, _gob); _error != nil {
		return _error
	}
	
	if true {
		_bufferBytes := _buffer.Bytes ()
		_bufferBytes = _bufferBytes [len (_bufferBytes) :]
		_data, _error := _gob.Marshal (_bufferBytes)
		if _error != nil {
			return errorw (0x19b25b1a, _error)
		}
		_buffer.Write (_data)
	} else {
		_encoder := gob.NewEncoder (_buffer)
		_error := _encoder.Encode (_gob)
		if _error != nil {
			return errorw (0x7e0b1a3d, _error)
		}
	}
	
	return nil
}


func IndexLoadData (_index *Index, _gob *IndexGob) (*Error) {
	
	_libraries := make (map[string]*Library, len (_gob.Libraries))
	for _, _library := range _gob.Libraries {
		if _error := libraryInitializeMatchers (_library); _error != nil {
			return _error
		}
		_libraries[_library.Identifier] = _library
	}
	
	_documents := make (map[string]*Document, len (_gob.Documents))
	for _, _document := range _gob.Documents {
		_documents[_document.Identifier] = _document
	}
	
	_libraryDocuments := make (map[string]map[string]bool, len (_gob.LibraryDocuments))
	for _, _libraryDocumentsGob := range _gob.LibraryDocuments {
		_libraryDocumentsMap := make (map[string]bool, len (_libraryDocumentsGob.Documents))
		for _, _documentIdentifier := range _libraryDocumentsGob.Documents {
			_libraryDocumentsMap[_documentIdentifier] = true
		}
		_libraryDocuments[_libraryDocumentsGob.Library] = _libraryDocumentsMap
	}
	
	_index.libraries = _libraries
	_index.documents = _documents
	_index.libraryDocuments = _libraryDocuments
	
	return nil
}


func IndexStoreData (_index *Index, _gob *IndexGob) (*Error) {
	
	_libraries := make ([]*Library, 0, len (_index.libraries))
	for _, _library := range _index.libraries {
		_libraries = append (_libraries, _library)
	}
	
	_documents := make ([]*Document, 0, len (_index.documents))
	for _, _document := range _index.documents {
		_documents = append (_documents, _document)
	}
	
	_libraryDocuments := make ([]IndexLibraryDocumentsGob, 0, len (_index.libraryDocuments))
	for _libraryIdentifier, _libraryDocumentsMap := range _index.libraryDocuments {
		_documentIdentifiers := make ([]string, 0, len (_libraryDocuments))
		for _documentIdentifier, _ := range _libraryDocumentsMap {
			_documentIdentifiers = append (_documentIdentifiers, _documentIdentifier)
		}
		_libraryDocumentsGob := IndexLibraryDocumentsGob { _libraryIdentifier, _documentIdentifiers }
		_libraryDocuments = append (_libraryDocuments, _libraryDocumentsGob)
	}
	
	_gob.Libraries = _libraries
	_gob.Documents = _documents
	_gob.LibraryDocuments = _libraryDocuments
	
	return nil
}





func IndexLibraryInclude (_index *Index, _library *Library) (*Error) {
	if _library.Identifier == "" {
		return errorw (0xcba1a612, nil)
	}
	if _, _exists := _index.libraries[_library.Identifier]; _exists {
		return errorw (0xb8990674, nil)
	}
	_index.libraries[_library.Identifier] = _library
	_index.libraryDocuments[_library.Identifier] = make (map[string]bool, 16 * 1024)
	if _index.dirtyCallback != nil {
		_index.dirtyCallback ()
	}
	return nil
}




func IndexDocumentInclude (_index *Index, _document *Document) (*Error) {
	if _document.Identifier == "" {
		return errorw (0x5567bd56, nil)
	}
	if _document.Library == "" {
		return errorw (0xf45e4633, nil)
	}
	if _, _exists := _index.libraries[_document.Library]; !_exists {
		return errorw (0x4c68b8a9, nil)
	}
	if _, _exists := _index.documents[_document.Identifier]; _exists {
//		logf ('d', 0x83e5bd38, "%s", _document.Path)
		return errorw (0x9c9f9c42, nil)
	}
	_index.documents[_document.Identifier] = _document
	_index.libraryDocuments[_document.Library][_document.Identifier] = true
	if _index.dirtyCallback != nil {
		_index.dirtyCallback ()
	}
	return nil
}


func IndexDocumentExclude (_index *Index, _document *Document) (*Error) {
	if _document.Identifier == "" {
		return errorw (0x18c096d9, nil)
	}
	if _document.Library == "" {
		return errorw (0x73f953db, nil)
	}
	if _, _exists := _index.libraries[_document.Library]; !_exists {
		return errorw (0x40bb92b4, nil)
	}
	if _, _exists := _index.documents[_document.Identifier]; !_exists {
		return errorw (0x20f5597e, nil)
	}
	delete (_index.documents, _document.Identifier)
	delete (_index.libraryDocuments[_document.Library], _document.Identifier)
	if _index.dirtyCallback != nil {
		_index.dirtyCallback ()
	}
	return nil
}


func IndexDocumentUpdate (_index *Index, _documentNew *Document, _documentOld *Document) (*Error) {
	if _error := IndexDocumentExclude (_index, _documentOld); _error != nil {
		return _error
	}
	if _error := IndexDocumentInclude (_index, _documentNew); _error != nil {
		return _error
	}
	return nil
}




func IndexLibrariesSelectAll (_index *Index) ([]*Library, *Error) {
	_libraries := make ([]*Library, 0, len (_index.libraries))
	for _, _library := range _index.libraries {
		_libraries = append (_libraries, _library)
	}
	LibrariesSort (_libraries)
	return _libraries, nil
}

func IndexDocumentsSelectAll (_index *Index) ([]*Document, *Error) {
	_documents := make ([]*Document, 0, len (_index.documents))
	for _, _document := range _index.documents {
		_documents = append (_documents, _document)
	}
	DocumentsSort (_documents)
	return _documents, nil
}

func IndexDocumentsSelectInLibrary (_index *Index, _libraryIdentifier string) ([]*Document, *Error) {
	_documents := make ([]*Document, 0, len (_index.documents))
	_libraryDocuments, _libraryExists := _index.libraryDocuments[_libraryIdentifier]
	if !_libraryExists {
		return nil, errorw (0xb14719d6, nil)
	}
	for _documentIdentifier, _ := range _libraryDocuments {
		_document, _ := _index.documents[_documentIdentifier]
		if _document == nil {
			return nil, errorw (0xef6c0449, nil)
		}
		_documents = append (_documents, _document)
	}
	DocumentsSort (_documents)
	return _documents, nil
}




func IndexLibraryResolve (_index *Index, _identifier string) (*Library, *Error) {
	if _identifier == "" {
		return nil, errorw (0xabe00a27, nil)
	}
	_library, _ := _index.libraries[_identifier]
	return _library, nil
}

func IndexDocumentResolve (_index *Index, _identifier string) (*Document, *Error) {
	if _identifier == "" {
		return nil, errorw (0x25e42042, nil)
	}
	_document, _ := _index.documents[_identifier]
	return _document, nil
}




func LibrariesSort (_libraries []*Library) () {
	sort.Slice (_libraries, func (_leftIndex, _rightIndex int) (bool) {
			_left := _libraries[_leftIndex]
			_right := _libraries[_rightIndex]
			return compareByNameOrIdentifier (_left.Name, _left.Identifier, _right.Name, _right.Identifier)
		})
}

func DocumentsSort (_documents []*Document) () {
	sort.Slice (_documents, func (_leftIndex, _rightIndex int) (bool) {
			_left := _documents[_leftIndex]
			_right := _documents[_rightIndex]
			return compareByNameOrIdentifier (_left.Title, _left.Identifier, _right.Title, _right.Identifier)
		})
}


func compareByNameOrIdentifier (_leftTitle, _leftIdentifier, _rightTitle, _rightIdentifier string) (bool) {
	if (_leftTitle != "") && (_rightTitle != "") {
		if _leftTitle == _rightTitle {
			return _leftIdentifier < _rightIdentifier
		}
		return _leftTitle < _rightTitle
	} else if (_leftTitle == "") && (_rightTitle == "") {
		return _leftIdentifier < _rightIdentifier
	} else if (_leftTitle != "") {
		return true
	} else if (_rightTitle != "") {
		return false
	} else {
		return false
	}
}

