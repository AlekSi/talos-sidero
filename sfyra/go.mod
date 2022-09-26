module github.com/talos-systems/sidero/sfyra

go 1.16

replace (
	github.com/talos-systems/sidero => ../

	github.com/talos-systems/talos/pkg/machinery => github.com/talos-systems/talos/pkg/machinery v0.11.5

	// See https://github.com/talos-systems/go-loadbalancer/pull/4
	// `go get github.com/smira/tcpproxy@combined-fixes`, then copy pseudo-version there
	inet.af/tcpproxy => github.com/smira/tcpproxy v0.0.0-20201015133617-de5f7797b95b

	// keep older versions of k8s.io packages to keep compatiblity with cluster-api
	k8s.io/api v0.21.3 => k8s.io/api v0.20.5
	k8s.io/apimachinery v0.21.3 => k8s.io/apimachinery v0.20.5
	k8s.io/client-go v0.21.3 => k8s.io/client-go v0.20.5

	sigs.k8s.io/cluster-api v0.3.20 => sigs.k8s.io/cluster-api v0.3.9
)

require (
	github.com/emicklei/go-restful v2.15.0+incompatible // indirect
	github.com/mdlayher/raw v0.0.0-20210412142147-51b895745faf // indirect
	github.com/spf13/cobra v1.5.0
	github.com/stretchr/testify v1.8.0
	github.com/talos-systems/cluster-api-bootstrap-provider-talos v0.2.0
	github.com/talos-systems/cluster-api-control-plane-provider-talos v0.1.1
	github.com/talos-systems/go-debug v0.2.1
	github.com/talos-systems/go-loadbalancer v0.1.3
	github.com/talos-systems/go-procfs v0.1.0
	github.com/talos-systems/go-retry v0.3.1
	github.com/talos-systems/net v0.3.2
	github.com/talos-systems/sidero v0.0.0-00010101000000-000000000000
	github.com/talos-systems/talos v1.2.3
	github.com/talos-systems/talos/pkg/machinery v1.2.3
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.25.1
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.25.1
	k8s.io/client-go v0.25.1
	sigs.k8s.io/cluster-api v0.3.20
	sigs.k8s.io/controller-runtime v0.6.3
)
