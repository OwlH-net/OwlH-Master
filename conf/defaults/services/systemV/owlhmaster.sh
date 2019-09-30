# # OWLH MASTER API

# mkdir -p /usr/local/owlh/src/bin
# mkdir -p /tmp/dispatcher/pcaps/
# touch /usr/local/owlh/src/owlhmaster/conf/stopdispatcher

# cat >> /etc/inittab << \EOF
# owlhmaster:2345:respawn:/usr/local/owlh/bin/manage-owlhmaster.sh
# EOF


# cat > /etc/init.d/owlhmaster <<\EOF
# #!/bin/sh

# # Copyright (C) 2019, OwlH.net
# # OwlH master: controls owlh master service
# # Author:       owlh team <support@olwh.net>
# #
# # chkconfig: 2345 99 15
# # description: starts and stop OwlH master service
# #

# # Source function library.
# export LANG=C

# start() {
#     echo -n "Starting OwlH MASTER: "
#     /usr/local/owlh/bin/manage-owlhmaster.sh start > /dev/null
#     RETVAL=$?
#     if [ $RETVAL -eq 0 ]; then
#         echo "success"
#     else
#         echo "failure"
#     fi
#     echo
#     return $RETVAL
# }

# stop() {
#     echo -n "Stopping OwlH MASTER: "
#     /usr/local/owlh/bin/manage-owlhmaster.sh stop > /dev/null
#     RETVAL=$?
#     if [ $RETVAL -eq 0 ]; then
#         echo "success"
#     else
#         echo "failure"
#     fi
#     echo
#     return $RETVAL
# }

# status() {
#     /usr/local/owlh/bin/manage-owlhmaster.sh status
#     RETVAL=$?
#     return $RETVAL
# }

# case "$1" in
# start)
#     start
#     ;;
# stop)
#     stop
#     ;;
# restart)
#     stop
#     start
#     ;;
# status)
#     status
#     ;;
# *)
#     echo "*** Usage: owlhmaster {start|stop|restart|status}"
#     exit 1
# esac

# exit $?
# EOF

# cat > /usr/local/owlh/bin/manage-owlhmaster.sh <<\EOF
# #!/bin/bash
# case "$1" in
# start)
#   cd /usr/local/owlh/src/owlhmaster/
#   export GOPATH=/usr/local/owlh
#   /usr/local/owlh/src/owlhmaster/owlhmaster > /dev/null 2>&1 &
#   ;;
# stop)
#   if [ $(pidof owlhmaster) ] ; then
#      kill -9 $(pidof owlhmaster)
#   fi
#   ;;
# status)
#   echo -n "OwlH MASTER status -> "
#   if [ $(pidof owlhmaster) ] ; then
#      echo -e "\e[92mRunning\e[0m"
#      exit 0
#   else
#      echo -e "\e[91mNot running\e[0m"
#      exit 1
#   fi
# esac

# exit $?
# EOF

# chmod +x /etc/init.d/owlhmaster
# chmod +x /usr/local/owlh/bin/manage-owlhmaster.sh
# chkconfig --add owlhmaster
# chkconfig owlhmaster on
# service owlhmaster start