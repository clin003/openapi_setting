package fonts

import (
	"embed"
)

//go:embed *.ttf
var FontList embed.FS
