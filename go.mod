module github.com/opensourceways/message-transfer

require (
	github.com/IBM/sarama v1.43.1 // indirect
	github.com/anshal21/json-flattener v1.0.0
	github.com/cloudevents/sdk-go/v2 v2.15.2
	github.com/lib/pq v1.10.9
	github.com/opensourceways/kafka-lib v0.0.0-20240410053850-1b04ade416f1
	github.com/opensourceways/server-common-lib v0.0.0-20240509113940-9336d8de8dba
	github.com/sirupsen/logrus v1.9.3
	gorm.io/datatypes v1.2.0
	gorm.io/driver/postgres v1.5.2
	gorm.io/gorm v1.25.4
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/jackc/pgx/v5 v5.5.5
	github.com/opensourceways/go-gitee v0.0.0-20240305060727-0df28a4f60c0
	github.com/todocoder/go-stream v1.1.2
	golang.org/x/xerrors v0.0.0-20240716161551-93cc26a95ae9
)

require (
	github.com/antihax/optional v1.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.6.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/oauth2 v0.12.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.4.7 // indirect
	k8s.io/apimachinery v0.30.1 // indirect
)

replace huaweicloud.com/apig/signer v0.0.0 => ./core

go 1.22.0

toolchain go1.22.2
