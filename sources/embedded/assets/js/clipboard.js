
(function () {
	
	"use strict";
	
	
	
	
	window.addEventListener ("DOMContentLoaded", (_event) => {
			_initialize ();
		});
	
	
	
	
	function _initialize () {
		
		for (let _target of document.querySelectorAll ("main.document :not(pre) code")) {
			_target.addEventListener ("click", (_event) => {
					if (_event.detail == 2) {
						_handle (_target, _event);
					}
				});
		}
		for (let _target of document.querySelectorAll ("main.document code")) {
			_target.addEventListener ("click", (_event) => {
					if (_event.detail == 3) {
						_handle (_target, _event);
					}
				});
		}
		for (let _target of document.querySelectorAll ("main.document > *")) {
			_target.addEventListener ("click", (_event) => {
					if (_event.detail == 3) {
						_handle (_target, _event);
					}
				});
		}
	}
	
	
	
	
	function _handle (_target, _event) {
		_copy (_target);
		_event.stopPropagation ();
		_event.preventDefault ();
	}
	
	function _copy (_source) {
		var _selection = window.getSelection ();
		_selection.removeAllRanges ();
		var _range = document.createRange ();
		_range.selectNodeContents (_source);
		_selection.addRange (_range);
		window.navigator.clipboard.writeText (_selection.toString ());
	}
	
	
	
	
} ());

