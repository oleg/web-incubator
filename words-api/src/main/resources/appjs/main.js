import React from "react"
import ReactDOM from "react-dom"
import Util from "./util"
import {worderApi} from "./api-client"
import {Word, WordsBox} from "./words"
import ListsBox from "./lists"
import Navigation from "./nav"
import ListEditor from "./manage"

class LoginBox extends React.Component {
    constructor() {
        super();
        Util._bind(this, 'handleSubmit');
    }

    handleSubmit(e) {
        e.preventDefault();

        var login = this.refs.login.value;
        var password = this.refs.password.value;
        var auth = "Basic " + btoa(login + ":" + password);
        var self = this;

        worderApi.storeLoginData(auth, login)
        //var success = this.props.onLoginSuccess.bind(this, {login: login, password: password});
        worderApi.getJson("/api/user/" + login, function (data) {
                self.props.onLoginSuccess({login: login, auth: auth});
            }
        );
    }

    render() {
        return <div className="loginBox">
            <form onSubmit={this.handleSubmit}>
                <input type="text" placeholder="Login" ref="login"/>
                <input type="password" placeholder="Password" ref="password"/>
                <input type="submit" value="login"/>
            </form>
        </div>
    }
}

class AnalyzeBox extends React.Component {
    constructor(props) {
        super(props);
        this.state = {activeList: "study", ignoreLists: [], words: []};
        Util._bind(this, '_showAnalyzed', '_changeList', '_submitToAnalyze', '_moveWord', '_onIgnoreListChange');
    }

    _showAnalyzed(newWords) {
        this.setState({words: newWords})
    }

    _changeList(listName) {
        this.setState({activeList: listName});
    }

    _onIgnoreListChange(added, listName) {
        Util.logDebug("ignoring list: " + listName)
        var newIgnoreList
        if (added) {
            newIgnoreList = this.state.ignoreLists.slice()
            newIgnoreList.push(listName)
        } else {
            newIgnoreList = this.state.ignoreLists.filter(l => l !== listName)
        }
        this.setState({ignoreLists: newIgnoreList});
    }

    _submitToAnalyze(e) {
        e.preventDefault();
        var text = this.refs.text.value;
        var json = {"text": text, "ignoreWordLists": this.state.ignoreLists};
        worderApi.postJson(json, "/api/tool/to-words", this._showAnalyzed);
    }

    _moveWord(word) {
        worderApi.postJson({"word": word}, worderApi.listUrl(this.state.activeList), function () {});
    }

    render() {
        return <div>
            <ListsBox listName={this.state.activeList} onListSelect={this._changeList}/>
            <AnalyzeListsBox onIgnoreListChange={this._onIgnoreListChange} />

            <div className="post">
                <h2 className="title">Text</h2>
                <div className="entry">

                    <div>{this.state.words.map(wc => <Word word={wc.word} count={wc.count} moveWord={this._moveWord}/>)}</div>

                    <button onClick={this._submitToAnalyze}>to words</button>
                    <br/>

                    <textarea ref="text"/>

                </div>
            </div>
        </div>
    }
}

class AnalyzeListsBox extends React.Component {
    constructor(props) {
        super(props);

        this.state = {lists: []};
        Util._bind(this, 'componentDidMount', '_loadListsFromServer', '_handleListSubmit');
    }

    componentDidMount() {
        this._loadListsFromServer();
    }

    _handleListSubmit(data) {
        worderApi.postJson(data, worderApi.listUrl(), this._loadListsFromServer)
    }

    _loadListsFromServer() {
        worderApi.getJson(worderApi.listUrl(), (data) => this.setState({lists: data}))
    }

    render() {
        var func = this.props.onIgnoreListChange;
        var checkboxes = this.state.lists.map((listName) => {
            return <div key={listName}>
                <input type="checkbox"
                       onChange={(data) => {func(data.target.checked, listName)}}/>{listName}
            </div>
        });
        return <div>{checkboxes}</div>
    }
}

class App extends React.Component {

    constructor(props) {
        super(props);

        let state = {activeList: "study"}
        worderApi.loadLoginData(state)

        this.state = state;
        Util._bind(this, 'handleSuccessLogin', 'handleListSelection', '_changeTopMenu', "_initTopMenu");
    }

    handleSuccessLogin(data) {
        this.setState({logged: true, username: data.login})
    }

    handleListSelection(data) {
        this.setState({activeList: data});
    }

    _changeTopMenu(data) {
        this.setState({topMenu: data})
    }

    _initTopMenu() {
        if (!this.state.topMenu) {
            if (window.location.hash) {
                this.state.topMenu = window.location.hash.substring(1)
            } else {
                this.state.topMenu = Navigation.Menu().Lists;
            }
        }
    }

    render() {
        this._initTopMenu();

        var body;
        if (!this.state.logged) {
            body = <LoginBox onLoginSuccess={this.handleSuccessLogin}/>
        } else if (this.state.topMenu == Navigation.Menu().Analyze) {
            body = <AnalyzeBox user={this.state.username}/>
        } else if (this.state.topMenu == Navigation.Menu().ManageLists) {
            body = <ListEditor />
        } else if (this.state.topMenu == Navigation.Menu().Lists) {
            body = <div>
                <ListsBox listName={this.state.activeList} onListSelect={this.handleListSelection}/>
                <WordsBox listName={this.state.activeList}/>
            </div>
        } else if (this.state.topMenu == Navigation.Menu().About) {
            body = <h2>About</h2>
        } else {
            body = <h2>Not found menu {this.state.topMenu}</h2>
        }
        return <div className="container">
            <Navigation topMenu={this.state.topMenu} onTopMenuClick={this._changeTopMenu}/>
            {body}
        </div>
    }
}

ReactDOM.render(<App/>, document.getElementById('application'));
