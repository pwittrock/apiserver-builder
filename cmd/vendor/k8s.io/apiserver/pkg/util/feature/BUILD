package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["feature_gate_test.go"],
    importpath = "k8s.io/apiserver/pkg/util/feature",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = ["//cmd/vendor/github.com/spf13/pflag:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = ["feature_gate.go"],
    importpath = "k8s.io/apiserver/pkg/util/feature",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/github.com/spf13/pflag:go_default_library",
    ],
)
