module github.com/itgram/tracking_service

go 1.19

replace github.com/itgram/minion => ../../minion

replace github.com/itgram/minion.encoding => ../../minion.encoding

replace github.com/itgram/minion.encoding.protobuf => ../../minion.encoding.protobuf

replace github.com/itgram/minion.persistence => ../../minion.persistence

replace github.com/itgram/minion.persistence.esdb => ../../minion.persistence.esdb

replace github.com/itgram/minion.system => ../../minion.system

replace github.com/itgram/tracking_domain => ../domain

replace github.com/itgram/tracking_service => ./

require (
	github.com/itgram/minion.encoding.protobuf v0.0.0-00010101000000-000000000000
	github.com/itgram/minion.persistence v0.0.0-00010101000000-000000000000
	github.com/itgram/minion.persistence.esdb v0.0.0-00010101000000-000000000000
	github.com/itgram/minion.system v0.0.0-00010101000000-000000000000
	github.com/itgram/tracking_domain v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/EventStore/EventStore-Client-Go/v3 v3.0.0 // indirect
	github.com/Workiva/go-datastructures v1.0.53 // indirect
	github.com/armon/go-metrics v0.4.0 // indirect
	github.com/asynkron/gofun v0.0.0-20220329210725-34fed760f4c2 // indirect
	github.com/asynkron/protoactor-go v0.0.0-20220910074408-d2ceff064d72 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/consul/api v1.12.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.2.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.9.8 // indirect
	github.com/itgram/minion v0.0.0-00010101000000-000000000000 // indirect
	github.com/itgram/minion.encoding v0.0.0-00010101000000-000000000000 // indirect
	github.com/lithammer/shortuuid/v4 v4.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/orcaman/concurrent-map v1.0.0 // indirect
	github.com/prometheus/client_golang v1.12.2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.34.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	go.opentelemetry.io/otel v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.30.0 // indirect
	go.opentelemetry.io/otel/metric v0.30.0 // indirect
	go.opentelemetry.io/otel/sdk v1.7.0 // indirect
	go.opentelemetry.io/otel/sdk/export/metric v0.28.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.7.0 // indirect
	golang.org/x/exp v0.0.0-20220518171630-0b5c67f07fdf // indirect
	golang.org/x/net v0.0.0-20220526153639-5463443f8c37 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220526192754-51939a95c655 // indirect
	google.golang.org/grpc v1.46.2 // indirect
)
