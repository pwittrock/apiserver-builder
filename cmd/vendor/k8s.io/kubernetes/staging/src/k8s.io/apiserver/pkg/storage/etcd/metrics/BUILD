package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["metrics.go"],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/apiserver/pkg/storage/etcd/metrics",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/prometheus/client_golang/prometheus:go_default_library"],
)
