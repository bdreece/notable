package options

import "flag"

var opts Options

type Options struct {
	Config       string
	HttpPort     int
	GrpcPort     int
}

func Parse() *Options {
    if flag.Parsed() {
        return &opts
    }

    flag.StringVar(&opts.Config, "c", "configs/server.yml", "config path")
	flag.IntVar(&opts.HttpPort, "H", 3000, "http port")
    flag.IntVar(&opts.GrpcPort, "G", 3001, "grpc port")
    flag.Parse()

    return &opts
}
