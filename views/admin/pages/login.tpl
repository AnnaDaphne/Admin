<div class="login-wrapper">
    <h1 class="login__title">Anna &amp; Daphne</h1>

    <div class="pure-g">
        <div class="login-block pure-u-1 pure-u-md-1-3">
            <h3 class="login__title--sub">ACP Login</h3>
            <form class="login__form pure-form pure-form-aligned">
                <fieldset>
                    {{if .flash.error}}
                    <div class="form-block pure-control-group">
                        <div class="flash">
                            {{.flash.error}}
                        </div>
                    </div>
                    {{end}}
                    <div class="form-block pure-control-group">
                        <label for="name">Username</label>
                        <input id="name" type="text" placeholder="Username" required>
                    </div>
                    <div class="form-block pure-control-group">
                        <label for="password">Password</label>
                        <input id="password" type="password" placeholder="Password" required>
                    </div>
                </fieldset>
                <button class="pure-button login__button">Login</button>
            </form>
        </div>
    </div>
</div>