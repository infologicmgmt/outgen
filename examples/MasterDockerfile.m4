# example docker master file

include(`m4docker.inc')

FROM PWOSIMAGEBASE


ENV DEBIAN_FRONTEND=noninteractive


# install packages
INSTALL_PKG(\
    gcc \
    git \
)

