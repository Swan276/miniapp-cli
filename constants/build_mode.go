package constants

type BuildMode int64

const merchant = "Merchant"
const wavepay = "Wavepay"
const empty = "Empty"
const buildUndefined = "UNDEFINED"

const merchantUrl = "/tnx-analytics-web-app/"
const wavepayUrl = "/wave-tnx-history-web-app/"
const emptyUrl = "/"

const (
	BuildUndefined BuildMode = iota
	Merchant
	Wavepay
	Empty
)

func (build BuildMode) String() string {
	switch build {
	case Merchant:
		return merchant
	case Wavepay:
		return wavepay
	case Empty:
		return empty
	}
	return buildUndefined
}

func (build BuildMode) BaseUrl() string {
	switch build {
	case Merchant:
		return merchantUrl
	case Wavepay:
		return wavepayUrl
	case Empty:
		return emptyUrl
	}
	return emptyUrl
}

func ToBuildMode(build string) BuildMode {
	switch build {
	case merchant:
		return Merchant
	case wavepay:
		return Wavepay
	case empty:
		return Empty
	}
	return BuildUndefined
}

func GetBuildModeList() []string {
	return []string{merchant, wavepay, empty}
}
