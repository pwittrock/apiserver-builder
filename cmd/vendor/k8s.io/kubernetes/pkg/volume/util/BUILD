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
        "atomic_writer.go",
        "device_util.go",
        "device_util_unsupported.go",
        "doc.go",
        "fs_unsupported.go",
        "io_util.go",
        "util.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "fs.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "device_util_linux.go",
            "fs.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "k8s.io/kubernetes/pkg/volume/util",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/api/storage/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/resource:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/labels:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
        "//cmd/vendor/k8s.io/client-go/kubernetes:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api/v1/helper:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/util/mount:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "util_test.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "atomic_writer_test.go",
            "device_util_linux_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "k8s.io/kubernetes/pkg/volume/util",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api/install:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api/v1/helper:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//cmd/vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
            "//cmd/vendor/k8s.io/client-go/util/testing:go_default_library",
        ],
        "//conditions:default": [],
    }),
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
        "//pkg/volume/util/nestedpendingoperations:all-srcs",
        "//pkg/volume/util/operationexecutor:all-srcs",
        "//pkg/volume/util/types:all-srcs",
        "//pkg/volume/util/volumehelper:all-srcs",
    ],
    tags = ["automanaged"],
)
