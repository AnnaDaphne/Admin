<div class="login-wrapper">
    <h1 class="login__title">Anna &amp; Daphne</h1>

    <div class="pure-g">
        <div class="login-block pure-u-1 pure-u-md-1-3">
            <h3 class="login__title--sub">ACP Login</h3>
            <form action="" method="post" class="login__form pure-form pure-form-aligned">
                <fieldset>
                    {{if .flash.error}}
                        <div class="flash">
                            {{.flash.error}}
                        </div>
                    {{end}}
                    <div class="form-block pure-control-group">
                        <label for="name">Username</label>
                        <input id="name" name="username" type="text" placeholder="Username" required>
                    </div>
                    <div class="form-block pure-control-group">
                        <label for="password">Password</label>
                        <input id="password" name="password" type="password" placeholder="Password" required>
                    </div>
                </fieldset>
                <input type="hidden" name="_xsrf" value="{{.Token}}">
                <button class="pure-button login__button">Login</button>
            </form>
        </div>
    </div>
</div>