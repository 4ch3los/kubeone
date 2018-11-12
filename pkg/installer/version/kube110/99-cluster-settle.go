package kube110

import (
	"time"

	"github.com/kubermatic/kubeone/pkg/installer/util"
)

func Wait(ctx *util.Context, t time.Duration) error {
	ctx.Logger.Infoln("Letting the cluster settle down…")
	time.Sleep(t)

	return nil
}
