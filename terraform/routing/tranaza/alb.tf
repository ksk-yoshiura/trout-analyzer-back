resource "aws_lb" "this" {
  count = var.enable_alb ? 1 : 0

  name = "${local.name_prefix}-tranaza"

  internal           = false
  load_balancer_type = "application"

  access_logs {
    bucket  = data.terraform_remote_state.log_alb.outputs.s3_bucket_this_id
    enabled = true
    prefix  = "tranaza"
  }

  security_groups = [
    data.terraform_remote_state.network_main.outputs.security_group_web_id,
    data.terraform_remote_state.network_main.outputs.security_group_vpc_id
  ]

  subnets = [
    for s in data.terraform_remote_state.network_main.outputs.subnet_public : s.id
  ]

  tags = {
    Name = "${local.name_prefix}-tranaza"
  }
}

resource "aws_lb_listener" "https" {
  count = var.enable_alb ? 1 : 0

  certificate_arn   = aws_acm_certificate.root.arn
  load_balancer_arn = aws_lb.this[0].arn
  port              = "443"
  protocol          = "HTTPS"

  ssl_policy = "ELBSecurityPolicy-2016-08"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.tranaza.arn

    // ALBレスポンス確認に便利
    # type = "fixed-response"
    # fixed_response { 
    #   content_type = "text/plain"
    #   message_body = "Fixed response content"
    #   status_code = "200"

    # }
  }

}

resource "aws_lb_listener" "redirect_http_to_https" {
  count = var.enable_alb ? 1 : 0

  load_balancer_arn = aws_lb.this[0].arn
  port              = "80"
  protocol          = "HTTP"


  default_action {
    type = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_lb_target_group" "tranaza" {
  name = "${local.name_prefix}-${local.service_name}"

  deregistration_delay = 60
  port                 = 80
  protocol             = "HTTP"
  target_type          = "ip"
  vpc_id               = data.terraform_remote_state.network_main.outputs.vpc_this_id

  health_check {
    healthy_threshold   = 2
    interval            = 30
    matcher             = 200
    path                = "/"
    port                = "traffic-port"
    protocol            = "HTTP"
    timeout             = 5
    unhealthy_threshold = 2
  }

  tags = {
    Name = "${local.name_prefix}-tranaza"
  }
}