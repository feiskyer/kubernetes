load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "azure_client_config.go",
        "doc.go",
    ],
    importmap = "k8s.io/kubernetes/vendor/k8s.io/legacy-cloud-providers/azure/clients",
    importpath = "k8s.io/legacy-cloud-providers/azure/clients",
    visibility = ["//visibility:public"],
    deps = [
        "//staging/src/k8s.io/client-go/util/flowcontrol:go_default_library",
        "//staging/src/k8s.io/legacy-cloud-providers/azure/retry:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/adal:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//staging/src/k8s.io/legacy-cloud-providers/azure/clients/armclient:all-srcs",
        "//staging/src/k8s.io/legacy-cloud-providers/azure/clients/vmssclient:all-srcs",
        "//staging/src/k8s.io/legacy-cloud-providers/azure/clients/vmssvmclient:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
