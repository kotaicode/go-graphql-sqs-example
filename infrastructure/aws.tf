provider "aws" {
  region = var.aws_region
}

resource "aws_sqs_queue" "go-sqs-queue" {
  name                      = var.aws_sqs_queue_name
  delay_seconds             = 0
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 0
}

resource "aws_iam_policy" "go-sqs-queue-policy" {
  name = "sqs-${var.aws_sqs_queue_name}-policy"
  policy = data.aws_iam_policy_document.go-sqs-queue-policy-doc.json
}

data "aws_iam_policy_document" "go-sqs-queue-policy-doc" {
  statement {
    actions = [
      "sqs:DeleteMessage",
      "sqs:GetQueueUrl",
      "sqs:ChangeMessageVisibility",
      "sqs:SendMessageBatch",
      "sqs:ReceiveMessage",
      "sqs:SendMessage",
      "sqs:GetQueueAttributes",
      "sqs:ListQueueTags",
      "sqs:ListDeadLetterSourceQueues",
      "sqs:DeleteMessageBatch",
      "sqs:PurgeQueue",
      "sqs:DeleteQueue",
      "sqs:CreateQueue",
      "sqs:ChangeMessageVisibilityBatch",
      "sqs:SetQueueAttributes"
    ]
    resources = [
      aws_sqs_queue.go-sqs-queue.arn
    ]
  }
}

resource "aws_iam_user" "go-sqs-iam-user" {
  name = var.aws_iam_user
}

resource "aws_iam_access_key" "go-sqs-iam-user" {
  user = aws_iam_user.go-sqs-iam-user.name
}

output "secret" {
  value = "${aws_iam_access_key.go-sqs-iam-user.secret}"
}

resource "aws_iam_user_policy_attachment" "go-sqs-iam-user-attachment" {
  user = aws_iam_user.go-sqs-iam-user.name
  policy_arn = aws_iam_policy.go-sqs-queue-policy.arn
}

resource "aws_iam_role" "go-sqs-iam-role" {
  name = "go-graphql-sqs-example-role"
  assume_role_policy = data.aws_iam_policy_document.go-sqs-iam-role-policy.json
}

data "aws_iam_policy_document" "go-sqs-iam-role-policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "AWS"
      identifiers = [aws_iam_user.go-sqs-iam-user.arn]
    }
  }
}

resource "aws_iam_role_policy_attachment" "go-sqs-iam-role-attach" {
  role = aws_iam_role.go-sqs-iam-role.name
  policy_arn = aws_iam_policy.go-sqs-queue-policy.arn
}
