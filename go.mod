module github.com/aliyun/terraform-provider-alicloud

replace github.com/googleapis/gnostic v0.5.6 => github.com/googleapis/gnostic v0.5.5

replace github.com/hashicorp/terraform-exec => ./local/terraform-exec

replace github.com/hashicorp/terraform-plugin-sdk => ./local/terraform-plugin-sdk

replace k8s.io/client-go => ./local/client-go

replace github.com/aliyun/alibaba-cloud-sdk-go => ./local/alibaba-cloud-sdk-go

require (
	cloud.google.com/go/iam v0.3.0 // indirect
	cloud.google.com/go/storage v1.21.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/PaesslerAG/gval v1.1.2 // indirect
	github.com/PaesslerAG/jsonpath v0.1.1
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/agiledragon/gomonkey/v2 v2.3.1
	github.com/alibabacloud-go/cs-20151215/v2 v2.4.5
	github.com/alibabacloud-go/darabonba-openapi v0.1.16
	github.com/alibabacloud-go/endpoint-util v1.1.1 // indirect
	github.com/alibabacloud-go/openapi-util v0.0.11 // indirect
	github.com/alibabacloud-go/tea v1.1.17
	github.com/alibabacloud-go/tea-roa v1.3.3
	github.com/alibabacloud-go/tea-rpc v1.3.3
	github.com/alibabacloud-go/tea-utils v1.4.4
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1264
	github.com/aliyun/aliyun-datahub-sdk-go v0.1.5
	github.com/aliyun/aliyun-log-go-sdk v0.1.28-0.20220211074227-7b7d575ad8d3
	github.com/aliyun/aliyun-mns-go-sdk v1.0.2
	github.com/aliyun/aliyun-oss-go-sdk v2.2.1+incompatible
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.3+incompatible
	github.com/aliyun/credentials-go v1.2.1
	github.com/aliyun/fc-go-sdk v0.0.0-20211130132118-64c3dff91779
	github.com/aws/aws-sdk-go v1.43.22 // indirect
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/deckarep/golang-set v1.8.0
	github.com/denverdino/aliyungo v0.0.0-20220321085828-46dabbd9e212
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/gogap/errors v0.0.0-20210818113853-edfbba0ddea9 // indirect
	github.com/gogap/stack v0.0.0-20150131034635-fef68dddd4f8 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/gnostic v0.6.7 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0
	github.com/googleapis/gax-go/v2 v2.2.0 // indirect
	github.com/googleapis/gnostic v0.5.6 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-getter v1.5.11 // indirect
	github.com/hashicorp/go-hclog v1.2.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.4.3 // indirect
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/go-version v1.4.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.11.1 // indirect
	github.com/hashicorp/terraform-config-inspect v0.0.0-20211115214459-90acf1ca460f // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.4.0
	github.com/hashicorp/terraform-plugin-test/v2 v2.2.1 // indirect
	github.com/hashicorp/vault v0.10.4
	github.com/hashicorp/yamux v0.0.0-20211028200310-0bc27b27de87 // indirect
	github.com/jmespath/go-jmespath v0.4.0
	github.com/keybase/go-crypto v0.0.0-20200123153347-de78d2cb44f4 // indirect
	github.com/klauspost/compress v1.15.1 // indirect
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/afero v1.8.2 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/valyala/fasthttp v1.34.0 // indirect
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	github.com/waigani/diffparser v0.0.0-20190828052634-7391f219313d
	github.com/zclconf/go-cty v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	google.golang.org/api v0.73.0 // indirect
	google.golang.org/genproto v0.0.0-20220322021311-435b647f9ef2 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.5
	//k8s.io/client-go v11.0.0+incompatible
	k8s.io/client-go v0.21.0-rc.0
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/kube-openapi v0.0.0-20220322033743-6a7b7046eec8 // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

go 1.13
