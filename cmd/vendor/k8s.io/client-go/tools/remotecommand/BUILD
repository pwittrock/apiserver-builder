package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = [
        "v2_test.go",
        "v4_test.go",
    ],
    importpath = "k8s.io/client-go/tools/remotecommand",
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/httpstream:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "errorstream.go",
        "remotecommand.go",
        "resize.go",
        "v1.go",
        "v2.go",
        "v3.go",
        "v4.go",
    ],
    importpath = "k8s.io/client-go/tools/remotecommand",
    tags = ["automanaged"],
    deps = [
        "//cmd/vendor/github.com/golang/glog:go_default_library",
        "//cmd/vendor/k8s.io/api/core/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/httpstream:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/remotecommand:go_default_library",
        "//cmd/vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//cmd/vendor/k8s.io/client-go/rest:go_default_library",
        "//cmd/vendor/k8s.io/client-go/transport:go_default_library",
        "//cmd/vendor/k8s.io/client-go/transport/spdy:go_default_library",
        "//cmd/vendor/k8s.io/client-go/util/exec:go_default_library",
    ],
)
