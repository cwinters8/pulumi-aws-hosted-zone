package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		config := config.New(ctx, "")
		domain := config.Require("domain")
		zone, err := route53.NewZone(ctx, "hostedZone", &route53.ZoneArgs{
			Name: pulumi.StringPtr(domain),
		})
		if err != nil {
			return err
		}

		ctx.Export("hostedZoneId", zone.ID())
		ctx.Export("nameServers", zone.NameServers)
		return nil
	})
}
