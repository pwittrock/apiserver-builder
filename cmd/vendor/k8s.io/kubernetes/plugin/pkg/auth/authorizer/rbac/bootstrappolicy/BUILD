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
        "controller_policy.go",
        "namespace_policy.go",
        "policy.go",
    ],
    importpath = "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac/bootstrappolicy",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/meta:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apiserver/pkg/authentication/user:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/apis/rbac:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["policy_test.go"],
    data = glob([
        "testdata/*",
    ]),
    importpath = "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac/bootstrappolicy_test",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/ghodss/yaml:go_default_library",
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/api/rbac/v1beta1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/meta:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/diff:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/api/install:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/apis/rbac:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/apis/rbac/install:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/registry/rbac/validation:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac/bootstrappolicy:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["controller_policy_test.go"],
    data = glob(["testdata/**"]),
    importpath = "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac/bootstrappolicy",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apimachinery/pkg/api/meta:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
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
