package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "daemonsets.go",
        "deployments.go",
        "job.go",
        "statefulset.go",
    ],
    importpath = "k8s.io/kubernetes/test/e2e/upgrades/apps",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/onsi/ginkgo:go_default_library",
        "//cmd/vendor/github.com/onsi/gomega:go_default_library",
        "//cmd/vendor/k8s.io/api/apps/v1beta1:go_default_library",
        "//cmd/vendor/k8s.io/api/batch/v1:go_default_library",
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/api/extensions/v1beta1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/labels:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/controller:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/controller/deployment/util:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/util/version:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/test/e2e/framework:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/test/e2e/upgrades:go_default_library",
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
