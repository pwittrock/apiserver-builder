package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "fake_image_service.go",
        "fake_runtime_service.go",
        "utils.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/kubelet/apis/cri/testing",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/util/sliceutils:go_default_library",
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
