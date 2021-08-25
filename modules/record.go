package modules

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ConfigureRecord(ctx pulumi.Context, resourceName string, name pulumi.StringInput, records pulumi.StringArray, ttl pulumi.Int, recordType pulumi.StringInput, hostedZoneId pulumi.StringInput) (*route53.Record, error) {
	record, err := route53.NewRecord(&ctx, resourceName, &route53.RecordArgs{
		Name:    name,
		Records: records,
		Ttl:     ttl,
		Type:    recordType,
		ZoneId:  hostedZoneId,
	})
	if err != nil {
		return nil, err
	}
	return record, nil
}
