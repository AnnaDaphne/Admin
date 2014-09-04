<h1>Create Place</h1>

<form action="" method="post" class="pure-form pure-form-aligned">
    <fieldset>
        <div class="pure-control-group">
            <label for="name">Name</label>
            <input id="name" name="name" type="text" placeholder="Name">
        </div>
        <div class="pure-control-group">
            <label for="epiceditor">Description</label>
            <textarea id="js-description" name="description"></textarea>
            <div id="epiceditor"></div>
        </div>
        <div class="pure-control-group">
            <label for="website">Website URL</label>
            <input id="website" name="website" type="text" placeholder="Website URL">
        </div>
        <div class="pure-control-group">
            <label for="gpsLat">GPS Latitude</label>
            <input id="gpsLat" name="gpsLat" type="text" placeholder="GPS Latitude">
        </div>
        <div class="pure-control-group">
            <label for="gpsLong">GPS Longitude</label>
            <input id="gpsLong" name="gpsLong" type="text" placeholder="GPS Longitude">
        </div>

        <div class="pure-controls">
            <button type="submit" class="pure-button pure-button-primary">Create</button>
        </div>
    </fieldset>
</form>

<script src="/static/js/epiceditor.js"></script>
<script>
    var opts = {
        textarea: "js-description",
        basePath: "/static/css/",
        theme: {
          base: "/epiceditor/base/epiceditor.css",
          preview: "/epiceditor/preview/github.css",
          editor: "/epiceditor/editor/epic-light.css"
        }
    }
    var editor = new EpicEditor(opts).load();
</script>