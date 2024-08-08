resource "aws_secretsmanager_secret" "jwt_key" {
  name        = "${data.aws_default_tags.default.tags.application}/${data.aws_default_tags.default.tags.account}/jwt-key"
  description = "JWT key for ${data.aws_default_tags.default.tags.application} in ${data.aws_default_tags.default.tags.account}, for use with Make and Register, and Use a LPA"
  # policy      = data.aws_iam_policy_document.jwt_key_cross_account_access.json
  # kms_key_id  = module.jwt_kms.eu_west_1_target_key_id
  replica {
    region = data.aws_region.eu_west_2.name
    # kms_key_id = module.jwt_kms.eu_west_2_target_key_id
  }
  provider = aws.management_eu_west_1
}

data "aws_iam_policy_document" "jwt_key_cross_account_access" {
  statement {
    effect = "Allow"
    actions = [
      "secretsmanager:GetSecretValue",
    ]
    resources = [
      "*"
    ]
    principals {
      type        = "AWS"
      identifiers = tolist(local.account.jwt_key_cross_account_access_roles)
    }
  }
}
