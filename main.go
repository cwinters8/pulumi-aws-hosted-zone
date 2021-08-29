package main

import (
	"pulumi-aws-hosted-zone/modules"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiConfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := pulumiConfig.New(ctx, "")
		domain := conf.Require("domain")
		zone, err := modules.ConfigureHostedZone(*ctx, domain)
		if err != nil {
			return err
		}

		ctx.Export("hostedZoneId", zone.ID())
		ctx.Export("nameServers", zone.NameServers)
		return nil
	})
}
