import React from 'react';

import './Header.css'

class Header extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    render() {
        return (
            <div className="Header">
                <img src={process.env.PUBLIC_URL + 'favicon.ico'} alt="ICON" />
                <h1>{this.props.title}</h1>
            </div>
        );
    }
}

export default Header;
