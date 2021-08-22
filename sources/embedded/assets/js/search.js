
(function () {
	
	"use strict";
	
	
	
	
	window.addEventListener ("DOMContentLoaded", (_event) => {
			_initialize ();
		});
	
	
	
	
	let _queryInput;
	let _resultsList;
	let _resultsCurrentIndex;
	let _resultsIdCounter;
	let _candidatesCache;
	
	
	
	
	function _initialize () {
		
		_queryInput = document.getElementById ("search-query");
		_resultsList = document.getElementById ("search-results");
		
		if ((_queryInput == null) || (_resultsList == null)) {
//			console.debug ("[dd][dd18a5aa]");
			return null;
		}
		
		_resultsCurrentIndex = null;
		_resultsIdCounter = 0;
		_candidatesCache = null;
		
		_queryInput.oninput = (_event) => {
				let _query = _queryInput.value;
				_search (_query);
			};
		
		_queryInput.onkeydown = (_event) => {
				switch (_event.key) {
					
					case "Escape" :
						_queryInput.value = "";
						_search (null);
						_event.stopPropagation ();
						_event.preventDefault ();
						break;
					
					case "ArrowUp" :
						if (_resultsCurrentIndex == null) {
							_resultsCurrentIndex = _resultsList.children.length;
						} else {
							_resultsCurrentIndex -= 1;
						}
						_highlightResult ();
						_event.stopPropagation ();
						_event.preventDefault ();
						break;
					
					case "ArrowDown" :
						if (_resultsCurrentIndex == null) {
							_resultsCurrentIndex = 0;
						} else {
							_resultsCurrentIndex += 1;
						}
						_highlightResult ();
						if (_resultsCurrentIndex == null) {
							_unfocus ();
						} else {
							_event.stopPropagation ();
							_event.preventDefault ();
						}
						break;
					
					case "End" :
						if (_queryInput.value == "") {
							_unfocus ();
						}
						break;
					
					case "Enter" :
						_activateResult (_event);
						_event.stopPropagation ();
						_event.preventDefault ();
						break;
					
					case "PageDown" :
					case "PageUp" :
						_unfocus ();
						break;
					
					default :
//						console.debug ("[dd][370cf584]", _event.key);
						break;
				}
			};
		
		window.addEventListener ("keypress", (_event) => {
				if (_event.key == "/") {
					if (document.activeElement != _queryInput) {
						_event.stopPropagation ();
						_event.preventDefault ();
						_focus ();
					}
				}
			});
		
		window.addEventListener ("scroll", (_event) => {
				if (window.scrollY == 0) {
					_focus ();
				} else if (document.activeElement == _queryInput) {
					_unfocus ();
				}
			});
		
		window.addEventListener ("load", (_event) => {
				if (window.scrollY == 0) {
					_focus ();
				}
				{
					let _query = _queryInput.value;
					_search (_query);
				}
			});
	}
	
	
	
	
	function _focus () {
		window.scrollTo (0, 0);
		_queryInput.focus ();
		_queryInput.select ();
	}
	
	function _unfocus () {
		document.activeElement.blur ();
	}
	
	
	
	
	function _search (_query) {
		
		let _filter = _resolveFilter (_query);
		let _candidates = _resolveCandidates ();
		let _results = [];
		
		for (let _candidate of _candidates) {
			if (_filter (_candidate)) {
//				console.debug ("[dd][0fe0c4a8]", _candidate);
				_candidate.resultLabel = _candidate.candidateLabel;
				_candidate.resultLink = _candidate.candidateLink;
				_candidate.resultDetails = _candidate.candidateDetails;
				_results.push (_candidate);
			}
		}
		
		_updateResults (_results);
	}
	
	
	
	
	function _resolveFilter (_query) {
		
		if (_query === null) {
			return ((_cadidate) => false);
		}
		_query = _query.toLowerCase () .trim ();
		if (_query == "?") {
			return ((_candidate) => true);
		}
		if (_query == "") {
			return ((_cadidate) => false);
		}
		if (_query[0] == "?") {
			_query = _query.substring (1) .trim ();
		}
		let _queryTokens = _query.split (/\s+/);
		
		function _candidateTextMatches (_candidateText) {
			for (let _token of _queryTokens) {
				if (! _candidateText.includes (_token)) {
					return (false);
				}
			}
			return (true);
		}
		
		function _candidateMatches (_candidate) {
			let _matches = false;
			_matches = _matches || _candidateTextMatches (_candidate.candidateText);
			if (_candidate.candidateDetails !== null) {
				_matches = _matches || _candidateTextMatches (_candidate.candidateDetails);
			}
			return (_matches);
		}
		
		return (_candidateMatches);
	}
	
	
	
	
	function _resolveCandidates () {
		
		if (_candidatesCache !== null) {
			return (_candidatesCache);
		}
		
		function _simplify (_node) {
			let _childFirst = null;
			let _childCount = 0;
			let _hasContent = false;
			for (let _child of _node.childNodes) {
				let _childText = _child.textContent.trim ();
				if (_childText == "") {
					continue;
				}
				_hasContent = true;
				if (_child.nodeType == Node.ELEMENT_NODE) {
					if (_childFirst === null) {
						_childFirst = _child;
					}
					_childCount += 1;
				}
			}
			if (!_hasContent) {
				return (null);
			}
			if (_node.tagName == "A") {
				return (_node);
			} else if (_childCount == 1) {
				if (_childFirst.tagName == "A") {
					return (_childFirst);
				} else {
					return (_simplify (_childFirst));
				}
			} else {
				return (_node);
			}
		}
		
		function _collect (_candidateNode, _labelPrefix, _textPrefix) {
			
			let _candidateLabel = _candidateNode.textContent;
			_candidateLabel = _candidateLabel.trim ();
			if (_labelPrefix != "") {
				_candidateLabel = _labelPrefix + " " + _candidateLabel;
			}
			
			let _candidateText = _candidateNode.textContent;
			_candidateText = _candidateText.toLowerCase () .trim ();
			if (_textPrefix != "") {
				_candidateText = _textPrefix + " " + _candidateText;
			}
			
			let _candidateLink = null;
			let _candidateDetails = null;
			if ((_candidateNode.tagName == "A") && (_candidateNode.href != "")) {
				_candidateLink = _candidateNode.href;
				_candidateDetails = _candidateNode.dataset.zsUrlOriginalHref || null;
			} else {
				if (_candidateNode.id == "") {
					_resultsIdCounter += 1;
					_candidateNode.id = "search-candidate-" + _resultsIdCounter;
				}
				_candidateLink = "#" + _candidateNode.id;
			}
			
			_candidates.push ({
					node : _candidateNode,
					candidateLabel : _candidateLabel,
					candidateText : _candidateText,
					candidateLink : _candidateLink,
					candidateDetails : _candidateDetails,
				});
		}
		
		let _candidates = [];
		
		for (let _candidateNode of document.querySelectorAll (".search-candidate")) {
			_candidateNode = _simplify (_candidateNode);
			if (_candidateNode !== null) {
				_collect (_candidateNode, "", "");
			}
		}
		
		for (let _mainNode of document.querySelectorAll ("main.document")) {
			for (let _candidateNode of _mainNode.querySelectorAll ("h1, h2, h3, h4, h5, h6, a[data-zs-url-type=\"external\"], a[data-zs-url-type=\"internal\"]")) {
				let _candidatePrefix = "<" + _candidateNode.tagName + ">";
				_candidateNode = _simplify (_candidateNode);
				if (_candidateNode !== null) {
					_collect (_candidateNode, _candidatePrefix, _candidatePrefix.toLowerCase ());
				}
			}
		}
		
		_candidatesCache = _candidates;
		
		return (_candidates);
	}
	
	
	
	
	function _updateResults (_results) {
		while (_resultsList.firstChild !== null) {
			_resultsList.removeChild (_resultsList.firstChild);
		}
		for (let _result of _results) {
			let _resultNode = document.createElement ("li");
			_resultNode.classList.add ("search-result");
			if (_result.resultLink !== null) {
				let _linkNode = document.createElement ("a");
				_linkNode.classList.add ("search-result-link");
				_linkNode.href = _result.resultLink;
				_linkNode.textContent = _result.resultLabel;
				_resultNode.appendChild (_linkNode);
			} else {
				_resultNode.textContent = _result.resultLabel;
			}
			if (_result.resultDetails !== null) {
				let _detailsNode = document.createElement ("span");
				_detailsNode.classList.add ("search-result-details");
				_detailsNode.textContent = _result.resultDetails;
				_resultNode.appendChild (document.createTextNode (" "));
				_resultNode.appendChild (_detailsNode);
			}
			_resultsList.appendChild (_resultNode);
		}
		_highlightResult ();
	}
	
	
	
	
	function _highlightResult () {
		for (let _node of _resultsList.querySelectorAll (".search-result-highlight")) {
			_node.classList.remove ("search-result-highlight");
		}
		let _highlighted = null;
		if (_resultsList.children.length == 1) {
			_highlighted = _resultsList.children[0];
		} else if (_resultsCurrentIndex !== null) {
			if (_resultsList.children.length == 0) {
				_resultsCurrentIndex = null;
			} else {
				if (_resultsCurrentIndex < 0) {
					_resultsCurrentIndex = 0;
				}
				if (_resultsCurrentIndex >= _resultsList.children.length) {
					_resultsCurrentIndex = _resultsList.children.length - 1;
				}
				_highlighted = _resultsList.children[_resultsCurrentIndex];
			}
		}
		if (_highlighted !== null) {
			_highlighted.classList.add ("search-result-highlight");
			_highlighted.scrollIntoView ({block : "center"});
		}
	}
	
	
	
	
	function _activateResult (_event) {
		
		let _highlighted = _resultsList.querySelector (".search-result-highlight");
		if (_highlighted === null) {
			return;
		}
		
		_queryInput.select ();
		
		let _target = _highlighted.firstChild;
		if (_target === null) {
			return;
		}
		
		if ((_target.tagName == "A") && (_target.href != "")) {
			let _newTab = false;
			if ((_event !== undefined) && _event.ctrlKey) {
				_newTab = true;
			}
			let _options = "noopener=yes,noreferrer=yes";
			if (_newTab) {
				window.open (_target.href, "_blank", _options);
			} else {
				window.open (_target.href, "_top", _options);
			}
		} else {
			_target.click ();
		}
	}
	
	
	
	
} ());

