package sms

type Template struct {
	Content map[Lang]string // {"en": "Hello {{.Name}}", "fa": "سلام {{.Name}}"}
	lang    Lang            // "en"
}

// func NewTemplate(content map[string]string, langCode string) (*Template, error) {
// 	if _, ok := content[langCode]; !ok {
// 		return nil, errors.New("default language must exist in content")
// 	}
// if !IsValidLanguage(defaultLang) {
// return nil, errors.New("invalid language")
// }
// 	return &Template{Content: content, lang: langCode}, nil
// }

func (t *Template) GetLocalizedContent(lang Lang) string {
	if content, ok := t.Content[lang]; ok {
		return content
	}

	return t.Content[t.lang]
}
