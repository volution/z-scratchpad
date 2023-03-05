

package embedded


import _ "embed"




//go:embed documentation/readme.txt
var ReadmeTxt string

//go:embed documentation/readme.html
var ReadmeHtml string


//go:embed documentation/z-scratchpad--help.txt
var ZscratchpadHelpTxt string

//go:embed documentation/z-scratchpad--manual.txt
var ZscratchpadManualTxt string

//go:embed documentation/z-scratchpad--manual.html
var ZscratchpadManualHtml string

//go:embed documentation/z-scratchpad--manual.man
var ZscratchpadManualMan string


//go:embed documentation/help--header.txt
var HelpHeader string

//go:embed documentation/help--footer.txt
var HelpFooter string


//go:embed documentation/sbom.txt
var SbomTxt string

//go:embed documentation/sbom.html
var SbomHtml string

//go:embed documentation/sbom.json
var SbomJson string

