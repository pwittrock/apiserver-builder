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
        "format.go",
        "metrics.go",
        "request.go",
        "scheme.go",
        "types.go",
        "union.go",
    ],
    importpath = "k8s.io/apiserver/pkg/audit",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/pborman/uuid:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//cmd/vendor/k8s.io/api/authentication/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/serializer:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/errors:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/net:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/apis/audit:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/apis/audit/v1alpha1:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/authorization/authorizer:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["union_test.go"],
    importpath = "k8s.io/apiserver/pkg/audit",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/apis/audit:go_default_library",
    ],
)
