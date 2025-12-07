#  File: m4docker.inc
#  Last Modification: yy/mm/dd
#  Author: Michael Moscovitch
#  Description:
#  Project:
#  History:


define(PKGINSTALLER,apt-get)
define(PKGINSTALLERFLAGS,-y)
define(PKGUPDATEFLAGS,-y)
define(PKGUPDATECMD,PKGINSTALLER update PKGUPDATEFLAGS)
define(PKGINSTALLCMD,PKGINSTALLER install PKGINSTALLERFLAGS)
define(PKGCLEANCMD,PKGINSTALLER clean)
define(PKGCLEANUP,rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*)

define(PWOSIMAGEBASE,ubuntu)

define(INSTALL_PKG,
RUN PKGUPDATECMD && PKGINSTALLCMD \
$1 \
&& PKGCLEANCMD \
&& PKGCLEANUP
)
