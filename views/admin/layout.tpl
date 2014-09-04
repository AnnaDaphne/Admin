<!doctype html>
<html class="no-js" lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">

        <title>Anna &amp; Daphne &mdash; {{.Title}}</title>
        <meta name="description" content="">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        {{range .HeadStyles}}
            <link rel="stylesheet" href="{{.}}">
        {{end}}

        {{range .HeadScripts}}
            <script src="{{.}}"></script>
        {{end}}
    </head>

    <body>
    <div class="acp-wrapper">
        <div class="navigation">
            {{template "admin/navigation.tpl"}}
        </div>
        <div class="main">
            {{.LayoutContent}}
        </div>
    </div>

    {{range .FooterScripts}}
        <script src="{{.}}"></script>
    {{end}}
    </body>
</html>