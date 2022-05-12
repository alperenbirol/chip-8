package widgets

import "github.com/AllenDang/giu"

func DisplayWidget(tex *giu.Texture) *giu.ImageWidget {
	return giu.Image(tex)
}
