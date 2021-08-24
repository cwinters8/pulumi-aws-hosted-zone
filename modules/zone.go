package modules

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ConfigureHostedZone(ctx pulumi.Context, domain string) (*route53.Zone, error) {
	zone, err := route53.NewZone(&ctx, "hostedZone", &route53.ZoneArgs{
		Name: pulumi.StringPtr(domain),
	})
	if err != nil {
		return nil, err
	}
	return zone, nil
}
