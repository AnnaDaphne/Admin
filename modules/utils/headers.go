package utils

import (
    "github.com/astaxie/beego/context"
)

func SecureHeaders(ctx *context.Context) {
    ctx.Output.Header("X-Content-Type-Options", "nosniff")
    ctx.Output.Header("X-FRAME-OPTIONS", "SAMEORIGIN")
    ctx.Output.Header("X-XSS-Protection", "1;mode=block")
}