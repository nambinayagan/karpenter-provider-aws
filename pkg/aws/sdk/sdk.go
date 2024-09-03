package sdk

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type EC2API interface {
	// EC2 Methods
	DescribeImages(ctx context.Context, *ec2.DescribeImagesInput, ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error)
	DescribeLaunchTemplates(ctx context.Context, *ec2.DescribeLaunchTemplatesInput, ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplatesOutput, error)
	DescribeSubnets(ctx context.Context, *ec2.DescribeSubnetsInput, ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error)
	DescribeSecurityGroups(ctx context.Context, *ec2.DescribeSecurityGroupsInput, ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error)
	DescribeInstanceTypes(ctx context.Context, *ec2.DescribeInstanceTypesInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error)
	DescribeInstanceTypeOfferings(ctx context.Context, *ec2.DescribeInstanceTypeOfferingsInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceTypeOfferingsOutput, error)
	DescribeAvailabilityZones(ctx context.Context, *ec2.DescribeAvailabilityZonesInput, ...func(*ec2.Options)) (*ec2.DescribeAvailabilityZonesOutput, error)
	DescribeSpotPriceHistory(ctx context.Context, *ec2.DescribeSpotPriceHistoryInput, ...func(*ec2.Options)) (*ec2.DescribeSpotPriceHistoryOutput, error)
	CreateFleet(ctx context.Context, *ec2.CreateFleetInput, ...func(*ec2.Options)) (*ec2.CreateFleetOutput, error)
	TerminateInstances(ctx context.Context, *ec2.TerminateInstancesInput, ...func(*ec2.Options)) (*ec2.TerminateInstancesOutput, error)
	DescribeInstances(ctx context.Context, *ec2.DescribeInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
	CreateTags(ctx context.Context, *ec2.CreateTagsInput, ...func(*ec2.Options)) (*ec2.CreateTagsOutput, error)
	DescribeInstancesPages(ctx context.Context, input *ec2.DescribeInstancesInput, fn func(*ec2.DescribeInstancesOutput, bool) bool) error
	CreateLaunchTemplate(ctx context.Context, params *ec2.CreateLaunchTemplateInput, optFns ...func(*ec2.Options)) (*ec2.CreateLaunchTemplateOutput, error)
	DeleteLaunchTemplate(ctx context.Context, params *ec2.DeleteLaunchTemplateInput, optFns ...func(*ec2.Options)) (*ec2.DeleteLaunchTemplateOutput, error)
	ModifyLaunchTemplate(ctx context.Context, params *ec2.ModifyLaunchTemplateInput, optFns ...func(*ec2.Options)) (*ec2.ModifyLaunchTemplateOutput, error)
}

type SQSAPI interface {
	// SQS Methods
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
	DeleteMessage(ctx context.Context, *sqs.DeleteMessageInput, ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error)
	GetQueueUrl(ctx context.Context, params *servicesqs.GetQueueUrlInput, optFns ...func(*servicesqs.Options)) (*sqs.GetQueueUrlOutput, error)
	ReceiveMessage(ctx context.Context, params *servicesqs.ReceiveMessageInput, optFns ...func(*servicesqs.Options)) (*sqs.ReceiveMessageOutput, error)
	CreateQueue(ctx context.Context, params *servicesqs.CreateQueueInput, optFns ...func(*servicesqs.Options)) (*sqs.CreateQueueOutput, error)
	DeleteQueue(ctx context.Context, params *servicesqs.DeleteQueueInput, optFns ...func(*servicesqs.Options)) (*sqs.DeleteQueueOutput, error)
}
type IAMAPI interface {
	// IAM Methods
	Reset()
	GetInstanceProfile(ctx context.Context, *iam.GetInstanceProfileInput, ...func(*iam.Options)) (*iam.GetInstanceProfileOutput, error)
	CreateInstanceProfile(ctx context.Context, *iam.CreateInstanceProfileInput, ...func(*iam.Options)) (*iam.CreateInstanceProfileOutput, error)
	DeleteInstanceProfile(ctx context.Context, *iam.DeleteInstanceProfileInput, ...func(*iam.Options)) (*iam.DeleteInstanceProfileOutput, error)
	AddRoleToInstanceProfile(ctx context.Context, *iam.AddRoleToInstanceProfileInput, ...func(*iam.Options)) (*iam.AddRoleToInstanceProfileOutput, error)
	TagInstanceProfile(ctx context.Context, *iam.TagInstanceProfileInput, ...func(*iam.Options)) (*iam.TagInstanceProfileOutput, error)
	RemoveRoleFromInstanceProfile(ctx context.Context, *iam.RemoveRoleFromInstanceProfileInput, ...func(*iam.Options)) (*iam.RemoveRoleFromInstanceProfileOutput, error)
	UntagInstanceProfile(ctx context.Context, params *iam.UntagInstanceProfileInput, optFns ...func(*iam.Options)) (*iam.UntagInstanceProfileOutput, error)
}

type EKSAPI interface {
	// EKS Methods
	DescribeCluster(ctx context.Context, params *eks.DescribeClusterInput, optFns ...func(*eks.Options)) (*eks.DescribeClusterOutput, error)
	Reset()
}

type PricingAPI interface {
	// Pricing Methods
	GetProductsPages(aws.Context, *pricing.GetProductsInput, func(*pricing.GetProductsOutput, bool) bool, ...func(*pricing.Options)) error
}

type SSMAPI interface{
	// SSM Methods
	GetParametersByPathPages(ctx context.Context, params *ssm.GetParametersByPathInput, fn func(*ssm.GetParametersByPathOutput, bool) bool) error
}