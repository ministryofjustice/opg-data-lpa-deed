module "eu_west_1" {
  source = "./region"

  app_version      = var.app_version
  dynamodb_arn     = aws_dynamodb_table.deeds_table.arn
  environment_name = local.environment_name

  providers = {
    aws.region     = aws.eu_west_1
    aws.management = aws.management_eu_west_1
  }
}

module "eu_west_2" {
  source = "./region"

  app_version      = var.app_version
  dynamodb_arn     = aws_dynamodb_table_replica.deeds_table.arn
  environment_name = local.environment_name

  providers = {
    aws.region     = aws.eu_west_2
    aws.management = aws.management_eu_west_2
  }
}
