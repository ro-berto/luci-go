luci.builder.defaults.task_backend.set("deafult-backend")

luci.project(
    name = "test",
    buildbucket = "cr-buildbucket.appspot.com",
)
luci.bucket(
    name = "ci",
)

luci.task_backend(
    name = "deafult-backend",
    target = "swarming://chromium-swarm-default",
    config = {"key": "value"},
)

luci.builder(
    name = "builder1",
    bucket = "ci",
    executable = luci.recipe(
        name = "recipe",
        cipd_package = "cipd/package",
        cipd_version = "refs/version",
    ),
)

# Expect configs:
#
# === cr-buildbucket.cfg
# buckets {
#   name: "ci"
#   swarming {
#     builders {
#       name: "builder1"
#       backend {
#         target: "swarming://chromium-swarm-default"
#         config_json:
#           '{'
#           '  "key": "value"'
#           '}'
#       }
#       recipe {
#         name: "recipe"
#         cipd_package: "cipd/package"
#         cipd_version: "refs/version"
#       }
#     }
#   }
# }
# ===
#
# === project.cfg
# name: "test"
# ===
#
# === realms.cfg
# realms {
#   name: "@root"
# }
# realms {
#   name: "ci"
# }
# ===
