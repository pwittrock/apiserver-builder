package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["file.go"],
    importpath = "k8s.io/kubernetes/cmd/kubeadm/app/discovery/file",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/clientcmd:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/clientcmd/api:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/cmd/kubeadm/app/constants:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/bootstrap/api:go_default_library",
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
