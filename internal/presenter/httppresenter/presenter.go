package httppresenter

type IPresenter interface {
	WithData(data interface{}) IPresenter
	WithMeta(meta map[string]interface{}) IPresenter
	WithExtension(key string, exten interface{}) IPresenter
	ToMap() map[string]interface{}
}

type Presenter struct {
	resource map[string]interface{}
}

func (p *Presenter) WithData(data interface{}) IPresenter {
	// This is the primary data of the document. It represents the main content that the API is delivering.
	// The data member can either be a single resource object or an array of resource objects.
	// "data": {
	//    "id": "1",
	//    "attributes": {
	//      "title": "Understanding JSON:API",
	//      "author": "John Doe"
	//    }
	//  }
	p.resource["data"] = data

	return p
}

func (p *Presenter) WithMeta(meta map[string]interface{}) IPresenter {
	// This member contains non-standard meta-information that may be relevant to the client
	// but is not part of the primary data. This can include pagination information, timestamps,
	// or other auxiliary data.
	// "meta": {
	//    "version": "1.0",
	//    "timestamp": "2024-10-30T12:00:00Z"
	//  }
	p.resource["meta"] = meta

	return p
}

func (p *Presenter) WithExtension(key string, exten interface{}) IPresenter {
	// Extensions allow the addition of new members to the document that are not part of the core JSON
	// specification. This can be useful for including additional context or functionality without breaking
	// existing standards
	// "custom_extension": {
	//    "custom_field": "This is additional information specific to the implementation."
	//  }
	p.resource[key] = exten

	return p
}

func (p *Presenter) ToMap() map[string]interface{} {
	return p.resource
}

func New(opts ...func(*Presenter)) *Presenter {
	p := &Presenter{resource: make(map[string]interface{})}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func WithData(data interface{}) func(*Presenter) {
	return func(p *Presenter) {
		p.resource["data"] = data
	}
}

func WithMeta(meta map[string]interface{}) func(*Presenter) {
	return func(p *Presenter) {
		p.resource["meta"] = meta
	}
}

func WithExtension(key string, ex interface{}) func(*Presenter) {
	return func(p *Presenter) {
		p.resource[key] = ex
	}
}

//httppresenter.New(
//httppresenter.WithData(resp.User),
//httppresenter.WithMeta(map[string]interface{}{"execution_duration": 3 * time.Microsecond}),
//httppresenter.WithExtension("status", map[string]interface{}{
//"uptime":           "172800",
//"last_maintenance": "2024-10-28T15:30:00Z",
//"load": map[string]interface{}{
//"cpu":    "65%",
//"memory": "75%",
//},
//}),
//).ToMap()

// type IPresenter interface {
//	WithData(data interface{}) *Presenter
//	WithMeta(m map[string]interface{}) *Presenter
//	WithExtention(key string, ex map[string]interface{}) *Presenter
//	ToMap() map[string]interface{}
//	Ok(data interface{}) map[string]interface{}
//	SuccessWithMSG(msg string, data interface{}) map[string]interface{}
//	Error(err error) (int, map[string]interface{})
//	ErrorWithMSG(msg string, err error) map[string]interface{}
//}
//
//type Presenter struct {
//	resource map[string]interface{}
//}
//
//func New() *Presenter {
//	return &Presenter{resource: make(map[string]interface{})}
//}
//
//func (p *Presenter) WithData(data interface{}) *Presenter {
//	p.resource["data"] = data
//
//	return p
//}
//
//func (p *Presenter) WithMeta(m map[string]interface{}) *Presenter {
//	p.resource["meta"] = m
//
//	return p
//}
//
//func (p *Presenter) WithExtention(key string, ex map[string]interface{}) *Presenter {
//	p.resource[key] = ex
//
//	return p
//}
//
//func (p *Presenter) ToMap() map[string]interface{} {
//	return p.resource
//}
//
//func (p *Presenter) Ok(data interface{}) map[string]interface{} {
//	return map[string]interface{}{
//		"status": true,
//		"data":   data,
//	}
//}
//
//func (p *Presenter) SuccessWithMSG(msg string, data interface{}) map[string]interface{} {
//	return map[string]interface{}{
//		"status":  true,
//		"message": msg,
//		"data":    data,
//	}
//}
//
//func (p *Presenter) Error(err error) (int, map[string]interface{}) {
//	richErr, _ := richerror.Analysis(err)
//	code := httpstatus.FromKind(richErr.Kind())
//
//	var errs interface{} = richErr.Error()
//	if len(errs.(string)) == 0 {
//		errs = richErr.Meta()
//	}
//	return code, map[string]interface{}{
//		"status":  false,
//		"message": richErr.Message(),
//		"errors":  errs,
//	}
//}
//
//func (p *Presenter) ErrorWithMSG(msg string, err error) map[string]interface{} {
//	return map[string]interface{}{
//		"status":  false,
//		"message": msg,
//		"errors":  err.Error(),
//	}
//}
