package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["kubeconfig_test.go"],
    importpath = "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig",
    library = ":go_default_library",
    tags = ["automanaged"],
)

go_library(
    name = "go_default_library",
    srcs = ["kubeconfig.go"],
    importpath = "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/client-go/kubernetes:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/clientcmd:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/clientcmd/api:go_default_library",
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
