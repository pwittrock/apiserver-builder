package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["expand.go"],
    importpath = "k8s.io/kubernetes/third_party/forked/golang/expansion",
    tags = ["automanaged"],
)

go_test(
    name = "go_default_test",
    srcs = ["expand_test.go"],
    importpath = "k8s.io/kubernetes/third_party/forked/golang/expansion",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = ["//cmd/vendor/k8s.io/kubernetes/pkg/api:go_default_library"],
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
