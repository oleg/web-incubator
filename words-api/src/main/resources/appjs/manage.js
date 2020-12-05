import React from "react"
import Util from "./util"
import {worderApi} from "./api-client"

class ListEditForm extends React.Component {
    constructor(props) {
        super(props);
        Util._bind(this, '_handleListSubmit');
    }

    _handleListSubmit(e) {
        e.preventDefault();
        var listName = this.refs.listName.value.trim();
        if (!listName) {
            return;
        }
        this.props.handleListSubmit({name: listName});
        this.refs.listName.value = '';
    };

    render() {
        return <div>
            <div className="modal fade add-new-list-modal" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel">
                <div className="modal-dialog modal-sm">

                    <div className="modal-content">

                        <div className="modal-header">
                            <button type="button" className="close" data-dismiss="modal" aria-label="Close">
                                <span aria-hidden="true">&times;</span></button>
                            <h4 className="modal-title">Add new list</h4>
                        </div>

                        <div className="modal-body">
                            <form className="new-lists-form">
                                <div class="form-group">
                                    <input className="form-control" type="text" placeholder="New list" ref="listName"/>
                                </div>
                            </form>
                        </div>

                        <div className="modal-footer">
                            <button type="button" className="btn btn-default" data-dismiss="modal">Close</button>
                            <button type="button" className="btn btn-primary" data-dismiss="modal"
                                    onClick={this._handleListSubmit}>Create
                            </button>
                        </div>

                    </div>

                </div>
            </div>
        </div>
    }

}

export default class ListEditor extends React.Component {
    constructor(props) {
        super(props);

        this.state = {lists: [this.props.listName], showAddListForm: false};
        Util._bind(this, 'componentDidMount', '_loadListsFromServer', '_handleListSubmit', '_removeList');
    }

    componentDidMount() {
        this._loadListsFromServer();
        // setInterval(this._loadListsFromServer, 5000);
    }

    _handleListSubmit(data) {
        worderApi.postJson(data, worderApi.listUrl(), this._loadListsFromServer)
    }

    _loadListsFromServer() {
        worderApi.getJson(worderApi.listUrl(), (data) => this.setState({lists: data}))
    }

    _removeList(listName) {
        worderApi.removeList(listName)
    }

    render() {
        var items = this.state.lists.map((listName) => {
            return <div className="row" key={"LE." + listName}>
                <div className="col-md-1">{listName}</div>
                <div className="col-md-2">
                    <div className="btn-group">
                        {/*<button className="btn btn-default btn-flat"><span className="glyphicon glyphicon-plus plus"/></button>*/}
                        <button onClick={()=>this._removeList(listName)} className="btn btn-default btn-flat"><span className="glyphicon glyphicon-trash plus"/></button>
                    </div>
                </div>
            </div>

        });
        return <div>
            {items}
        </div>
    }
}
