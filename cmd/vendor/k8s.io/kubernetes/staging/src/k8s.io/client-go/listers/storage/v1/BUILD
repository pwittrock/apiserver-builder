package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "expansion_generated.go",
        "storageclass.go",
    ],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/client-go/listers/storage/v1",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/api/storage/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/labels:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/cache:go_default_library",
    ],
)
