module github.com/go-sanpang/go-admin-core/plugins/logger/zap

go 1.18

require (
	github.com/go-sanpang/go-admin-core v1.5.2
	go.uber.org/zap v1.26.0
)

require go.uber.org/multierr v1.11.0 // indirect

replace github.com/go-sanpang/go-admin-core => ../../../
