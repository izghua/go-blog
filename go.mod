module github.com/izghua/go-blog

go 1.12

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-errors/errors v1.0.1
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-xorm/xorm v0.7.1
	github.com/izghua/zgh v0.0.24
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/mojocn/base64Captcha v0.0.0-20190509095025-87c9c59224d8
	github.com/qiniu/api.v7 v7.2.5+incompatible
	github.com/qiniu/x v7.0.8+incompatible // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/speps/go-hashids v2.0.0+incompatible
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f
	gopkg.in/russross/blackfriday.v2 v2.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.2.2
	qiniupkg.com/x v7.0.8+incompatible // indirect
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday v2.0.0+incompatible

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f

replace gopkg.in/yaml.v2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8
