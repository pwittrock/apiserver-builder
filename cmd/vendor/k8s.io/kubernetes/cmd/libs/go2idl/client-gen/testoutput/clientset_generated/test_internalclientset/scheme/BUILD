package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "register.go",
    ],
    importpath = "k8s.io/kubernetes/cmd/libs/go2idl/client-gen/testoutput/clientset_generated/test_internalclientset/scheme",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apimachinery/pkg/apimachinery/announced:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apimachinery/registered:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/serializer:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/cmd/libs/go2idl/client-gen/test_apis/testgroup/install:go_default_library",
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
