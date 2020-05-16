module main

go 1.13

require (
	github.com/aws/aws-sdk-go v1.30.29
	github.com/gin-contrib/logger v0.0.2
	github.com/gin-gonic/gin v1.4.0
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/lib/pq v1.0.0 // indirect
	github.com/rs/zerolog v1.18.0
	github.com/ugorji/go/codec v0.0.0-20190320090025-2dc34c0b8780 // indirect
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
