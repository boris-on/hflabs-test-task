package pprofing

import (
	"context"
	"net/http"
	"net/http/pprof"

	"github.com/sirupsen/logrus"
)

func InitPprofing(ctx context.Context, log logrus.FieldLogger, host string) {
	log.Infof("pprofing at %s", host)
	r := http.NewServeMux()

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	if err := http.ListenAndServe(host, r); err != nil {
		log.Errorf(err.Error())
	}
}
