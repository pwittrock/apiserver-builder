package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["portforward_test.go"],
    importpath = "k8s.io/client-go/tools/portforward",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = ["//cmd/vendor/k8s.io/apimachinery/pkg/util/httpstream:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "portforward.go",
    ],
    importpath = "k8s.io/client-go/tools/portforward",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/httpstream:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
    ],
)
