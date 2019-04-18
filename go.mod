module github.com/izghua/go-blog

go 1.12

require (
	github.com/gin-gonic/gin v1.3.0
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-xorm/xorm v0.7.1
	github.com/izghua/zgh v0.0.20
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/qiniu/api.v7 v7.2.5+incompatible
	github.com/qiniu/x v7.0.8+incompatible // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/speps/go-hashids v2.0.0+incompatible
	gopkg.in/russross/blackfriday.v2 v2.0.0-00010101000000-000000000000
	qiniupkg.com/x v7.0.8+incompatible // indirect
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday v2.0.0+incompatible
