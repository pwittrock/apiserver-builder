package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["factory_interfaces.go"],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/kube-aggregator/pkg/client/informers/externalversions/internalinterfaces",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/cache:go_default_library",
        "//cmd/vendor/k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset:go_default_library",
    ],
)
