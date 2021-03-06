package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = [
        "discovery_client.go",
        "helper.go",
        "restmapper.go",
        "unstructured.go",
    ],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/client-go/discovery",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/emicklei/go-restful-swagger12:go_default_library",
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/googleapis/gnostic/OpenAPIv2:go_default_library",
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/meta:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/serializer:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/version:go_default_library",
        "//cmd/vendor/k8s.io/client-go/kubernetes/scheme:go_default_library",
        "//cmd/vendor/k8s.io/client-go/rest:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "discovery_client_test.go",
        "helper_blackbox_test.go",
        "restmapper_test.go",
    ],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/client-go/discovery_test",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/emicklei/go-restful-swagger12:go_default_library",
        "//cmd/vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/googleapis/gnostic/OpenAPIv2:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/version:go_default_library",
        "//cmd/vendor/k8s.io/client-go/discovery:go_default_library",
        "//cmd/vendor/k8s.io/client-go/kubernetes/scheme:go_default_library",
        "//cmd/vendor/k8s.io/client-go/rest:go_default_library",
        "//cmd/vendor/k8s.io/client-go/rest/fake:go_default_library",
    ],
)
