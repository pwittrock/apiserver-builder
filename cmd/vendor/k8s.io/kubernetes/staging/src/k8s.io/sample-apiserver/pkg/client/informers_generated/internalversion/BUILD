package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "factory.go",
        "generic.go",
    ],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/sample-apiserver/pkg/client/informers_generated/internalversion",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/cache:go_default_library",
        "//cmd/vendor/k8s.io/sample-apiserver/pkg/apis/wardle:go_default_library",
        "//cmd/vendor/k8s.io/sample-apiserver/pkg/client/clientset_generated/internalclientset:go_default_library",
        "//cmd/vendor/k8s.io/sample-apiserver/pkg/client/informers_generated/internalversion/internalinterfaces:go_default_library",
        "//cmd/vendor/k8s.io/sample-apiserver/pkg/client/informers_generated/internalversion/wardle:go_default_library",
    ],
)
