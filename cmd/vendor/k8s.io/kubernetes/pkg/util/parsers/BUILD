package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["parsers.go"],
    importpath = "k8s.io/kubernetes/pkg/util/parsers",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/docker/distribution/reference:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["parsers_test.go"],
    importpath = "k8s.io/kubernetes/pkg/util/parsers",
    library = ":go_default_library",
    tags = ["automanaged"],
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
