package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["uuid.go"],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/uuid",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/pborman/uuid:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/types:go_default_library",
    ],
)
