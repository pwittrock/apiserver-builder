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
        "prober.go",
        "prober_manager.go",
        "worker.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/kubelet/prober",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/intstr:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/record:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api/v1/pod:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/container:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/events:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/prober/results:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/status:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/util/format:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/probe:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/probe/exec:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/probe/http:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/probe/tcp:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/util/exec:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "common_test.go",
        "prober_manager_test.go",
        "prober_test.go",
        "worker_test.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/kubelet/prober",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/intstr:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//cmd/vendor/k8s.io/client-go/kubernetes/fake:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/record:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/container:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/container/testing:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/pod:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/prober/results:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/status:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/kubelet/status/testing:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/probe:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/probe/exec:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/util/exec:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//pkg/kubelet/prober/results:all-srcs",
        "//pkg/kubelet/prober/testing:all-srcs",
    ],
    tags = ["automanaged"],
)
