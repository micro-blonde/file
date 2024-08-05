module github.com/micro-blonde/file/client

go 1.22.3

require (
	github.com/ginger-core/compound/registry v0.0.0-20230608151919-2963b75416c3
	github.com/ginger-core/errors v0.0.0-20230703084505-b10c3f9cedfb
	github.com/ginger-core/errors/grpc v0.0.0-20230703084505-b10c3f9cedfb
	github.com/ginger-core/log v0.0.0-20240629145652-3b2876535940
	github.com/micro-blonde/file v0.0.0-20240805211824-0b9fff068040
	github.com/micro-blonde/file/proto v0.0.0-20240805211729-18bd31844877
	google.golang.org/grpc v1.65.0
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/ginger-core/compound v0.0.0-20230608151919-2963b75416c3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.16.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace github.com/micro-blonde/file/proto => ../proto
// replace github.com/micro-blonde/file => ../
