import React from "react";
import Util from "./util";
import {worderApi} from "./api-client";

class ListsForm extends React.Component {
    constructor(props) {
        super(props);
        Util._bind(this, '_handleListSubmit');
    }

    _handleListSubmit(e) {
        e.preventDefault();
        let listName = this.refs.listName.value.trim();
        if (!listName) {
            return;
        }
        this.props.handleListSubmit({name: listName});
        this.refs.listName.value = '';
    };

    render() {
        return <div>
            <div className="modal fade add-new-list-modal" tabIndex="-1" role="dialog" aria-labelledby="mySmallModalLabel">
                <div className="modal-dialog modal-sm">

                    <div className="modal-content">

                        <div className="modal-header">
                            <button type="button" className="close" data-dismiss="modal" aria-label="Close">
                                <span aria-hidden="true">&times;</span></button>
                            <h4 className="modal-title">Add new list</h4>
                        </div>

                        <div className="modal-body">
                            <form className="new-lists-form">
                                <div className="form-group">
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

export default class ListsBox extends React.Component {
    constructor(props) {
        super(props);

        this.state = {lists: [this.props.listName], showAddListForm: false};
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
        var links = this.state.lists.map((listName) => {
            var className = (listName == this.props.listName) ? "active" : ""
            return <li className={className} key={listName}>
                <a href="#" onClick={(e) => {e.preventDefault(); this.props.onListSelect(listName)}}>
                    {listName}
                </a>
            </li>
        });
        return <div>
            <ul className="nav nav-pills">
                {links}
                <li>
                    <a href="#" className="glyphicon glyphicon-plus plus" data-toggle="modal" data-target=".add-new-list-modal"></a>
                </li>
            </ul>
            <ListsForm handleListSubmit={this._handleListSubmit}/>
        </div>
    }
}

//<button className="btn btn-sm btn-default">add new list</button>


