package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = ["csr_test.go"],
    data = glob(["testdata/**"]),
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/client-go/util/cert",
    library = ":go_default_library",
    tags = ["automanaged"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "cert.go",
        "csr.go",
        "io.go",
        "pem.go",
    ],
    data = [
        "testdata/dontUseThisKey.pem",
    ],
    importpath = "k8s.io/kubernetes/staging/src/k8s.io/client-go/util/cert",
    tags = ["automanaged"],
)
