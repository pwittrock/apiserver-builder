package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["crdregistration_controller.go"],
    importpath = "k8s.io/kubernetes/pkg/master/controller/crdregistration",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions:go_default_library",
        "//cmd/vendor/k8s.io/apiextensions-apiserver/pkg/client/informers/internalversion/apiextensions/internalversion:go_default_library",
        "//cmd/vendor/k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/internalversion:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/labels:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/cache:go_default_library",
        "//cmd/vendor/k8s.io/client-go/util/workqueue:go_default_library",
        "//cmd/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration:go_default_library",
        "//cmd/vendor/k8s.io/kubernetes/pkg/controller:go_default_library",
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

go_test(
    name = "go_default_test",
    srcs = ["crdregistration_controller_test.go"],
    importpath = "k8s.io/kubernetes/pkg/master/controller/crdregistration",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions:go_default_library",
        "//cmd/vendor/k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/internalversion:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//cmd/vendor/k8s.io/client-go/tools/cache:go_default_library",
        "//cmd/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration:go_default_library",
    ],
)
