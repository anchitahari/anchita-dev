terraform {
    required_providers {
        aws = {
            source  = "hashicorp/aws"
            version = "~> 5.0"
        }
    }
}

provider "aws" {
    region = "us-east-1"
}

data "aws_availability_zones" "available" {}

# VPC
module "vpc" {
    source  = "terraform-aws-modules/vpc/aws"
    version = "5.0.0"

    name = "anchita-dev-vpc"
    cidr = "10.0.0.0/16"

    azs             = slice(data.aws_availability_zones.available.names, 0, 2)
    private_subnets = ["10.0.1.0/24", "10.0.2.0/24"]
    public_subnets  = ["10.0.4.0/24", "10.0.5.0/24"]

    enable_nat_gateway   = true
    single_nat_gateway   = true
    enable_dns_hostnames = true

    public_subnet_tags = {
        "kubernetes.io/role/elb" = 1
    }

    private_subnet_tags = {
        "kubernetes.io/role/internal-elb" = 1
    }
}

# EKS cluster
module "eks" {
    source  = "terraform-aws-modules/eks/aws"
    version = "19.15.3"

    cluster_name    = "anchita-dev"
    cluster_version = "1.30"

    vpc_id                         = module.vpc.vpc_id
    subnet_ids                     = module.vpc.private_subnets
    cluster_endpoint_public_access = true

    eks_managed_node_groups = {
        main = {
            min_size     = 1
            max_size     = 2
            desired_size = 1
            instance_types = ["t3.small"]
        }
    }
}

output "cluster_name" {
    value = module.eks.cluster_name
}

output "cluster_endpoint" {
    value = module.eks.cluster_endpoint
}