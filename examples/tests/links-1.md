## links




## Auto-links


### Conclusions

* always prefix with one of `http://`, `https://` or `ftp://`;
* always bracket by `<...>`, to make sure no other punctuation or context is mistaken as part of the URL;


### Experiments

These should work:
* no bracketing:
  * www.example.com | www.example.com/path
  * https://example.com | https://example.com/path
  * ftp://example.com | ftp://example.com/path
  * user@example.com
* bracketed by `(...)`:
  * (www.example.com) | (www.example.com/path)
  * (https://example.com) | https://example.com/path)
  * (ftp://example.com) | (ftp://example.com/path)
  * (user@example.com)
* bracketed by `<...>`:
  * <https://example.com> | <https://example.com/path>
  * <ftp://example.com> | <ftp://example.com/path>
  * <user@example.com>
* bracketed by `<<...>>`:
  * <<https://example.com>> | <<https://example.com/path>>
  * <<ftp://example.com>> | <<ftp://example.com/path>>
  * <<user@example.com>>
* bracketed by `(<...>)`:
  * (<https://example.com>) | (<https://example.com/path>)
  * (<ftp://example.com>) | (<ftp://example.com/path>)
  * (<user@example.com>)

These should technically work, some are (but some aren't) what we expect:
* punctuation after the URL:
  * OK -- https://example.com/.
  * NOK -- https://example.com/, | https://example.com/; | https://example.com/? | https://example.com/!

These shouldn't work:
* example.com | example.com/path
* file://example.com | file://example.com/path
* gopher://example.com | gopher://example.com/path
* gemini://example.com | gemini://example.com/path
* mailto:user@example.com | mailto:user@example.com

These apparently don't work:
* <www.example.com> | <www.example.com/path>
* <<www.example.com>> | <<www.example.com/path>>
* [https://example.com/] | {https://example.com/} | |https://example.com/|
* [user@example.com] | {user@example.com} | |user@example.com|




## Markdown links


### Conclusions

* always use with a protocol;


### Experiments

With protocol:
* [https://example.com](https://example.com) | [https://example.com/path](https://example.com/path)
* [mailto:user@example.com](mailto:user@example.com)
* [gopher://example.com](gopher://example.com) | [gopher://example.com/path](gopher://example.com/path)
* [gemini://example.com](gemini://example.com) | [gemini://example.com/path](gemini://example.com/path)

Without protocol:
* [www.example.com](www.example.com) | [www.example.com/path](www.example.com/path)
* [user@example.com](user@example.com)

Internal links:
* [sl:some-library](sl:some-library) | [s:l/some-library](s:l/some-library)
* [sd:some-document](sd:some-document) | [s:d/some-document](s:d/some-document)
* [sd:some-library:some-document](sd:some-library:some-document) | [s:d/some-library:some-document](s:d/some-library:some-document)
* [w:some-term](w:some-term) | [s:w/some-term](s:w/some-term)

Links without label or target:
* [](https://example.com) | [](mailto:user@example.com)
* [](gopher://example.com)
* [](sl:some-library)
* [](sd:some-document)
* [](www.example.com) | [](user@example.com)
* [no-target]()
* (<a></a>) | (<a href=""></a>) | <a href="https://example.com/"></a>

