package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["proxy_test.go"],
    importpath = "k8s.io/apiserver/pkg/util/proxy",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/intstr:go_default_library",
        "//cmd/vendor/k8s.io/client-go/listers/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/cache:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["proxy.go"],
    importpath = "k8s.io/apiserver/pkg/util/proxy",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/intstr:go_default_library",
        "//cmd/vendor/k8s.io/client-go/listers/core/v1:go_default_library",
    ],
)
