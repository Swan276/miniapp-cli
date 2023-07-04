package constants

type Env int64

const tnd = "TND"
const sit = "SIT"
const preprod = "PREPROD"
const prod = "PROD"
const envUndefined = "UNDEFINED"

const (
	EnvUndefined Env = iota
	TND
	SIT
	PREPROD
	PROD
)

func (env Env) String() string {
	switch env {
	case TND:
		return tnd
	case SIT:
		return sit
	case PREPROD:
		return preprod
	case PROD:
		return prod
	}
	return envUndefined
}

func ToEnv(env string) Env {
	switch env {
	case tnd:
		return TND
	case sit:
		return SIT
	case preprod:
		return PREPROD
	case prod:
		return PROD
	}
	return EnvUndefined
}

func GetEnvList() []string {
	return []string{tnd, sit, preprod, prod}
}
