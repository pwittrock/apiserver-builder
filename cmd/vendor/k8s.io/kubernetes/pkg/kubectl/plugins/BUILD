package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = [
        "env.go",
        "loader.go",
        "plugins.go",
        "runner.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/kubectl/plugins",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/ghodss/yaml:go_default_library",
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/github.com/spf13/pflag:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/clientcmd:go_default_library",
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

go_test(
    name = "go_default_test",
    srcs = [
        "env_test.go",
        "loader_test.go",
        "plugins_test.go",
        "runner_test.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/kubectl/plugins",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = ["//cmd/vendor/github.com/spf13/pflag:go_default_library"],
)
