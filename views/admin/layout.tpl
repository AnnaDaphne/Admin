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
    {{if .Login}}
        {{.LayoutContent}}
    {{else}}
    <div class="acp-wrapper">
        <div class="navigation">
            <div class="pure-menu pure-menu-open">
                <a class="pure-menu-heading">Anna &amp; Daphne</a>
                <ul>
                    <li><a href="{{urlfor "PlaceController.Get"}}">View Places</a></li>
                    <li><a href="{{urlfor "PlaceController.AddPlace"}}">Add Place</a></li>
                    <li class="divider"><a href="#">View Tags</a></li>
                    <li><a href="#">Add Tags</a></li>
                </ul>
            </div>
        </div>
        <div class="main">
            {{.LayoutContent}}
        </div>
    </div>
    {{end}}

    {{range .FooterScripts}}
        <script src="{{.}}"></script>
    {{end}}
    </body>
</html>