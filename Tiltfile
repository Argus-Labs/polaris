
# version_settings() enforces a minimum Tilt version
# https://docs.tilt.dev/api.html#api.version_settings
version_settings(constraint='>=0.22.2')

# build the base polard container. This is the base image
# for all the other types of nodes.
custom_build(
  'base',
  'docker build -t $EXPECTED_REF -f ./k8s/base.Dockerfile --build-arg GOOS=linux --build-arg GOARCH=arm64 ./',
  ['./'],
)

docker_build(
  'seed',
  context='.',
  dockerfile='./k8s/seed.Dockerfile',
)
# docker_build(
#     'polard',
#     context='.',
#     dockerfile='./e2e/testapp/docker/local/Dockerfile',
#     build_args={'SERVICE_NAME': 'account's},
#     ssh='default',
#     # only=['./.git/', './services/account/', './pkg/', './Makefile'],
#     # live_update=[
#     #     sync('./services/account/', '/services/account/'),
#     #     sync('./pkg/', '/pkg/'),
#     #     sync('./Makefile', '/Makefile'),
#     #     sync('./.git/', '/.git/'),
#     #     run(
#     #         'go mod tidy',
#     #         trigger=['./services/account/go.mod']
#     #     )
#     # ]
# )

update_settings(suppress_unused_image_warnings=["base"])

# k8s_yaml automatically creates resources in Tilt for the entities
# and will inject any images referenced in the Tiltfile when deploying
# https://docs.tilt.dev/api.html#api.k8s_yaml
k8s_yaml('k8s/account.yaml')

# k8s_resource allows customization where necessary such as adding port forwards and labels
# https://docs.tilt.dev/api.html#api.k8s_resource
k8s_resource(
    'seed',
    port_forwards='8545:8545',
    labels=['seed']
)

# config.main_path is the absolute path to the Tiltfile being run
# there are many Tilt-specific built-ins for manipulating paths, environment variables, parsing JSON/YAML, and more!
# https://docs.tilt.dev/api.html#api.config.main_path
tiltfile_path = config.main_path
