

package embedded


import _ "embed"




//go:embed manual/z-scratchpad.txt
var ManualTxt string

//go:embed manual/z-scratchpad.html
var ManualHtml string

//go:embed manual/z-scratchpad.man
var ManualMan string

