<h1>Add Place</h1>

{{if .flash.notice}}
<div class="flash">
    {{.flash.notice}}
</div>
{{end}}

<form action="" method="post" class="pure-form pure-form-aligned">
    <fieldset>
        <div class="pure-g">
            <div class="form-block pure-u-1">
                <input class="pure-input-1 form__input--title" id="name" name="name" type="text" placeholder="Place name" value="{{.Place.Name}}" required>
                {{if .Errors.Name}}<span class="form__hint--warn">{{.Errors.Name}}</span>{{end}}
            </div>
            <div class="form-block pure-u-1">
                <label for="epiceditor">Description</label>
                <textarea id="js-markdownSync" name="description">{{.Place.Description}}</textarea>
                <div id="epiceditor"></div>
                {{if .Errors.Description}}<span class="form__hint--warn">{{.Errors.Description}}</span>{{end}}
            </div>
            <div class="form-block pure-u-1 pure-u-md-1-3">
                <label for="website">Website URL</label>
                <input id="website" name="website" type="url" placeholder="Website URL" value="{{.Place.Website}}">
                {{if .Errors.Website}}<span class="form__hint--warn">{{.Errors.Website}}</span>{{end}}
            </div>
            <div class="form-block pure-u-1-2 pure-u-md-1-3">
                <label for="latitude">GPS Latitude</label>
                <input id="latitude" name="latitude" type="number" placeholder="GPS Latitude" step="any" value="{{.Place.Latitude}}">
                {{if .Errors.Latitude}}<span class="form__hint--warn">{{.Errors.Latitude}}</span>{{end}}
            </div>
            <div class="form-block pure-u-1-2 pure-u-md-1-3">
                <label for="longitude">GPS Longitude</label>
                <input id="longitude" name="longitude" type="number" placeholder="GPS Longitude" step="any" value="{{.Place.Longitude}}">
                {{if .Errors.Longitude}}<span class="form__hint--warn">{{.Errors.Longitude}}</span>{{end}}
            </div>
            <div class="form-block">
                <input type="hidden" name="_xsrf" value="{{.Token}}">
                <button type="submit" class="pure-button pure-button-primary">Add</button>
            </div>
        </div>
    </fieldset>
</form>

<!-- TODO: Move to controller -->
<script src="/static/js/epiceditor.js"></script>
<script>
    var opts = {
        textarea: "js-markdownSync",
        basePath: "/static/css/",
        theme: {
          base: "/epiceditor/base/epiceditor.css",
          preview: "/epiceditor/preview/github.css",
          editor: "/epiceditor/editor/epic-light.css"
        }
    }
    var editor = new EpicEditor(opts).load();
</script>