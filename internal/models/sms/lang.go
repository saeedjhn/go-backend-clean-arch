package sms

type Lang string

const (
	LangFA Lang = "fa"
	LangEN Lang = "en"
	LangAR Lang = "ar"

	LangDefault Lang = "en"
)

func IsValidLang(lang Lang) bool {
	switch lang {
	case LangFA, LangEN, LangAR:
		return true
	default:
		return false
	}
}

func (l Lang) ToLang(rawLang string) Lang {
	switch rawLang {
	case string(LangEN):
		return LangEN
	case string(LangFA):
		return LangFA
	case string(LangAR):
		return LangAR
	default:
		return LangDefault
	}
}
