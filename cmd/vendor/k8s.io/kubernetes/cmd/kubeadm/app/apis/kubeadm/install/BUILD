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
        "doc.go",
        "install.go",
    ],
    importpath = "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/install",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apimachinery/pkg/apimachinery/announced:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apimachinery/registered:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api:go_default_library",
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
    srcs = ["install_test.go"],
    importpath = "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/install",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/testing:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/fuzzer:go_default_library",
    ],
)
