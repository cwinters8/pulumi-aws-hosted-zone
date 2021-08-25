package main

import (
	"pulumi-aws-hosted-zone/modules"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/acm"
	awsConfig "github.com/pulumi/pulumi-aws/sdk/v4/go/aws/config"
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

		type Certificate struct {
			region   string
			acmCert  *acm.Certificate
			provider *aws.Provider
		}

		var certs []Certificate
		// TODO: pass a list of cert regions via config
		regions := []string{"us-east-1", "us-west-2"}
		for _, v := range regions {
			profile := awsConfig.GetProfile(ctx)
			provider, err := aws.NewProvider(ctx, v+"-provider", &aws.ProviderArgs{
				Profile: pulumi.StringPtr(profile),
				Region:  pulumi.StringPtr(v),
			})
			if err != nil {
				return err
			}
			cert, err := modules.ConfigureCert(*ctx, zone.ZoneId, domain, v, provider)
			if err != nil {
				return err
			}
			ctx.Export(v+"-certificateArn", cert.Arn)
			certs = append(certs, Certificate{acmCert: cert, region: v, provider: provider})
		}
		domainValidationOpts := certs[0].acmCert.DomainValidationOptions.Index(pulumi.Int(0))
		recordName := domainValidationOpts.ResourceRecordName().Elem()
		records := pulumi.StringArray{domainValidationOpts.ResourceRecordValue().Elem()}
		recordType := domainValidationOpts.ResourceRecordType().Elem()
		validationRecord, err := modules.ConfigureRecord(*ctx, domain+"-record", recordName, records, pulumi.Int(60), recordType, zone.ZoneId)
		if err != nil {
			return err
		}
		for _, v := range certs {
			_, err := modules.ConfigureCertValidation(*ctx, v.region, v.acmCert.Arn, validationRecord.Fqdn, v.provider)
			if err != nil {
				return err
			}
		}

		ctx.Export("hostedZoneId", zone.ID())
		ctx.Export("nameServers", zone.NameServers)
		return nil
	})
}
