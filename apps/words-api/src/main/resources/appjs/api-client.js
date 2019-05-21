import Util from "./util";

class WorderApiClient {

    static myHandleError(xhr, status, err) {
        console.error(status, err.toString());
        alert("error: " + err.toString());
    }

    storeLoginData(auth, login) {
        this.authHeader = auth;
        this.loggedUser = login;

        Cookies.set("auth", auth);
        Cookies.set("login", login);
    }

    loadLoginData(state) {
        this.authHeader = Cookies.get("auth");
        this.loggedUser = Cookies.get("login");

        state.auth = this.authHeader;
        state.username = this.loggedUser;
        state.logged = state.username != null
    }

    listUrl(listName) {
        return "/api/user/" + this.loggedUser + "/list" + (listName != null ? ('/' + listName) : '');
    }

    postText(data, url, success) {
        Util.logDebug("POST text at " + url)

        $.ajax({
            url: url,
            type: 'POST',
            data: data,
            contentType: 'plain/text; charset=utf-8',
            dataType: 'json',
            success: success,
            headers: {"Authorization": this.authHeader},
            error: this.myHandleError
        });
    }

    postJson(data, url, success) {
        Util.logDebug("POST json at " + url)

        $.ajax({
            url: url,
            type: 'POST',
            data: JSON.stringify(data),
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            success: success,
            headers: {"Authorization": this.authHeader},
            error: this.myHandleError
        });
    }

    getJson(url, success) {
        Util.logDebug("GET json at " + url)

        $.ajax({
            url: url,
            type: 'GET',
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            success: success,
            headers: {"Authorization": this.authHeader},
            error: this.myHandleError
        });
    }

    removeList(listName) {
        $.ajax({
            url: "/api/user/" + this.loggedUser + "/list/" + listName,
            type: 'DELETE',
            contentType: 'application/json; charset=utf-8',
            headers: {"Authorization": this.authHeader},
            error: this.myHandleError
        });
    }
}

export let worderApi = new WorderApiClient();