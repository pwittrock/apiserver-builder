package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["storage_test.go"],
    importpath = "k8s.io/kubernetes/pkg/registry/core/service/allocator/storage",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//vendor/golang.org/x/net/context:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/storage/etcd/testing:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/storage/storagebackend:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/registry/core/service/allocator:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/registry/registrytest:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["storage.go"],
    importpath = "k8s.io/kubernetes/pkg/registry/core/service/allocator/storage",
    tags = ["automanaged"],
    deps = [
        "//vendor/golang.org/x/net/context:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/registry/generic:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/storage:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/storage/errors:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/storage/storagebackend:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/registry/core/rangeallocation:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/registry/core/service/allocator:go_default_library",
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
