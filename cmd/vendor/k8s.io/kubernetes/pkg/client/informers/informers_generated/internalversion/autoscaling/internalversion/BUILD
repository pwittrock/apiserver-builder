package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "horizontalpodautoscaler.go",
        "interface.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/autoscaling/internalversion",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/watch:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/cache:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/apis/autoscaling:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/client/listers/autoscaling/internalversion:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)
