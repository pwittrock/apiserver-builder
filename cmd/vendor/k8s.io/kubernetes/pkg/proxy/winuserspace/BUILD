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
        "loadbalancer.go",
        "proxier.go",
        "proxysocket.go",
        "roundrobin.go",
        "types.go",
        "udp_server.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/proxy/winuserspace",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/miekg/dns:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/net:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api/helper:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/proxy:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/util/exec:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/util/ipconfig:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/util/netsh:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/util/slice:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "proxier_test.go",
        "roundrobin_test.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/proxy/winuserspace",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/proxy:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/util/netsh/testing:go_default_library",
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
