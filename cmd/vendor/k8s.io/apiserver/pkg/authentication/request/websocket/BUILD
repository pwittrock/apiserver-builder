package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["protocol.go"],
    importpath = "k8s.io/apiserver/pkg/authentication/request/websocket",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apiserver/pkg/authentication/authenticator:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/authentication/user:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/util/wsstream:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["protocol_test.go"],
    importpath = "k8s.io/apiserver/pkg/authentication/request/websocket",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apiserver/pkg/authentication/authenticator:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/authentication/user:go_default_library",
    ],
)
