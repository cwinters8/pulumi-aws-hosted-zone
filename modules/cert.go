package modules

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/acm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ConfigureCert(ctx pulumi.Context, hostedZoneId pulumi.StringOutput, domain string, region string, provider *aws.Provider) (*acm.Certificate, error) {
	subjAltName := pulumi.String("*." + domain)
	cert, err := acm.NewCertificate(&ctx, region+"-domainCert", &acm.CertificateArgs{
		DomainName:              pulumi.StringPtr(domain),
		ValidationMethod:        pulumi.StringPtr("DNS"),
		SubjectAlternativeNames: pulumi.StringArray{subjAltName},
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}

	return cert, nil
}
