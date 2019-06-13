# `go_repository` + go modules + vendor directory with added BUILD files

This repository contains a minimal example for a specific issue in the interplay of the aforementioned components.

## Steps to reproduce

Run `bazel build @com_github_envoyproxy_protoc_gen_validate//:protoc-gen-validate`. Observe the build failing:

```
ERROR: /home/jannis/.cache/bazel/_bazel_jannis/cafebabedeadbeef/external/com_github_envoyproxy_protoc_gen_validate/BUILD:16:1: no such package '@com_github_envoyproxy_protoc_gen_validate//vendor/github.com/lyft/protoc-gen-star/lang/go': BUILD file not found on package path and referenced by '@com_github_envoyproxy_protoc_gen_validate//:go_default_library'
ERROR: Analysis of target '@com_github_envoyproxy_protoc_gen_validate//:protoc-gen-validate' failed; build aborted: no such package '@com_github_envoyproxy_protoc_gen_validate//vendor/github.com/lyft/protoc-gen-star/lang/go': BUILD file not found on package path
INFO: Elapsed time: 80.201s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (14 packages loaded, 92 targets configured)
```

This problem starts occuring with https://github.com/bazelbuild/bazel-gazelle/commit/b661aec20274d7b8201fcff9d550564e3451c8ef because `build_file_generation="auto"` is fixed with this commit. The problem doesn't occur with `"always"`

## Leads

Apparently `go mod download` strips vendor directories from downloaded modules. The [module in question](https://github.com/envoyproxy/protoc-gen-validate), however, contains custom-made BUILD files in its vendor directory. The packages declared by those are referenced in its top-level BUILD file, which is not stripped. Hence, the downloaded module cannot be built with Bazel.

