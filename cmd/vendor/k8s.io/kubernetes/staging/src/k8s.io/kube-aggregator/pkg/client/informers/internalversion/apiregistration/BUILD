package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["interface.go"],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/kube-aggregator/pkg/client/informers/internalversion/apiregistration",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/kube-aggregator/pkg/client/informers/internalversion/apiregistration/internalversion:go_default_library",
        "//cmd/vendor/k8s.io/kube-aggregator/pkg/client/informers/internalversion/internalinterfaces:go_default_library",
    ],
)
