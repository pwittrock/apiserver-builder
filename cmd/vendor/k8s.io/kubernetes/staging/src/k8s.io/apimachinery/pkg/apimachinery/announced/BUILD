package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["announced_test.go"],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/apimachinery/announced",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = ["//cmd/vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "announced.go",
        "group_factory.go",
    ],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/apimachinery/announced",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/meta:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apimachinery:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apimachinery/registered:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
    ],
)
