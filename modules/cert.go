package modules

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/acm"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/config"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ConfigureCert(ctx pulumi.Context, hostedZoneId pulumi.StringOutput, domain string) (*acm.Certificate, error) {
	profile := config.GetProfile(&ctx)
	eastRegion, err := aws.NewProvider(&ctx, "east", &aws.ProviderArgs{
		Profile: pulumi.StringPtr(profile),
	})
	if err != nil {
		return nil, err
	}

	subjAltName := pulumi.String("*." + domain)
	cert, err := acm.NewCertificate(&ctx, "domainCert", &acm.CertificateArgs{
		DomainName:              pulumi.StringPtr(domain),
		ValidationMethod:        pulumi.StringPtr("DNS"),
		SubjectAlternativeNames: pulumi.StringArray{subjAltName},
	}, pulumi.Provider(eastRegion))
	if err != nil {
		return nil, err
	}

	domainValidationOpts := cert.DomainValidationOptions.Index(pulumi.Int(0))
	validationRecord, err := route53.NewRecord(&ctx, "certValidation", &route53.RecordArgs{
		Name:    domainValidationOpts.ResourceRecordName().Elem(),
		Records: pulumi.StringArray{domainValidationOpts.ResourceRecordValue().Elem()},
		Ttl:     pulumi.Int(60),
		Type:    domainValidationOpts.ResourceRecordType().Elem(),
		ZoneId:  hostedZoneId,
	})
	if err != nil {
		return nil, err
	}

	_, err = acm.NewCertificateValidation(&ctx, "certValidation", &acm.CertificateValidationArgs{
		CertificateArn:        cert.Arn,
		ValidationRecordFqdns: pulumi.StringArray{validationRecord.Fqdn},
	}, pulumi.Provider(eastRegion))
	if err != nil {
		return nil, err
	}

	return cert, nil
}
