package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["anonymous_test.go"],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/apiserver/pkg/authentication/request/anonymous",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/authentication/authenticator:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/authentication/user:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["anonymous.go"],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/apiserver/pkg/authentication/request/anonymous",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apiserver/pkg/authentication/authenticator:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/authentication/user:go_default_library",
    ],
)
