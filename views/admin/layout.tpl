<!doctype html>
<html class="no-js" lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">

        <title>Anna &amp; Daphne &mdash; {{.Title}}</title>
        <meta name="description" content="">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <link rel="stylesheet" href="//cdn.jsdelivr.net/pure/0.5.0/pure-min.css">
        {{str2html "<!--[if lte IE 8]>"}}
            <link rel="stylesheet" href="//yui.yahooapis.com/pure/0.5.0/grids-responsive-old-ie-min.css">
        {{str2html "<![endif]-->"}}
        {{str2html "<!--[if gt IE 8]><!-->"}}
            <link rel="stylesheet" href="//yui.yahooapis.com/pure/0.5.0/grids-responsive-min.css">
        {{str2html "<!--<![endif]-->"}}
        <link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:400,500,300">
        <link rel="stylesheet" href="/static/css/admin.css">

        <script src="//cdn.jsdelivr.net/modernizr/2.8.3/modernizr.min.js"></script>
    </head>

    <body>
    <div class="acpWrapper">
        <div class="navigation">
            {{template "admin/navigation.tpl"}}
        </div>
        <div class="main">
            {{.LayoutContent}}
        </div>
    </div>
    </body>
</html>