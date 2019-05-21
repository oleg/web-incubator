import React from "react"

class NavigationLink extends React.Component {

    constructor(props) {
        super(props)
    }

    render() {
        var className = this.props.itemName == this.props.topMenu ? "active" : "";
        return <li className={className}>
            <a href={"#" + this.props.itemName} onClick={() => this.props.onClick(this.props.itemName)}>
                {this.props.itemName}
            </a>
        </li>
    }
}

const Menu = {
    ManageLists: "Manage-Lists",
    About: "About",
    Analyze: "Analyze",
    Lists: "Lists",
}

export default class Navigation extends React.Component {

    static Menu() { return Menu }

    render() {
        return <nav className="navbar navbar-default">
            
            <div className="container-fluid">
                <div className="navbar-header">
                    <a className="navbar-brand" href="#">Worder</a>
                </div>
                <div className="navbar-collapse collapse" id="navbar">
                    <ul className="nav navbar-nav">
                        <NavigationLink onClick={this.props.onTopMenuClick} topMenu={this.props.topMenu} itemName={Menu.Lists}/>
                        <NavigationLink onClick={this.props.onTopMenuClick} topMenu={this.props.topMenu} itemName={Menu.ManageLists}/>
                        <NavigationLink onClick={this.props.onTopMenuClick} topMenu={this.props.topMenu} itemName={Menu.Analyze}/>
                        <NavigationLink onClick={this.props.onTopMenuClick} topMenu={this.props.topMenu} itemName={Menu.About}/>
                    </ul>
                </div>
            </div>
            
        </nav>
    }
}





