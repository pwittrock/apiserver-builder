package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["patch_test.go"],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/jsonmergepatch",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/davecgh/go-spew/spew:go_default_library",
        "//vendor/github.com/evanphx/json-patch:go_default_library",
        "//vendor/github.com/ghodss/yaml:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/json:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["patch.go"],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/jsonmergepatch",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/evanphx/json-patch:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/json:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/mergepatch:go_default_library",
    ],
)
