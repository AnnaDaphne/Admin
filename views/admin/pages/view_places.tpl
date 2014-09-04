<h1>Places</h1>

{{if .flash.notice}}
<div class="flash">
    {{.flash.notice}}
</div>
{{else}}
<table class="places pure-table pure-table-horizontal">
    <thead>
        <tr>
            <th>Name</th>
            <th>Description</th>
            <th>GPS (X, Y)</th>
        </tr>
    </thead>

    <tbody>
        {{range $Place := .Places}}
            <tr>
                <td>{{$Place.Name}}</td>
                <td>{{substr $Place.Description 0 200}} ...</td>
                <td>{{$Place.Longitude}}, {{$Place.Latitude}}</td>
            </tr>
        {{end}}
    </tbody>
</table>
{{end}}