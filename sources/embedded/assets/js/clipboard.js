
(function () {
	
	"use strict";
	
	
	
	
	window.addEventListener ("DOMContentLoaded", (_event) => {
			_initialize ();
		});
	
	
	
	
	function _initialize () {
		
		for (let _target of document.querySelectorAll ([
				"html:root > body > header > h1",
				"html:root > body > header code",
				"html:root > body > header time",
		].join (", "))) {
			_register (_target, 1);
		}
		for (let _target of document.querySelectorAll ([
				"main.document:not(.document-format-text) h1",
				"main.document:not(.document-format-text) h2",
				"main.document:not(.document-format-text) h3",
				"main.document:not(.document-format-text) h4",
				"main.document:not(.document-format-text) h5",
				"main.document:not(.document-format-text) h6",
				"main.document:not(.document-format-text) p",
				"main.document:not(.document-format-text) ul",
				"main.document:not(.document-format-text) ol",
				"main.document:not(.document-format-text) dl",
				"main.document:not(.document-format-text) li",
				"main.document:not(.document-format-text) dt",
				"main.document:not(.document-format-text) dd",
				"main.document:not(.document-format-text) blockquote",
				"main.document:not(.document-format-text) pre",
				"main.document:not(.document-format-text) :not(pre) > code",
		].join (", "))) {
			_register (_target, 5);
		}
	}
	
	
	
	
	function _handle (_target, _event, _expectedClicks) {
		
		if (_target !== _currentTarget) {
			if (_currentTarget !== null) {
				_currentTarget.classList.remove ("clipboard-active");
				_currentTarget.classList.remove ("clipboard-copied");
			}
			_currentTarget = _target;
			_currentClicks = 1;
		} else {
			_currentClicks += 1;
		}
		
		if (_event.button == 1) {
			if ((_currentClicks == 1) && (_currentClicks < _expectedClicks)) {
				_currentClicks = _expectedClicks - 1;
			}
		}
		
		if (_currentClicks < _expectedClicks) {
			
			if (_currentClicks == (_expectedClicks - 1)) {
				_currentTarget.classList.add ("clipboard-active");
			}
			
			if (_currentTimeout !== null) {
				window.clearTimeout (_currentTimeout);
			}
			_currentTimeout = window.setTimeout (() => {
					_currentTarget.classList.remove ("clipboard-active");
					_currentTarget = null;
					_currentClicks = 0;
					_currentTimeout = null;
				}, 500);
			
		} else if (_currentClicks == _expectedClicks) {
			
			_copy (_target);
			
			_currentTarget.classList.remove ("clipboard-active");
			_currentTarget.classList.add ("clipboard-copied");
			if (_currentTimeout !== null) {
				window.clearTimeout (_currentTimeout);
			}
			_currentTimeout = window.setTimeout (() => {
					_currentTarget.classList.remove ("clipboard-copied");
					_currentTarget = null;
					_currentClicks = 0;
					_currentTimeout = null;
				}, 6000);
		}
		
		_event.stopPropagation ();
		_event.preventDefault ();
	}
	
	function _register (_target, _expectedClicks) {
		_target.classList.add ("clipboard-target");
		_target.addEventListener ("click", (_event) => {
				_handle (_target, _event, _expectedClicks);
			});
		_target.addEventListener ("auxclick", (_event) => {
				_handle (_target, _event, _expectedClicks);
			});
	}
	
	var _currentTarget = null;
	var _currentClicks = 0;
	var _currentTimeout = null;
	
	function _copy (_source) {
		
		var _selection = window.getSelection ();
		_selection.removeAllRanges ();
		
		var _range = document.createRange ();
		_range.selectNodeContents (_source);
		_selection.addRange (_range);
		
		var _selectionText = _selection.toString ();
		_selection.removeAllRanges ();
		
		if (true) {
			var _selectionEncoded = _selectionText;
			_selectionEncoded = encodeURIComponent (_selectionEncoded);
			_selectionEncoded = btoa (_selectionEncoded);
			_selectionEncoded = _selectionEncoded.replaceAll ("+", "-") .replaceAll ("/", "_") .replaceAll ("=", "");
			fetch ("/cs/" + _selectionEncoded)
				.then (_response => {
					if (_response.status != 204) {
						console.error ("[ee][76677dbf]", _response.status);
					}})
				.catch (_error => {
						console.error ("[ee][46101c4c]", _error);
					});
		} else {
			if (window.navigator.clipboard !== undefined) {
				window.navigator.clipboard.writeText (_selectionText);
			}
		}
	}
	
	
	
	
} ());

