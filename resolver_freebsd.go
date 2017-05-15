package libnetwork

import (
	"fmt"
	//"net"
	"os"
	"os/exec"
	//"runtime"

	//"github.com/Sirupsen/logrus"
	"github.com/docker/docker/pkg/reexec"
	//"github.com/docker/libnetwork/iptables"
	//"github.com/vishvananda/netns"
)

func init() {
	reexec.Register("setup-resolver", reexecSetupResolver)
}

const (
	// outputChain used for docker embed dns
	outputChain = "DOCKER_OUTPUT"
	//postroutingchain used for docker embed dns
	postroutingchain = "DOCKER_POSTROUTING"
)

func reexecSetupResolver() {
}

func (r *resolver) setupIPTable() error {
	if r.err != nil {
		return r.err
	}
	laddr := r.conn.LocalAddr().String()
	ltcpaddr := r.tcpListen.Addr().String()

	cmd := &exec.Cmd{
		Path:   reexec.Self(),
		Args:   append([]string{"setup-resolver"}, r.resolverKey, laddr, ltcpaddr),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("reexec failed: %v", err)
	}
	return nil
}
