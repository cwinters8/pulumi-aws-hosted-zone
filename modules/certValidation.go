package modules

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/acm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ConfigureCertValidation(ctx pulumi.Context, region string, certArn pulumi.StringInput, validationRecordFqdn pulumi.StringInput, provider *aws.Provider) (*acm.CertificateValidation, error) {
	// profile := config.GetProfile(&ctx)
	// provider, err := aws.NewProvider(&ctx, region+"-provider", &aws.ProviderArgs{
	// 	Profile: pulumi.StringPtr(profile),
	// 	Region:  pulumi.StringPtr(region),
	// })
	// if err != nil {
	// 	return nil, err
	// }

	certValidation, err := acm.NewCertificateValidation(&ctx, region+"-certValidation", &acm.CertificateValidationArgs{
		CertificateArn:        certArn,
		ValidationRecordFqdns: pulumi.StringArray{validationRecordFqdn},
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}
	return certValidation, nil
}
