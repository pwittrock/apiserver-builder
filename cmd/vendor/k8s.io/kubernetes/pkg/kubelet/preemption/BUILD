package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["preemption.go"],
    importpath = "k8s.io/kubernetes/pkg/kubelet/preemption",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/util/feature:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/record:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api/v1/helper/qos:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api/v1/resource:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/features:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/events:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/eviction:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/lifecycle:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/types:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/util/format:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/plugin/pkg/scheduler/algorithm:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/plugin/pkg/scheduler/algorithm/predicates:go_default_library",
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
    srcs = ["preemption_test.go"],
    importpath = "k8s.io/kubernetes/pkg/kubelet/preemption",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/resource:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/record:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/types:go_default_library",
    ],
)
