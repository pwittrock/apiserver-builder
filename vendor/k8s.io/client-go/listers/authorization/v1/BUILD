package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "expansion_generated.go",
        "localsubjectaccessreview.go",
        "selfsubjectaccessreview.go",
        "subjectaccessreview.go",
    ],
    importpath = "k8s.io/client-go/listers/authorization/v1",
    tags = ["automanaged"],
    deps = [
        "//vendor/k8s.io/api/authorization/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/labels:go_default_library",
        "//vendor/k8s.io/client-go/tools/cache:go_default_library",
    ],
)
