package constants

type WebRenderer int64

const html = "html"
const canvaskit = "canvaskit"
const rendererUndefined = "UNDEFINED"

const (
	RendererUndefined WebRenderer = iota
	Html
	Canvaskit
)

func (renderer WebRenderer) String() string {
	switch renderer {
	case Html:
		return html
	case Canvaskit:
		return canvaskit
	}
	return rendererUndefined
}

func ToWebRenderer(renderer string) WebRenderer {
	switch renderer {
	case html:
		return Html
	case canvaskit:
		return Canvaskit
	}
	return RendererUndefined
}

func GetRendererList() []string {
	return []string{html, canvaskit}
}
