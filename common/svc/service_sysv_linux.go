// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

package svc

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"
)

type sysv struct {
	i Interface
	*Config
}

func newSystemVService(i Interface, c *Config) (Service, error) {
	s := &sysv{
		i:      i,
		Config: c,
	}

	return s, nil
}

func (s *sysv) String() string {
	if len(s.DisplayName) > 0 {
		return s.DisplayName
	}
	return s.Name
}

var errNoUserServiceSystemV = errors.New("User services are not supported on SystemV.")

func (s *sysv) configPath() (cp string, err error) {
	if s.Option.bool(optionUserService, optionUserServiceDefault) {
		err = errNoUserServiceSystemV
		return
	}
	cp = "/etc/init.d/" + s.Config.Name
	return
}
func (s *sysv) template() *template.Template {
	return template.Must(template.New("").Funcs(tf).Parse(sysvScript))
}

func (s *sysv) Install() error {
	confPath, err := s.configPath()
	if err != nil {
		return err
	}
	_, err = os.Stat(confPath)
	if err == nil {
		return fmt.Errorf("Init already exists: %s", confPath)
	}

	f, err := os.Create(confPath)
	if err != nil {
		return err
	}
	defer f.Close()

	path, err := s.execPath()
	if err != nil {
		return err
	}

	var to = &struct {
		*Config
		Path string
	}{
		s.Config,
		path,
	}

	err = s.template().Execute(f, to)
	if err != nil {
		return err
	}

	if err = os.Chmod(confPath, 0755); err != nil {
		return err
	}
	for _, i := range [...]string{"2", "3", "4", "5"} {
		if err = os.Symlink(confPath, "/etc/rc"+i+".d/S50"+s.Name); err != nil {
			continue
		}
	}
	for _, i := range [...]string{"0", "1", "6"} {
		if err = os.Symlink(confPath, "/etc/rc"+i+".d/K02"+s.Name); err != nil {
			continue
		}
	}

	return nil
}

func (s *sysv) Uninstall() error {
	cp, err := s.configPath()
	if err != nil {
		return err
	}
	if err := os.Remove(cp); err != nil {
		return err
	}
	return nil
}

func (s *sysv) Logger(errs chan<- error) (Logger, error) {
	if system.Interactive() {
		return ConsoleLogger, nil
	}
	return s.SystemLogger(errs)
}
func (s *sysv) SystemLogger(errs chan<- error) (Logger, error) {
	return newSysLogger(s.Name, errs)
}

func (s *sysv) Run() (err error) {
	err = s.i.Start(s)
	if err != nil {
		return err
	}

	s.Option.funcSingle(optionRunWait, func() {
		var sigChan = make(chan os.Signal, 3)
		signal.Notify(sigChan, syscall.SIGTERM, os.Interrupt)
		<-sigChan
	})()

	return s.i.Stop(s)
}

func (s *sysv) Start() error {
	return run("service", s.Name, "start")
}

func (s *sysv) Stop() error {
	return run("service", s.Name, "stop")
}

func (s *sysv) Restart() error {
	err := s.Stop()
	if err != nil {
		return err
	}
	time.Sleep(50 * time.Millisecond)
	return s.Start()
}

const sysvScript = `#!/bin/sh
### BEGIN INIT INFO
# Provides:          NPD
# Required-Start:    $network $local_fs $remote_fs $syslog
# Required-Stop:     $remote_fs
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: {{.Description}}
# Description:       {{.Description}}
### END INIT INFO

. /etc/rc.d/init.d/functions

PATH=/sbin:/usr/sbin:/bin:/usr/bin
DESC="{{.DisplayName}}"
NAME="{{.Name}}"
DAEMON="{{.Path}}{{range .Arguments}} {{.|cmd}}{{end}}"
PIDFILE=/var/lib/$NAME/$NAME.pid
SCRIPTNAME=/etc/init.d/$NAME
USER={{.UserName}}
LOCKFILE=/var/lib/$NAME/$NAME.lock

start() {
    echo -n $"Starting $NAME: "
    daemon --user $USER --pidfile $PIDFILE "$DAEMON &>/dev/null & echo \$! > $PIDFILE"
    retval=$?
    echo
    [ $retval -eq 0 ] && touch $LOCKFILE
    return $retval
}

stop() {
    echo -n $"Stopping $NAME: "
    killproc -p $PIDFILE $NAME
    retval=$?
    echo
    [ $retval -eq 0 ] && rm -f $LOCKFILE
    return $retval
}

restart() {
    stop
    start
}

reload() {
    restart
}

force_reload() {
    restart
}

rh_status() {
    status -p $PIDFILE $NAME
}

rh_status_q() {
    rh_status >/dev/null 2>&1
}

case "$1" in
    start)
        rh_status_q && exit 0
        $1
        ;;
    stop)
        rh_status_q || exit 0
        $1
        ;;
    restart)
        $1
        ;;
    reload)
        rh_status_q || exit 7
        $1
        ;;
    force-reload)
        force_reload
        ;;
    status)
        rh_status
        ;;
    condrestart|try-restart)
        rh_status_q || exit 0
        restart
        ;;
    *)
        echo $"Usage: $0 {start|stop|status|restart|condrestart|try-restart|reload|force-reload}"
        exit 2
esac
exit $?
`
