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
        "fake_authentication_client.go",
        "fake_tokenreview.go",
        "fake_tokenreview_expansion.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/client/clientset_generated/clientset/typed/authentication/v1beta1/fake",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/api/authentication/v1beta1:go_default_library",
        "//cmd/vendor/k8s.io/client-go/rest:go_default_library",
        "//cmd/vendor/k8s.io/client-go/testing:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/client/clientset_generated/clientset/typed/authentication/v1beta1:go_default_library",
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
